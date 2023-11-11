package kafkaclient

import (
	"context"
	"encoding/json"
	"log"
	"signature-app/database/repository"

	kafka "github.com/segmentio/kafka-go"
)

type kafkaClient struct {
	ReaderToken     *kafka.Reader
	WriterToken     *kafka.Writer
	ReaderPendingTx *kafka.Reader
	WriterPendingTx *kafka.Writer
	ReaderNewBlock  *kafka.Reader
	WriterNewBlock  *kafka.Writer
	Topic           KafkaTopics
	Ctx             context.Context
	Db              *repository.Database
}

type KafkaTopics struct {
	Token     string
	PendingTx string
	NewBlock  string
}

type NewPendingTransaction struct {
	TxId      string
	TimeStamp string
}

type KafkaClientInterface interface {
	ProduceMessageToken(msg []byte) error
	ReadMessagesToken()
	ProduceMessagePendingTx(msg []byte) error
	ReadMessagesPendingTx()
}

func NewKafkaClient(ctx context.Context, kafkaServerAddress, group string, db *repository.Database, topics KafkaTopics) KafkaClientInterface {
	_, err := kafka.DialLeader(ctx, "tcp", kafkaServerAddress, topics.Token, 0)
	if err != nil {
		log.Fatal("error connecting to kafka: ", err)
	}

	readerToken := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaServerAddress},
		GroupID:  group,
		Topic:    topics.Token,
		MaxBytes: 10e6,
	})

	writerToken := &kafka.Writer{
		Addr:                   kafka.TCP(kafkaServerAddress),
		Topic:                  topics.Token,
		AllowAutoTopicCreation: true,
	}

	readerPendingTx := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaServerAddress},
		GroupID:  group,
		Topic:    topics.PendingTx,
		MaxBytes: 10e6,
	})

	writerPendingTx := &kafka.Writer{
		Addr:                   kafka.TCP(kafkaServerAddress),
		Topic:                  topics.PendingTx,
		AllowAutoTopicCreation: true,
	}

	readerNewBlock := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaServerAddress},
		GroupID:  group,
		Topic:    topics.NewBlock,
		MaxBytes: 10e6,
	})

	writerNewBlock := &kafka.Writer{
		Addr:                   kafka.TCP(kafkaServerAddress),
		Topic:                  topics.NewBlock,
		AllowAutoTopicCreation: true,
	}

	return &kafkaClient{
		ReaderToken:     readerToken,
		WriterToken:     writerToken,
		ReaderPendingTx: readerPendingTx,
		WriterPendingTx: writerPendingTx,
		ReaderNewBlock:  readerNewBlock,
		WriterNewBlock:  writerNewBlock,
		Topic:           topics,
		Ctx:             ctx,
		Db:              db,
	}
}

func (kc *kafkaClient) ProduceMessageToken(msg []byte) error {
	err := kc.WriterToken.WriteMessages(
		kc.Ctx,
		kafka.Message{
			Value: msg,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (kc *kafkaClient) ReadMessagesToken() {
	var data map[string]string

	for {
		msg, err := kc.ReaderToken.ReadMessage(kc.Ctx)
		if err != nil {
			log.Printf("Error reading kafka data : %+v", err)
		}
		
		log.Println("new kafka token message arrived")

		err = json.Unmarshal(msg.Value, &data)
		if err == nil {
			kc.Db.AddNewToken(data["address"], data["token"])
		} else {
			log.Println("[ReadMessagesToken]error unmarshall the kafka message")
		}
	}
}

func (kc *kafkaClient) ProduceMessagePendingTx(msg []byte) error {
	err := kc.WriterPendingTx.WriteMessages(
		kc.Ctx,
		kafka.Message{
			Value: msg,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (kc *kafkaClient) ReadMessagesPendingTx() {
	var data NewPendingTransaction

	for {
		msg, err := kc.ReaderPendingTx.ReadMessage(kc.Ctx)
		if err != nil {
			log.Printf("Error reading kafka data : %+v", err)
		}

		err = json.Unmarshal(msg.Value, &data)
		if err == nil {
			ask, err := kc.Db.GetOneAsk(data.TxId)
			if err != nil || len(ask) != 1 {
				log.Println("Error getting an ask : ", err)
				continue
			}

			err = kc.Db.UpdatePendingAsk(data.TxId, data.TimeStamp)
			if err != nil {
				log.Println("Error updating pending ask")
				continue
			}
			log.Println("updating new pending tx")
		} else {
			log.Println("[ReadMessagesPendingTx()] error unmarshall the kafka message")
		}
	}
}
