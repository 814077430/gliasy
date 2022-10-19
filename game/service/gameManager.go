package service

import (
	"time"
)

var internal = time.Millisecond * 100

func Init(){
	PlayerMap.Init()
	time.AfterFunc(internal, Tick)
}

func Tick(){
	//tick
	now := uint64(time.Now().UnixNano())/1000
	PlayerMap.Tick(now);

	time.AfterFunc(internal, Tick)
}