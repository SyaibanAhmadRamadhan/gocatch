package Jlog

import (
	"fmt"
	"sync"

	"github.com/rs/zerolog"
)

type ZerologHook struct {
	Notify string
	Level  zerolog.Level
	sync.WaitGroup
}

func (z *ZerologHook) Run(e *zerolog.Event, level zerolog.Level, message string) {
	if level > z.Level {
		z.Add(1)
		go func() {
			// event
			fmt.Println("send")
			z.Done()
		}()
	}
}
