package helper

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ARunni/ConnetHub_job/pkg/config"
	"github.com/ARunni/ConnetHub_job/pkg/utils/models"
	"github.com/IBM/sarama"
)

func SendNotification(data models.Notification, msg []byte) {
	data.Message = string(msg)
	err := KafkaNotificationProducer(data)
	if err != nil {
		fmt.Println("error sending notification to Kafka:", err)
		return
	}

	fmt.Println("==sent job notification successfully to Kafka==")
}

func KafkaNotificationProducer(message models.Notification) error {
	cfg, _ := config.LoadConfig()
	configs := sarama.NewConfig()
	configs.Producer.Return.Successes = true
	configs.Producer.RequiredAcks = sarama.WaitForAll
	configs.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer([]string{cfg.KafkaBrokers}, configs)
	if err != nil {
		return err
	}
	result, err := json.Marshal(message)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{Topic: cfg.KafkaTopic, Key: sarama.StringEncoder("Notifications"), Value: sarama.StringEncoder(result)}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("err send message in kafka ", err)
	}
	log.Printf("[producer] partition id: %d; offset:%d, value: %v\n", partition, offset, msg)
	return nil
}
