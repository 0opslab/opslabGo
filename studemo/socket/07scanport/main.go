package main

import (
	"container/list"
	"net"
	"strconv"
	"sync"
	"time"
	"sort"
	"fmt"
	"strings"
)

type Queue struct {
	data *list.List
}

func NewQueue() *Queue {
	q := new(Queue)
	q.data = list.New()
	return q
}

func (q *Queue) push(v interface{}) {
	defer lock.Unlock()
	lock.Lock()
	q.data.PushFront(v)
}

func ScanPort(ip string, portlist []int, q *Queue) {
	for i := 0; i < len(portlist); i++ {
		_, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(portlist[i]), time.Second*3)
		if err != nil {
			continue
		}
		q.push(portlist[i])
	}
	wg.Done()
}

var lock sync.Mutex
var wg sync.WaitGroup

func main() {
	q := NewQueue()
	pport := 1
	for i := 0; i < 5000; i++ {
		all_port := make([]int, 14)
		for j := 0; j < 14; j++ {
			pport++
			all_port[j] = pport
		}
		wg.Add(1)
		go ScanPort("127.0.0.1", all_port, q)
	}
	wg.Wait()
	res_port := []int{}
	for iter := q.data.Back(); iter != nil; iter = iter.Prev() {
		switch v := iter.Value.(type) {
		case int:
			res_port = append(res_port, v)
		}
	}
	sort.Ints(res_port)
	ss := strings.Replace(strings.Trim(fmt.Sprint(res_port), "[]"), " ", ",", -1)
	println("127.0.0.1 open port:", ss)
}
