package kafka_consumer

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
	"time"
)

var expiryTime = 3000000

type KafkaConsumerHandler struct{}

func (h KafkaConsumerHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h KafkaConsumerHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h KafkaConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	//fnc := "ConsumeClaim"

	for kafkaMessage := range claim.Messages() {
		/* 오래된 메세지 파기 */
		if kafkaMessage.Timestamp.Before(time.Now().Add(-time.Duration(expiryTime) * time.Millisecond)) {
			//util.PrintServiceWarnLog(Func: fnc, ServiceType: ServiceType, Status: "DELETE MESSAGE",
			//	Message: "this message's send time has passed. time: " + kafkaMessage.Timestamp.String(), Data: string(kafkaMessage.Value)})
			session.MarkMessage(kafkaMessage, "") // offset commit : 읽은 '메시지 위치 정보(offset)'를 기록(commit)
			continue
		}

		/* message byte로 파싱 */
		metricCollectorInfoBytes := kafkaMessage.Value
		maps := map[string]interface{}{}
		if err := parseMessage(metricCollectorInfoBytes, &maps); err != nil {
			session.MarkMessage(kafkaMessage, "")
			continue
		}

		/* message Input interface로 파싱 */
		var newData struct {
			// parsing할 data 타입...
		}

		log.Printf("%v", newData)

		session.MarkMessage(kafkaMessage, "") // offset commit : 읽은 '메시지 위치 정보(offset)'를 기록(commit)
	}
	return nil

}

func parseMessage(msgByte []byte, msgStruct interface{}) error {
	fnc := "parseMessage"
	/* Kafka Message 기반 수집 */
	if err := json.Unmarshal(msgByte, &msgStruct); err != nil {

		log.Printf("%s: marshaling is failed...", fnc)
		return err
	}
	return nil
}
