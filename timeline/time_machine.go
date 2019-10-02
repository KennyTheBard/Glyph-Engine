package time

import (
	"container/heap"
	"time"
)

type TimeMachine struct {
	ticker   *time.Ticker
	timeHeap *TimeHeap
	done     chan bool
}

var Timeline TimeMachine

func (tm *TimeMachine) Init() {
	tm.timeHeap = new(TimeHeap)
	heap.Init(tm.timeHeap)
	tm.done = make(chan bool)
}

func (tm *TimeMachine) Start() {
	tm.ticker = time.NewTicker(1000 * time.Millisecond)

	go func() {
		for {
			select {
			case <-tm.done:
				return
			case t := <-tm.ticker.C:
				for tm.timeHeap.Len() > 0 && tm.timeHeap.Peek().Point.Before(t) {
					tp := heap.Pop(tm.timeHeap).(TimePoint)
					tp.Action()

					if tp.IsRepetable() {
						heap.Push(tm.timeHeap, TimePoint{
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
	heap.Push(tm.timeHeap, tp)
}

func (tm *TimeMachine) Stop() {
	tm.ticker.Stop()
	tm.done <- true
}
