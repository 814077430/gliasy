package service

import (
	"server/common"
)

type entity struct{
	id int32
	typ int32
	x int32
	y int32
}

type Player struct{
	b entity
	uid int32
}

type Castle struct{
	b entity
	uid int32
	name string
	uuid int32
	uname string
}

type Monster struct {
	b entity
}

type Resource struct {
	b entity
	uid int32
}

type WorldEntityManager struct{
	EntityIncrId int32
	Players map[int32](*Player)
	Castles map[int32](*Castle)
	Monsters map[int32](*Monster)
	Resources map[int32](*Resource)
}

func (we *WorldEntityManager) Init(){
	we.EntityIncrId = 0;
	we.Players = make(map[int32](*Player))
	we.Castles = make(map[int32](*Castle))
	we.Monsters = make(map[int32](*Monster))
	we.Resources = make(map[int32](*Resource))
}

func (we *WorldEntityManager) GenUuid() int32{
	we.EntityIncrId++
	return common.Const_BaseEntityId + we.EntityIncrId
}

var(
	WorldEntity = new(WorldEntityManager)
)
