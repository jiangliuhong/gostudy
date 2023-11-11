package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

// Producer kafka 消息生产者
type Producer struct {
	Brokers   []string
	Topic     string
	Partition int32
	Conf      *sarama.Config
}

// ConnAndSend 连接并发送消息
func (k *Producer) ConnAndSend(msg []byte) (err error) {

	syncProducer, err := sarama.NewSyncProducer(k.Brokers, k.Conf)
	if err != nil {
		return err
	}
	partition, offset, err := syncProducer.SendMessage(&sarama.ProducerMessage{
		Topic:     k.Topic,
		Partition: k.Partition,
		Value:     sarama.ByteEncoder(msg),
	})
	if err != nil {
		return err
	}
	_ = syncProducer.Close()
	log.Printf(fmt.Sprintf("send message success: %d %d", partition, offset))
	return nil
}

var producer *Producer

type NoOpenProducer struct {
}

func (p *NoOpenProducer) ConnAndSend(msg []byte) (err error) {
	return nil
}

func InitKafkaProducerBuilder() {
	conf := sarama.NewConfig()
	conf.Producer.Retry.Max = 3
	conf.Producer.RequiredAcks = sarama.WaitForLocal
	conf.Producer.Return.Successes = true
	conf.Producer.MaxMessageBytes = 1024 * 1024 * 10 // 10M
	conf.Metadata.Full = true
	conf.Version = sarama.V2_5_0_0

	conf.Net.SASL.Enable = true
	conf.Net.SASL.User = "admin"
	conf.Net.SASL.Password = "zaq1@WSX"
	conf.Net.SASL.Handshake = true

	producer = &Producer{
		Brokers:   []string{"10.100.44.58:9092"},
		Topic:     "BACKUP_TASK_REPLY",
		Partition: 0,
		Conf:      conf,
	}
}

func main() {
	InitKafkaProducerBuilder()
	msg := `{
    "sourceId": "5",
    "status": "complete",
    "message": "执行成功",
    "filePath": "/backup/0235hnCDihMFgHUsxWHM.tar.gz",
    "fileType": "minio"
}`
	producer.ConnAndSend([]byte(msg))
}
