package time

type TimeHeap []TimePoint

func (th TimeHeap) Len() int {
	return len(th)
}

func (th TimeHeap) Less(i, j int) bool {
	return th[i].Point.Before(th[j].Point)
}

func (th TimeHeap) Swap(i, j int) {
	th[i], th[j] = th[j], th[i]
}

func (th TimeHeap) Peek() TimePoint {
	return th[0]
}

func (th *TimeHeap) Push(x interface{}) {
	tp := x.(TimePoint)
	*th = append(*th, tp)
}

func (th *TimeHeap) Pop() interface{} {
	old := *th
	n := len(old)
	tp := old[n-1]
	*th = old[0 : n-1]
	return tp
}
