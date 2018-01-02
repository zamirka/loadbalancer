package loadbalancer

import (
	"math/rand"
	"testing"
	"time"
)

const nWorker = 10
const nRequester = 100

func BenchmarkBalance(b *testing.B) {

	work := make(chan Request)
	for i := 0; i < nRequester; i++ {
		go requester(work)
	}
	NewBalancer(nWorker, nRequester).Balance(work)
}

func op() int {
	n := rand.Int63n(int64(time.Second))
	time.Sleep(time.Duration(nWorker * n))
	return int(n)
}

func requester(work chan Request) {
	c := make(chan int)
	for {
		time.Sleep(time.Duration(rand.Int63n(int64(nWorker * 2 * time.Second))))
		work <- Request{op, c}
		<-c
	}
}
