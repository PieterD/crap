package zoohandler

import (
	"fmt"
	"log"
	"time"

	"github.com/PieterD/kafka/internal/killchan"
	"github.com/samuel/go-zookeeper/zk"
)

type ZooHandler struct {
	kill   *killchan.Killchan
	dead   *killchan.Killchan
	conn   *zk.Conn
	logger *log.Logger
}

func New(peers []string, logger *log.Logger) (*ZooHandler, error) {
	conn, _, err := zk.Connect(peers, time.Second)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to zookeeper: %v", err)
	}
	zh := &ZooHandler{
		kill:   killchan.New(),
		dead:   killchan.New(),
		conn:   conn,
		logger: logger,
	}
	go zh.run()
	return zh, nil
}

func (zh *ZooHandler) Close() {
	zh.kill.Kill()
	zh.dead.Wait()
}

func (zh *ZooHandler) run() {
	defer zh.dead.Kill()
	defer zh.conn.Close()
	zh.kill.Wait()
}
