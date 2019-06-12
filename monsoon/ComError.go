package monsoon

import (
	"fmt"
	"runtime/debug"
	"time"
)

const logFormat = "%v : %d - %s \n %s "

type ComError struct {
	time  time.Time
	code  int
	msg   string
	stack string
}

func (m *ComError) Error() string {
	return fmt.Sprintf(logFormat, m.time, m.code, m.msg, m.stack)
}

// 根据消息生成错误消息
func NewComError(strMsg string) *ComError {
	return &ComError{
		time:  time.Now(),
		code:  0,
		msg:   strMsg,
		stack: string(debug.Stack()),
	}
}
// 根据信息包装指定的错误消息
func NewComWithError(strMsg string, err error) *ComError {
	if (err != nil) {
		strMsg += " " + err.Error()
	}
	return &ComError{
		time:  time.Now(),
		code:  0,
		msg:   strMsg,
		stack: string(debug.Stack()),
	}
}
// 根据错误编码和消息生成错误消息
func NewComErrorCode(i int, strMsg string) *ComError {
	return &ComError{
		time:  time.Now(),
		code:  i,
		msg:   strMsg,
		stack: string(debug.Stack()),
	}
}
// 根据错误编码和消息包装指定错误消息
func NewComWithErrorCode(i int, strMsg string, err error) *ComError {
	if (err != nil) {
		strMsg += " " + err.Error()
	}
	return &ComError{
		time:  time.Now(),
		code:  i,
		msg:   strMsg,
		stack: string(debug.Stack()),
	}
}
