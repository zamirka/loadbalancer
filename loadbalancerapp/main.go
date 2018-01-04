package main

import (
	"math/rand"
	"time"

	"github.com/zamirka/loadbalancer"
)

const nWorker = 10
const nRequester = 100

func main() {

	work := make(chan loadbalancer.Request)
	for i := 0; i < nRequester; i++ {
		go requester(work)
	}
	loadbalancer.NewBalancer(nWorker, nRequester).Balance(work)
}

func op() int {
	n := rand.Int63n(int64(time.Second))
	time.Sleep(time.Duration(nWorker * n))
	return int(n)
}

func requester(work chan loadbalancer.Request) {
	c := make(chan int)
	for {
		time.Sleep(time.Duration(rand.Int63n(int64(nWorker * 2 * time.Second))))
		work <- loadbalancer.Request{op, c}
		<-c
	}
}
