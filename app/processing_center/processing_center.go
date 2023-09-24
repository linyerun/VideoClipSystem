package processing_center

import (
	"VideoClipSystem/app/global"
	"math/rand"
	"time"
)

var (
	chs []chan func()
)

func HandleTask(f func()) {
	rand.Seed(time.Now().UnixMilli())
	chs[rand.Uint64()%uint64(global.ProcessingCenterPoolSize)] <- f
}

func Init() {
	chs = make([]chan func(), global.ProcessingCenterPoolSize)
	for i := uint(0); i < global.ProcessingCenterPoolSize; i++ {
		chs[i] = make(chan func(), global.ProcessingCenterChanLen)
		go processHandler(chs[i])
	}
}

func processHandler(ch chan func()) {
	for {
		handle := <-ch
		handle()
	}
}
