package service

import (
	"server/common"
	"server/db"
	"server/msg"

	"github.com/golang/protobuf/proto"
)

type Player struct {
	Acc string
	Uid int32
	Online bool

	RoleData PlayerRoleManager
}

func (p *Player) Init(){
	p.Online = true
	p.RoleData.Name = "aliasy"
}

type PlayerManager struct {
	Acc2Uid map[string]int32
	Players map[int32](*Player)
	Onlines map[int32](*Player)
	PlayerIncrId int32
	ServerStartTime int32
	IsServerStart bool
}

func (pm *PlayerManager) Init(){
	pm.Acc2Uid = make(map[string]int32)
	pm.Players = make(map[int32](*Player))
	pm.Onlines = make(map[int32](*Player))
	common.ConInit()
}

func (pm *PlayerManager) Tick(now uint64){

}

func (pm *PlayerManager) GenUid() int32{
	pm.PlayerIncrId++

	data := msg.G2D_PlayerIncrId{
		IncrId : proto.Int32(pm.PlayerIncrId),
	}

	db.ChanRPC.Go("G2D_PlayerIncrId", &msg.G2D_PlayerIncrId{}, data)
	return common.Const_BasePlayerId + pm.PlayerIncrId;
}

var (
	PlayerMap = new(PlayerManager)
)