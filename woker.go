package loadbalancer

type Worker struct {
	requests chan Request
	pending  int
	index    int
}

func (w *Worker) work(done chan *Worker) {
	for {
		req := <-w.requests
		req.C <- req.Fn()
		done <- w
	}
}
