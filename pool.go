package loadbalancer

type Pool []*Worker

func (p Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

func (p Pool) Len() int {
	return len(p)
}

func (p Pool) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *Pool) Pop() interface{} {
	old := *p
	n := len(old)
	w := old[n-1]
	*p = old[0 : n-1]
	return w
}

func (p *Pool) Push(x interface{}) {
	*p = append(*p, x.(*Worker))
}
