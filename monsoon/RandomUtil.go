package monsoon

import (
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"sync/atomic"
	crand "crypto/rand"
	"time"
)

const (
	Uppercase              = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lowercase              = "abcdefghijklmnopqrstuvwxyz"
	Alphabetic             = Uppercase + Lowercase
	Numeric                = "0123456789"
	Alphanumeric           = Alphabetic + Numeric
	Symbols                = "`" + `~!@#$%^&*()-_+={}[]|\;:"<>,./?`
	Hex                    = Numeric + "abcdef"
	MAXUINT32              = 4294967295
	DEFAULT_UUID_CNT_CACHE = 512
)

type Random struct {
	Prefix       string
	idGen        uint32
	internalChan chan uint32
}

func NewRandom() *Random {
	rand.Seed(time.Now().UnixNano())
	random := &Random{
		Prefix:       "",
		idGen:        0,
		internalChan: make(chan uint32, DEFAULT_UUID_CNT_CACHE),
	}
	random.startGen()
	return random
}

func NewRandomGen(prefix string, startValue uint32) *Random {
	rand.Seed(time.Now().UnixNano())
	random := &Random{
		Prefix:       prefix,
		idGen:        startValue,
		internalChan: make(chan uint32, DEFAULT_UUID_CNT_CACHE),
	}
	random.startGen()
	return random
}

//开启 goroutine, 把生成的数字形式的UUID放入缓冲管道
func (this *Random) startGen() {
	go func() {
		for {
			if this.idGen == MAXUINT32 {
				this.idGen = 1
			} else {
				this.idGen += 1
			}
			this.internalChan <- this.idGen
		}
	}()
}

//获取带前缀的字符串形式的UUID
func (this *Random) Get() string {
	idgen := <-this.internalChan
	return fmt.Sprintf("%s%d", this.Prefix, idgen)
}

//获取uint32形式的UUID
func (this *Random) GetUint32() uint32 {
	return <-this.internalChan
}

func (r *Random) String(length uint8, charsets ...string) string {
	charset := strings.Join(charsets, "")
	if charset == "" {
		charset = Alphanumeric
	}
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Int63()%int64(len(charset))]
	}
	return string(b)
}

func TimeUUID() string {
	u := [16]byte{}
	var timeBase = time.Date(1582, time.October, 15, 0, 0, 0, 0, time.UTC).Unix()
	var hardwareAddr []byte
	var clockSeq uint32

	utcTime := time.Now().In(time.UTC)
	t := uint64(utcTime.Unix()-timeBase)*10000000 + uint64(utcTime.Nanosecond()/100)
	u[0], u[1], u[2], u[3] = byte(t>>24), byte(t>>16), byte(t>>8), byte(t)
	u[4], u[5] = byte(t>>40), byte(t>>32)
	u[6], u[7] = byte(t>>56)&0x0F, byte(t>>48)

	clock := atomic.AddUint32(&clockSeq, 1)
	u[8] = byte(clock >> 8)
	u[9] = byte(clock)

	copy(u[10:], hardwareAddr)

	u[6] |= 0x10
	u[8] &= 0x3F
	u[8] |= 0x80

	var offsets = [...]int{0, 2, 4, 6, 9, 11, 14, 16, 19, 21, 24, 26, 28, 30, 32, 34}
	r := make([]byte, 36)
	for i, b := range u {
		r[offsets[i]] = Hex[b>>4]
		r[offsets[i]+1] = Hex[b&0xF]
	}
	r[8] = '-'
	r[13] = '-'
	r[18] = '-'
	r[23] = '-'
	return string(r)

}

//生成Guid字串
func Guid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		return ""
	}

	return Md5String(base64.URLEncoding.EncodeToString(b))
}
