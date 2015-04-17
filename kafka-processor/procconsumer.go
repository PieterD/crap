package main

import (
	"bytes"
	"fmt"
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
	fmt.Printf("started %s:%d -> %s\n", pc.src, pc.partition, pc.dst)
	for msg := range pc.consumer.Messages() {
		//TODO: more general function
		newval := bytes.ToUpper(msg.Value)
		message := &sarama.ProducerMessage{
			Topic: pc.dst,
			Value: sarama.ByteEncoder(newval),
		}
		_, _, err := pc.producer.SendMessage(message)
		if err != nil {
			logger.Printf("Failed to send message to %s (offset %d): %v", pc.src, msg.Offset, err)
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
