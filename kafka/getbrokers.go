package kafka

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func (k *Kafka) GetKafkaBrokersFromZookeeper() ([]BrokerInfo, error) {
	brokerids, _, err := k.zooConn.Children("/brokers/ids")
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch Kafka brokers: %v", err)
	}

	var brokers []BrokerInfo

	for _, brokerid := range brokerids {
		data, _, err := k.zooConn.Get("/brokers/ids/" + brokerid)
		if err != nil {
			k.logger.Printf("Failed to get Kafka broker: %v", err)
			continue
		}
		var broker BrokerInfo
		broker.Id, err = strconv.Atoi(brokerid)
		if err != nil {
			k.logger.Printf("Failed to parse brokerid: %v", err)
			continue
		}
		err = json.Unmarshal(data, &broker)
		if err != nil {
			k.logger.Printf("Failed to parse Kafka broker json: %v", err)
			continue
		}
		brokers = append(brokers, broker)
	}

	return brokers, nil
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

func brokerStrings(brokers []BrokerInfo) []string {
	strs := make([]string, len(brokers))
	for i := range strs {
		strs[i] = brokers[i].Addr()
	}
	return strs
}
