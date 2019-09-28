package time

import (
	"container/heap"
	"time"
)

type TimeMachine struct {
	ticker   *time.Ticker
	timeHeap TimeHeap
	done     chan bool
}

func NewTimeMachine() *TimeMachine {
	var timeMachine TimeMachine
	heap.Init(&timeMachine.timeHeap)
	timeMachine.done = make(chan bool)

	return &timeMachine
}

func (tm *TimeMachine) Start() {
	tm.ticker = time.NewTicker(1000 * time.Millisecond)

	go func() {
		for {
			select {
			case <-tm.done:
				return
			case t := <-tm.ticker.C:
				for tm.timeHeap.Peek().Point.Before(t) {
					tp := heap.Pop(&tm.timeHeap).(TimePoint)
					tp.Action()

					if tp.IsRepetable() {
						heap.Push(&tm.timeHeap, TimePoint{
							Point:       time.Now().Add(tp.WaitTime),
							WaitTime:    tp.WaitTime,
							IsRepetable: tp.IsRepetable,
							Action:      tp.Action,
						})
					}
				}
			}
		}
	}()
}

func (tm *TimeMachine) AddTimePoint(tp TimePoint) {
	heap.Push(&tm.timeHeap, tp)
}

func (tm *TimeMachine) Stop() {
	tm.ticker.Stop()
	tm.done <- true
}
