package kafka

import (
	"github.com/devexps/go-micro/v2/broker"
	kafkaGo "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	"github.com/segmentio/kafka-go/sasl/scram"
	"hash"
	"time"
)

type BalancerName string

const (
	LeastBytesBalancer    BalancerName = "LeastBytes"
	RoundRobinBalancer    BalancerName = "RoundRobin"
	HashBalancer          BalancerName = "Hash"
	ReferenceHashBalancer BalancerName = "ReferenceHash"
	Crc32Balancer         BalancerName = "CRC32Balancer"
	Murmur2Balancer       BalancerName = "Murmur2Balancer"
)

type ScramAlgorithm string

const (
	ScramAlgorithmSHA256 ScramAlgorithm = "SHA256"
	ScramAlgorithmSHA512 ScramAlgorithm = "SHA512"
)

///
/// Option
///

type retriesCountKey struct{}
type queueCapacityKey struct{}
type minBytesKey struct{}
type maxBytesKey struct{}
type maxWaitKey struct{}
type readLagIntervalKey struct{}
type heartbeatIntervalKey struct{}
type commitIntervalKey struct{}
type partitionWatchIntervalKey struct{}
type watchPartitionChangesKey struct{}
type sessionTimeoutKey struct{}
type rebalanceTimeoutKey struct{}
type retentionTimeKey struct{}
type startOffsetKey struct{}
type mechanismKey struct{}
type readerConfigKey struct{}
type writerConfigKey struct{}
type dialerConfigKey struct{}
type dialerTimeoutKey struct{}
type loggerKey struct{}
type errorLoggerKey struct{}
type enableLoggerKey struct{}
type enableErrorLoggerKey struct{}
type enableOneTopicOneWriterKey struct{}
type batchSizeKey struct{}
type batchTimeoutKey struct{}
type batchBytesKey struct{}
type asyncKey struct{}
type maxAttemptsKey struct{}
type readTimeoutKey struct{}
type writeTimeoutKey struct{}
type allowPublishAutoTopicCreationKey struct{}

// WithReaderConfig .
func WithReaderConfig(cfg kafkaGo.ReaderConfig) broker.Option {
	return broker.WithContextAndValue(readerConfigKey{}, cfg)
}

// WithWriterConfig .
func WithWriterConfig(cfg WriterConfig) broker.Option {
	return broker.WithContextAndValue(writerConfigKey{}, cfg)
}

// WithEnableOneTopicOneWriter .
func WithEnableOneTopicOneWriter(enable bool) broker.Option {
	return broker.WithContextAndValue(enableOneTopicOneWriterKey{}, enable)
}

// WithPlainMechanism .
func WithPlainMechanism(username, password string) broker.Option {
	mechanism := plain.Mechanism{
		Username: username,
		Password: password,
	}
	return broker.WithContextAndValue(mechanismKey{}, mechanism)
}

// WithScramMechanism .
func WithScramMechanism(algoName ScramAlgorithm, username, password string) broker.Option {
	var algo scram.Algorithm
	switch algoName {
	case ScramAlgorithmSHA256:
		algo = scram.SHA256
	case ScramAlgorithmSHA512:
		algo = scram.SHA512
	}
	mechanism, err := scram.Mechanism(algo, username, password)
	if err != nil {
		panic(err)
	}
	return broker.WithContextAndValue(mechanismKey{}, mechanism)
}

// WithDialer .
func WithDialer(cfg *kafkaGo.Dialer) broker.Option {
	return broker.WithContextAndValue(dialerConfigKey{}, cfg)
}

// WithDialerTimeout .
func WithDialerTimeout(tm time.Duration) broker.Option {
	return broker.WithContextAndValue(dialerTimeoutKey{}, tm)
}

// WithRetries .
func WithRetries(cnt int) broker.Option {
	return broker.WithContextAndValue(retriesCountKey{}, cnt)
}

// WithQueueCapacity .
func WithQueueCapacity(cap int) broker.Option {
	return broker.WithContextAndValue(queueCapacityKey{}, cap)
}

// WithMinBytes fetch.min.bytes
func WithMinBytes(bytes int) broker.Option {
	return broker.WithContextAndValue(minBytesKey{}, bytes)
}

// WithMaxBytes .
func WithMaxBytes(bytes int) broker.Option {
	return broker.WithContextAndValue(maxBytesKey{}, bytes)
}

// WithMaxWait fetch.max.wait.ms
func WithMaxWait(time time.Duration) broker.Option {
	return broker.WithContextAndValue(maxWaitKey{}, time)
}

// WithReadLagInterval .
func WithReadLagInterval(interval time.Duration) broker.Option {
	return broker.WithContextAndValue(readLagIntervalKey{}, interval)
}

// WithHeartbeatInterval .
func WithHeartbeatInterval(interval time.Duration) broker.Option {
	return broker.WithContextAndValue(heartbeatIntervalKey{}, interval)
}

// WithCommitInterval .
func WithCommitInterval(interval time.Duration) broker.Option {
	return broker.WithContextAndValue(commitIntervalKey{}, interval)
}

// WithPartitionWatchInterval .
func WithPartitionWatchInterval(interval time.Duration) broker.Option {
	return broker.WithContextAndValue(partitionWatchIntervalKey{}, interval)
}

// WithWatchPartitionChanges .
func WithWatchPartitionChanges(enable bool) broker.Option {
	return broker.WithContextAndValue(watchPartitionChangesKey{}, enable)
}

// WithSessionTimeout .
func WithSessionTimeout(timeout time.Duration) broker.Option {
	return broker.WithContextAndValue(sessionTimeoutKey{}, timeout)
}

// WithRebalanceTimeout .
func WithRebalanceTimeout(timeout time.Duration) broker.Option {
	return broker.WithContextAndValue(rebalanceTimeoutKey{}, timeout)
}

