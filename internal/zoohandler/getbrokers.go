package zoohandler

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func (zh *ZooHandler) GetKafkaBrokers() (BrokerList, error) {
	brokerids, _, err := zh.conn.Children("/brokers/ids")
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch Kafka brokers: %v", err)
	}

	var brokers BrokerList

	for _, brokerid := range brokerids {
		data, _, err := zh.conn.Get("/brokers/ids/" + brokerid)
		if err != nil {
			zh.logger.Printf("Failed to get Kafka broker: %v", err)
			continue
		}
		var broker BrokerInfo
		broker.Id, err = strconv.Atoi(brokerid)
		if err != nil {
			zh.logger.Printf("Failed to parse brokerid: %v", err)
			continue
		}
		err = json.Unmarshal(data, &broker)
		if err != nil {
			zh.logger.Printf("Failed to parse Kafka broker json: %v", err)
			continue
		}
		brokers = append(brokers, broker)
	}

	return brokers, nil
}

type BrokerList []BrokerInfo

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

func (bl BrokerList) Strings() []string {
	strs := make([]string, len(bl))
	for i := range strs {
		strs[i] = bl[i].Addr()
	}
	return strs
}
