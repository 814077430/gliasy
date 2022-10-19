package handler

import (
	"server/common"
	"server/db"
	"server/msg"
	"server/game/service"

	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/gate"
	"github.com/golang/protobuf/proto"
)

func C2S_Login(args []interface{}) {
	m := args[0].(*msg.C2S_Login)
	a := args[1].(gate.Agent)

	var cbMsg msg.S2C_Login
	var userData msg.UserData
	cbMsg.Code = proto.Int32(common.Error_SUCCES)

	var acc string = *(m.Account)
	uid, ok := service.PlayerMap.Acc2Uid[acc]
	if !ok {
		uid = service.PlayerMap.GenUid()
		var p *(service.Player) = new(service.Player)
		p.Init()
		service.PlayerMap.Acc2Uid[acc] = uid
		service.PlayerMap.Players[uid] = p
		service.PlayerMap.Onlines[uid] = p

		var info msg.G2D_CreatePlayer

		var userData msg.UserData
		userData.Account = &acc
		userData.Uid = &uid
		var roleData msg.RoleData
		roleData.Name = &(p.RoleData.Name)
		userData.RoleData = &roleData

		info.Data = &userData
		db.ChanRPC.Go("G2D_CreatePlayer", &msg.G2D_CreatePlayer{}, info, Cb_G2D_CreatePlayer)
	}
	common.AddCon(uid, a)

	userData.Account = &acc
	userData.Uid = &uid
	cbMsg.Data = &userData
	common.Cb(a, &cbMsg)
	return
}

func Cb_G2D_CreatePlayer(code int32){
	log.Debug("Cb_G2D_CreatePlayer %v:", code)
}