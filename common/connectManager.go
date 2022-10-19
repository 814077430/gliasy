package common

import (
	"github.com/name5566/leaf/gate"
)

type ConnectManager struct{
	PlayerCons map[int32](gate.Agent)
}

func ConInit(){
	conects.PlayerCons = make(map[int32](gate.Agent))
}

func Cb(_a gate.Agent, msg interface{}) {
	_a.WriteMsg(msg)
}

func PushMsg(uid int32, msg interface{}) {
	conects.PlayerCons[uid].WriteMsg(msg)
}

func BroadcastMsg(msg interface{}) {
	for a := range conects.PlayerCons {
		PushMsg(a, msg)
	}
}

func AddCon(uid int32, _a gate.Agent){
	conects.PlayerCons[uid] = _a
}

func DelCon(uid int32){
	delete(conects.PlayerCons, uid)
}

var (
	conects = new(ConnectManager)
)