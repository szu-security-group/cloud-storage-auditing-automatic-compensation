package core

import (
	"fmt"
	"time"
)

type TimeController struct {
	timeQueue []string
	phase     string
	startTime int64
	endTime   int64
}

func (t *TimeController) Start(phase string) {
	if t.timeQueue == nil {
		t.timeQueue = []string{}
	}
	t.phase = phase
	t.startTime = time.Now().UnixMicro()
}
func (t *TimeController) End() {
	t.endTime = time.Now().UnixMicro()
	t.timeQueue = append(t.timeQueue, fmt.Sprintf("%v:%v ms", t.phase, t.endTime-t.startTime))
}

func (t *TimeController) Output() {
	s := ""
	for i := range t.timeQueue {
		s += fmt.Sprintf("%v|", t.timeQueue[i])
	}
	fmt.Println(s)
}
