package main

import (
	"encoding/xml"
	"os"
)

type Config struct {
	Name      string          `xml:"Name"`
	Threads   int             `xml:"Threads"`
	Zookeeper ConfigZooKeeper `xml:"Zookeeper"`
	Stream    []ConfigStream  `xml:"Streams>Stream"`
}

type ConfigZooKeeper struct {
	Peer []string `xml:"Peers>Peer"`
}

type ConfigKafka struct {
	Peer []string `xml:"Peers>Peer"`
}

type ConfigStream struct {
	TopicSrc  string  `xml:"Topic>Src"`
	TopicDst  string  `xml:"Topic>Dst"`
	Partition []int32 `xml:"Partitions>Partition"`
	Function  string  `xml:"Function"`
}

func ParseConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := xml.NewDecoder(file)
	config := new(Config)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
