package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/samuel/go-zookeeper/zk"
)

func GetBrokers(zooConn *zk.Conn) []BrokerInfo {
	brokerids, _, err := zooConn.Children("/brokers/ids")
	if err != nil {
		logger.Panicf("Failed to fetch Kafka brokers: %v", err)
	}

	var brokers []BrokerInfo

	for _, brokerid := range brokerids {
		data, _, err := zooConn.Get("/brokers/ids/" + brokerid)
		if err != nil {
			logger.Printf("Failed to get Kafka broker: %v", err)
			continue
		}
		var broker BrokerInfo
		broker.Id, err = strconv.Atoi(brokerid)
		if err != nil {
			logger.Printf("Failed to parse brokerid: %v", err)
			continue
		}
		err = json.Unmarshal(data, &broker)
		if err != nil {
			logger.Printf("Failed to parse Kafka broker json: %v", err)
			continue
		}
		brokers = append(brokers, broker)
	}

	return brokers
}

type BrokerInfo struct {
	Id      int    `xml:"-"`
	JmxPort int64  `xml:"jmx_port"`
	Host    string `xml:"host"`
	Version int    `xml:"version"`
	Port    int    `xml:"port"`
}

func (bi BrokerInfo) Addr() string {
	return fmt.Sprintf("%s:%d", bi.Host, bi.Port)
}