// WithRetentionTime .
func WithRetentionTime(time time.Duration) broker.Option {
	return broker.WithContextAndValue(retentionTimeKey{}, time)
}

// WithStartOffset .
func WithStartOffset(offset int64) broker.Option {
	return broker.WithContextAndValue(startOffsetKey{}, offset)
}

// WithMaxAttempts .
func WithMaxAttempts(cnt int) broker.Option {
	return broker.WithContextAndValue(maxAttemptsKey{}, cnt)
}

// WithLogger inject info logger
func WithLogger(l kafkaGo.Logger) broker.Option {
	return broker.WithContextAndValue(loggerKey{}, l)
}

// WithErrorLogger inject error logger
func WithErrorLogger(l kafkaGo.Logger) broker.Option {
	return broker.WithContextAndValue(errorLoggerKey{}, l)
}

// WithEnableLogger enable go-micro info logger
func WithEnableLogger(enable bool) broker.Option {
	return broker.WithContextAndValue(enableLoggerKey{}, enable)
}

// WithEnableErrorLogger enable go-micro error logger
func WithEnableErrorLogger(enable bool) broker.Option {
	return broker.WithContextAndValue(enableErrorLoggerKey{}, enable)
}

// WithBatchSize batch.size
// default：100
func WithBatchSize(size int) broker.Option {
	return broker.WithContextAndValue(batchSizeKey{}, size)
}

// WithBatchTimeout linger.ms
// default：10ms
func WithBatchTimeout(timeout time.Duration) broker.Option {
	return broker.WithContextAndValue(batchTimeoutKey{}, timeout)
}

// WithBatchBytes
// default：1048576 bytes
func WithBatchBytes(by int64) broker.Option {
	return broker.WithContextAndValue(batchBytesKey{}, by)
}

// WithAsync
// default：true
func WithAsync(enable bool) broker.Option {
	return broker.WithContextAndValue(asyncKey{}, enable)
}

// WithPublishMaxAttempts .
func WithPublishMaxAttempts(cnt int) broker.Option {
	return broker.WithContextAndValue(maxAttemptsKey{}, cnt)
}

// WithReadTimeout
// default：10s
func WithReadTimeout(timeout time.Duration) broker.Option {
	return broker.WithContextAndValue(readTimeoutKey{}, timeout)
}

// WithWriteTimeout
// default：10s
func WithWriteTimeout(timeout time.Duration) broker.Option {
	return broker.WithContextAndValue(writeTimeoutKey{}, timeout)
}

// WithAllowPublishAutoTopicCreation .
func WithAllowPublishAutoTopicCreation(enable bool) broker.Option {
	return broker.WithContextAndValue(allowPublishAutoTopicCreationKey{}, enable)
}

///
/// PublishOption
///

type messageHeadersKey struct{}
type messageKeyKey struct{}
type messageOffsetKey struct{}
type balancerKey struct{}
type balancerValue struct {
	Name       BalancerName
	Consistent bool
	Hasher     hash.Hash32
}

// WithHeaders .
func WithHeaders(headers map[string]interface{}) broker.PublishOption {
	return broker.WithPublishContextAndValue(messageHeadersKey{}, headers)
}

// WithMessageKey .
func WithMessageKey(key []byte) broker.PublishOption {
	return broker.WithPublishContextAndValue(messageKeyKey{}, key)
}

// WithMessageOffset .
func WithMessageOffset(offset int64) broker.PublishOption {
	return broker.WithPublishContextAndValue(messageOffsetKey{}, offset)
}

// WithLeastBytesBalancer .
func WithLeastBytesBalancer() broker.PublishOption {
	return broker.WithPublishContextAndValue(balancerKey{},
		&balancerValue{
			Name: LeastBytesBalancer,
		},
	)
}

// WithRoundRobinBalancer .
func WithRoundRobinBalancer() broker.PublishOption {
	return broker.WithPublishContextAndValue(balancerKey{},
		&balancerValue{
			Name: RoundRobinBalancer,
		},
	)
}

// WithHashBalancer .
func WithHashBalancer(hasher hash.Hash32) broker.PublishOption {
	return broker.WithPublishContextAndValue(balancerKey{},
		&balancerValue{
			Name:   HashBalancer,
			Hasher: hasher,
		},
	)
}

// WithReferenceHashBalancer .
func WithReferenceHashBalancer(hasher hash.Hash32) broker.PublishOption {
	return broker.WithPublishContextAndValue(balancerKey{},
		&balancerValue{
			Name:   ReferenceHashBalancer,
			Hasher: hasher,
		},
	)
}

// WithCrc32Balancer .
func WithCrc32Balancer(consistent bool) broker.PublishOption {
	return broker.WithPublishContextAndValue(balancerKey{},
		&balancerValue{
			Name:       Crc32Balancer,
			Consistent: consistent,
		},
	)
}

// WithMurmur2Balancer .
func WithMurmur2Balancer(consistent bool) broker.PublishOption {
	return broker.WithPublishContextAndValue(balancerKey{},
		&balancerValue{
			Name:       Murmur2Balancer,
			Consistent: consistent,
		},
	)
}

///
/// SubscribeOption
///

type autoSubscribeCreateTopicKey struct{}
type autoSubscribeCreateTopicValue struct {
	Topic             string
	NumPartitions     int
	ReplicationFactor int
}

func WithSubscribeAutoCreateTopic(topic string, numPartitions, replicationFactor int) broker.SubscribeOption {
	return broker.WithSubscribeContextAndValue(autoSubscribeCreateTopicKey{},
		&autoSubscribeCreateTopicValue{
			Topic:             topic,
			NumPartitions:     numPartitions,
			ReplicationFactor: replicationFactor,
		},
	)
}
