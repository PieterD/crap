package main

import (
	"bytes"
	"sync/atomic"

	"github.com/Shopify/sarama"
)

type ProcConsumer struct {
	consumer  sarama.PartitionConsumer
	producer  sarama.SyncProducer
	src, dst  string
	partition int32
	function  []string
	killchan  chan struct{}
	isclosed  int32
}

func (pc *ProcConsumer) Close() {
	if atomic.CompareAndSwapInt32(&pc.isclosed, 0, 1) {
		pc.consumer.Close()
		<-pc.killchan
	}
}

func (pc *ProcConsumer) Run() {
	defer pc.Close()
	defer close(pc.killchan)
	for msg := range pc.consumer.Messages() {
		//TODO: more general function
		newval := bytes.ToUpper(msg.Value)
		message := &sarama.ProducerMessage{
			Topic:     pc.dst,
			Partition: pc.partition,
			Value:     sarama.ByteEncoder(newval),
		}
		_, _, err := pc.producer.SendMessage(message)
		if err != nil {
			logger.Panicf("Failed to send message: %v", err)
		}
	}
}

type ConsumerList struct {
	pc       []*ProcConsumer
	killchan chan struct{}
}

func NewConsumerList() *ConsumerList {
	cl := new(ConsumerList)
	cl.killchan = make(chan struct{})
	return cl
}

func (cl *ConsumerList) Close() {
	defer close(cl.killchan)
	for i := range cl.pc {
		cl.pc[i].Close()
	}
}

func (cl *ConsumerList) Add(pc *ProcConsumer) {
	go pc.Run()
	cl.pc = append(cl.pc, pc)
}
