package helper

import (
	config "connectHub_gateway/pkg/config"
	"connectHub_gateway/pkg/utils/models"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/websocket"
)

type Helper struct {
	config *config.Config
}

func NewHelper(config *config.Config) *Helper {
	return &Helper{
		config: config,
	}
}

func (r *Helper) SendMessageToUser(User map[string]*websocket.Conn, msg []byte, userID string) {
	var message models.Message
	if err := json.Unmarshal([]byte(msg), &message); err != nil {
		fmt.Println("error while unmarshel ", err)
	}

	message.SenderID = userID
	recipientConn, ok := User[message.RecipientID]

	if ok {
		recipientConn.WriteMessage(websocket.TextMessage, msg)
	}
	err := KafkaProducer(message)
	fmt.Println("==sending succesful==", err)
}


func KafkaProducer(message models.Message) error {
	fmt.Println("from kafka ", message)

	cfg, _ := config.LoadConfig()
	configs := sarama.NewConfig()
	configs.Producer.Return.Successes = true
	configs.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer([]string{cfg.KafkaPort}, configs)
	if err != nil {
		fmt.Println("error creating producer:", err)
		return err
	}
	fmt.Println("producer created successfully:", producer)

	result, err := json.Marshal(message)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: cfg.KafkaTopic,
		Key:   sarama.StringEncoder("Friend message"),
		Value: sarama.StringEncoder(result),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("error sending message:", err)
		return err
	}

	log.Printf("[producer] partition id: %d; offset:%d, value: %v\n", partition, offset, msg)
	fmt.Println("==sending successful==")
	return nil
}

func (h *Helper) ValidateTokenJobseeker(tokenString string) (int, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.config.JobSeekerAccessKey), nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["id"].(float64)
		intUserID := int(userID)

		if !ok {
			return 0, errors.New("user_id not found in token")
		}
		return intUserID, nil
	} else {
		return 0, errors.New("invalid token")
	}
}

func (h *Helper) ValidateTokenRecruiter(tokenString string) (int, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.config.RecruiterAccessKey), nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["id"].(float64)
		intUserID := int(userID)

		if !ok {
			return 0, errors.New("user_id not found in token")
		}
		return intUserID, nil
	} else {
		return 0, errors.New("invalid token")
	}
}
