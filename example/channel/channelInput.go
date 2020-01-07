package channel

/*
type ProducerMessage struct {
	Key         Encoder
	Value       Encoder
	Headers     []RecordHeader
	Metadata    interface{}
	Offset      int64
	Partition   int32
	Timestamp   time.Time
	retries     int
	flags       flagSet
	expectation chan *ProducerError
}
type ProducerError struct {
	Msg *ProducerMessage
	Err error
}

type asyncProducer struct {
	client                    Client
	conf                      *Config
	ownClient                 bool
	errors                    chan *ProducerError
	input, successes, retries chan *ProducerMessage
	inFlight                  sync.WaitGroup
	brokers                   map[*Broker]chan<- *ProducerMessage
	brokerRefs                map[chan<- *ProducerMessage]int
	brokerLock                sync.Mutex
}

func (p *asyncProducer) Input() chan<- *ProducerMessage {
	return p.input
}

func TestInput() {

	msg := &ProducerMessage{
		Topic: p.kafkaTopicCardCommonRegister,
		Key:   string(cardCommon.CardNo),
		Value: []byte(b),
	}

	expectation := make(chan *ProducerError, 1)
	msg.expectation = expectation
	sp.producer.Input() <- msg

	if err := <-expectation; err != nil {
		return -1, -1, err.Err
	}
}
*/
