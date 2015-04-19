package transmithandler

import (
	"github.com/PieterD/kafka-processor/kafka/internal/killchan"
	"github.com/Shopify/sarama"
)

type TransmitHandler struct {
	kill     *killchan.Killchan
	dead     *killchan.Killchan
	bus      chan Transmit
	producer sarama.SyncProducer
}

type Transmit struct {
	ch chan<- transmitRequest
}

type transmitRequest struct {
	topic    string
	key, val []byte
	response chan<- transmitResponse
}

type transmitResponse struct {
	partition int32
	offset    int64
	err       error
}

func New(kfkConn sarama.Client) (*TransmitHandler, error) {
	producer, err := sarama.NewSyncProducerFromClient(kfkConn)
	if err != nil {
		return nil, err
	}
	th := &TransmitHandler{
		kill:     killchan.New(),
		dead:     killchan.New(),
		bus:      make(chan Transmit),
		producer: producer,
	}
	go th.run()
	return th, nil
}

func (th *TransmitHandler) Close() {
	th.kill.Kill()
	th.dead.Wait()
}

func (th *TransmitHandler) run() {
	defer th.dead.Kill()
	defer th.producer.Close()
	defer close(th.bus)
	transChan := make(chan transmitRequest)
	trans := Transmit{ch: transChan}
	for {
		select {
		case <-th.kill.Chan():
			return
		case th.bus <- trans:
			req := <-transChan
			var key sarama.ByteEncoder
			if req.key != nil {
				key = sarama.ByteEncoder(req.key)
			}
			message := &sarama.ProducerMessage{
				Topic: req.topic,
				Key:   key,
				Value: sarama.ByteEncoder(req.val),
			}
			partition, offset, err := th.producer.SendMessage(message)
			req.response <- transmitResponse{
				partition: partition,
				offset:    offset,
				err:       err,
			}
		}
	}
}

func (th *TransmitHandler) Outgoing() <-chan Transmit {
	return th.bus
}

func (t Transmit) Send(key, val []byte, topic string) (partition int32, offset int64, err error) {
	respChan := make(chan transmitResponse, 1)
	req := transmitRequest{
		topic:    topic,
		key:      key,
		val:      val,
		response: respChan,
	}
	t.ch <- req
	resp := <-respChan
	return resp.partition, resp.offset, resp.err
}
