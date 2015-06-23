package zoohandler

import "fmt"

func (zh *ZooHandler) GetKafkaBrokers() (BrokerList, error) {
	node, err := zh.NewNode("/brokers/ids")
	if err != nil {
		return nil, err
	}

	result := node.GetChildren(false)
	if result.Err != nil {
		return nil, result
	}

	var brokers BrokerList
	for _, childnode := range result.Children {
		brokerid, ok := childnode.NumberName()
		if !ok {
			continue
		}
		child := childnode.Get(false)
		if child.Err != nil {
			zh.logger.Printf("Failed to get Kafka broker (%s): %v", child.name, child)
			continue
		}
		var broker BrokerInfo
		broker.Id = int(brokerid)
		err = child.Json(&broker)
		if err != nil {
			zh.logger.Printf("Failed to parse Kafka broker (%s) json: %v", child.name, err)
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
