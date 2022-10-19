package handler

import (
	"server/common"
	"server/msg"
	"server/game/service"

	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/gate"
	"github.com/golang/protobuf/proto"
)

func C2S_Logout(args []interface{}) {
	m := args[0].(*msg.C2S_Login)
	a := args[1].(gate.Agent)

	var cbMsg msg.S2C_Logout
	cbMsg.Code = proto.Int32(common.Error_SUCCES)

	var acc string = *(m.Account)
	uid, ok := service.PlayerMap.Acc2Uid[acc]
	if !ok {
		log.Debug("C2S_Logout player not online acc: %s", acc)
		cbMsg.Code = proto.Int32(common.Error_PlayerNotOnline)
		common.Cb(a, &cbMsg)
		return
	}
	var p *(service.Player)= service.PlayerMap.Players[uid]
	if p == nil{
		log.Debug("C2S_Logout player not online acc: %s", acc)
		cbMsg.Code = proto.Int32(common.Error_PlayerNotOnline)
		common.Cb(a, &cbMsg)
		return
	}
	p.Online = false
	delete(service.PlayerMap.Onlines, uid)
	common.DelCon(uid)
	common.Cb(a, &cbMsg)
	return
}