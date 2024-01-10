package kafka_consumer

import "testing"

var (
	servers []string
	group   string
	topics  []string
)

func Test_kafka(t *testing.T) {
	servers = []string{"{servers}"}
	group = "{group}"
	topics = []string{"{topics}"}

	Consumer(servers, group, topics)
}
