package snowflake

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// SnowflakeSeqGenerator snowflake gen ids
// ref: https://en.wikipedia.org/wiki/Snowflake_ID

var (
	// set the beginning time
	epoch = time.Date(2024, time.January, 01, 00, 00, 00, 00, time.UTC).UnixMilli()
)

const (
	timestampBits    = 41 // timestamp occupancy bits
	dataCenterIdBits = 5  // dataCenterId occupancy bits
	workerIdBits     = 5  // workerId occupancy bits
	seqBits          = 12 // sequence occupancy bits

	// timestamp max value, just like 2^41-1 = 2199023255551
	timestampMaxValue = -1 ^ (-1 << timestampBits)
	// dataCenterId max value, just like 2^5-1 = 31
	dataCenterIdMaxValue = -1 ^ (-1 << dataCenterIdBits)
	// workId max value, just like 2^5-1 = 31
	workerIdMaxValue = -1 ^ (-1 << workerIdBits)
	// sequence max value, just like 2^12-1 = 4095
	seqMaxValue = -1 ^ (-1 << seqBits)

	workIdShift       = 12 // number of workId offsets (seqBits)
	dataCenterIdShift = 17 // number of dataCenterId offsets (seqBits + workerIdBits)
	timestampShift    = 22 // number of timestamp offsets (seqBits + workerIdBits + dataCenterIdBits)

	defaultInitValue = 0
)

type SnowflakeSeqGenerator struct {
	mu           *sync.Mutex
	timestamp    int64
	dataCenterId int64
	workerId     int64
	sequence     int64
}

// NewSnowflakeSeqGenerator initiates the snowflake generator
func NewSnowflakeSeqGenerator(dataCenterId, workId int64) (r *SnowflakeSeqGenerator, err error) {
	if dataCenterId < 0 || dataCenterId > dataCenterIdMaxValue {
		err = fmt.Errorf("dataCenterId should between 0 and %d", dataCenterIdMaxValue-1)
		return
	}

	if workId < 0 || workId > workerIdMaxValue {
		err = fmt.Errorf("workId should between 0 and %d", dataCenterIdMaxValue-1)
		return
	}

	return &SnowflakeSeqGenerator{
		mu:           new(sync.Mutex),
		timestamp:    defaultInitValue - 1,
		dataCenterId: dataCenterId,
		workerId:     workId,
		sequence:     defaultInitValue,
	}, nil
}

// GenerateId timestamp + dataCenterId + workId + sequence
func (S *SnowflakeSeqGenerator) GenerateId(entity string, ruleName string) string {
	S.mu.Lock()
	defer S.mu.Unlock()

	now := time.Now().UnixMilli()
	//timestamp存储的是上一次的计算时间，如果当前的时间比上一次的时间还要小，那么说明发生了时钟回拨，
	//那么此时我们不进行生产id，并且记录错误日志。
	if S.timestamp > now { // Clock callback
		log.Fatalf("Clock moved backwards. Refusing to generate ID, last timestamp is %d, now is %d", S.timestamp, now)
		return ""
	}
	//如果时间相等的话，那就说明这是在 同一毫秒时间戳内生成的 ，
	//那么就进行seq的自旋，在这同一毫秒内最多生成 4095 个。
	if S.timestamp == now { // generate multiple IDs in the same millisecond, incrementing the sequence number to prevent conflicts
		S.sequence = (S.sequence + 1) & seqMaxValue
		if S.sequence == 0 { // sequence overflow, waiting for next millisecond
			// sequence overflow, waiting for next millisecond
			for now <= S.timestamp {
				now = time.Now().UnixMilli()
			}
		}
	} else { // initialized sequences are used directly at different millisecond timestamps
		S.sequence = defaultInitValue //那么如果是不在同一毫秒内的话，seq 直接用初始值就好了
	}
	//如果超过了69年，也就是时间戳超过了69年，也不能再继续生成了
	tmp := now - epoch
	if tmp > timestampMaxValue {
		log.Fatalf("epoch should between 0 and %d", timestampMaxValue-1)
		return ""
	}
	//记录这一次的计算时间，这样就可以和下一次的生成的时间做对比
	S.timestamp = now
	// combine the parts to generate the final ID and convert the 64-bit binary to decimal digits.
	r := (tmp)<<timestampShift |
		(S.dataCenterId << dataCenterIdShift) |
		(S.workerId << workIdShift) |
		(S.sequence)
	return fmt.Sprintf("%d", r)
}
