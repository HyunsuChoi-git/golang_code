package kafka_consumer

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
	"time"
)

var (
	consumerGroup sarama.ConsumerGroup
)

func connectKafkaConsumer(servers []string, group string) error {
	fnc := "connectKafkaConsumer"

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	var err error
	consumerGroup, err = sarama.NewConsumerGroup(servers, group, config)
	if err != nil {
		return err
	}

	log.Printf("%s: Kafka Consumer Connect Successfully...", fnc)
	return nil
}

func Consumer(servers []string, group string, topics []string) sarama.PartitionConsumer {
	fnc := "Consumer"
	log.Printf("%s: Kafka Consumer Connect...", fnc)

	go func() {

		var err error

		defer func() {
			if r := recover(); r != nil {
				log.Error().Msgf("consume error: %v", r)
			}
		}()

		for {
			/* kafka connection - 실패시 2초 뒤 반복*/
			err = connectKafkaConsumer(servers, group)
			if err != nil {
				log.Printf("%s: Kafka Consumer Connection Error!", fnc)
				time.Sleep(2 * time.Second)
				continue
			}

			handler := KafkaConsumerHandler{}
			err = consumerGroup.Consume(context.Background(), topics, handler)

			if err != nil {
				time.Sleep(2 * time.Second)
			}

		}
	}()

	return nil
}
