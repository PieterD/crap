package main

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/PieterD/kafka-processor/kafka"
	"github.com/PieterD/kafka-processor/kafka/internal/killchan"
)

type threadHandler struct {
	kfk     *kafka.Kafka
	actions map[string]threadAction
	kill    *killchan.Killchan
	wg      sync.WaitGroup
}

type threadAction struct {
	dst      string
	function string
}

func newThreadHandler(k *kafka.Kafka, cfg *Config) (*threadHandler, error) {
	th := &threadHandler{
		kfk:     k,
		actions: make(map[string]threadAction),
		kill:    killchan.New(),
	}

	for _, stream := range cfg.Stream {
		for _, partition := range stream.Partition {
			err := k.ListenNewest(stream.TopicSrc, partition)
			if err != nil {
				//TODO: Is this right?
				return nil, fmt.Errorf("Failed to add listener to stream %s:%d: %v", stream.TopicSrc, partition, err)
			}
			th.actions[stream.TopicSrc] = threadAction{
				dst:      stream.TopicDst,
				function: stream.Function,
			}
		}
	}

	th.wg.Add(cfg.Threads)
	for i := 0; i < cfg.Threads; i++ {
		go th.run()
	}

	return th, nil
}

func (th *threadHandler) wait() {
	th.wg.Wait()
}

func (th *threadHandler) close() {
	th.kill.Kill()
}

func (th *threadHandler) run() {
	defer th.wg.Done()
	for {
		var msg kafka.Message

		select {
		case <-th.kill.Chan():
			return
		case msg = <-th.kfk.Incoming():
		}

		action, ok := th.actions[msg.Topic]
		if !ok {
			logger.Panicf("Message received from topic we are not supposed to be listening to")
		}

		select {
		case <-th.kill.Chan():
			return
		case trans := <-th.kfk.Outgoing():
			_, _, err := trans.Send(msg.Key, bytes.ToUpper(msg.Val), action.dst)
			if err != nil {
				logger.Panicf("Message send failed to %s: %v", action.dst, err)
			}
		}
	}
}
