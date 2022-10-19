package service

import (
	"math"
	"sync"
	"server/common"
	"github.com/name5566/leaf/log"
)

type tile struct {
	coordX int32
	coordY int32
	playerIds map[int32]int32
	pIdLock sync.RWMutex
	entityIds map[int32]int32
	eIdLock sync.RWMutex
}

func (t *tile) init(x int32, y int32){
	t.coordX = x
	t.coordY = y
	t.playerIds = make(map[int32]int32)
	t.entityIds = make(map[int32]int32)
}

func (t *tile) addObj(id int32, typ int32){
	t.eIdLock.Lock()
	defer t.eIdLock.Unlock()
	t.entityIds[id] = typ
}

func (t *tile) removeObj(id int32){
	t.eIdLock.Lock()
	defer t.eIdLock.Unlock()
	delete(t.entityIds, id)
}

func (t *tile) addPlayer(id int32){
	t.pIdLock.Lock()
	defer  t.pIdLock.Unlock()
	t.playerIds[id] = common.Const_Player
}

func (t *tile) removePlayer(id int32){
	t.pIdLock.Lock()
	defer  t.pIdLock.Unlock()
	delete(t.playerIds, id)
}

type tiles struct{
	long int32
	width int32
	tileLong int32
	tileWidth int32
	longLimit int32
	widthLimit int32
	aoi [][](*tile)
}

func (ts *tiles) Init(){
	ts.long = common.Const_Long
	ts.width = common.Const_Width
	ts.tileLong = common.Const_LTile
	ts.tileWidth = common.Const_WTile
	ts.longLimit = (int32)(math.Ceil((float64)(ts.long/ts.tileLong)))
	ts.widthLimit = (int32)(math.Ceil((float64)(ts.width/ts.tileWidth)))

	var i int32
	for i = 0; i < ts.longLimit; i++{
		var tmp [](*tile)
		var j int32
		for j = 0; j < ts.widthLimit; j++{
			var t *tile = new(tile)
			t.init(i, j)
			tmp = append(tmp, t)
		}
		ts.aoi = append(ts.aoi, tmp)
	}

	log.Debug("World aoi init success: %d %d", ts.longLimit, ts.widthLimit)
}

func (ts *tiles) CheckPos(x int32, y int32) bool{
	return x >=0 && x < ts.long && y >= 0 && y < ts.width
}

func (ts *tiles) AddObj(id int32, typ int32, x int32, y int32){
	if !ts.CheckPos(x, y){
		log.Debug("World aoi add obj pos error:%d %d", x, y)
		return
	}

	xx := (int32)(math.Floor((float64)(x/ts.tileLong)))
	yy := (int32)(math.Floor((float64)(y/ts.tileWidth)))

	var t *tile = ts.aoi[xx][yy]
	t.addObj(id, typ)

	switch typ {
		case common.Const_Castle:
			var c *Castle = new(Castle)
			c.b.id = WorldEntity.GenUuid()
			WorldEntity.Castles[c.b.id] = c
			break;
		case common.Const_Monster:
			var m *Monster = new(Monster)
			m.b.id = WorldEntity.GenUuid()
			WorldEntity.Monsters[m.b.id] = m
			break;
		case common.Const_Resource:
			var r *Resource = new(Resource)
			r.b.id = WorldEntity.GenUuid()
			WorldEntity.Resources[r.b.id] = r
			break;
	}
}

func (ts *tiles) RemoveObj(id int32, x int32, y int32){
	if !ts.CheckPos(x, y){
		log.Debug("World aoi remove obj pos error:%d %d", x, y)
		return
	}

	xx := (int32)(math.Floor((float64)(x/ts.tileLong)))
	yy := (int32)(math.Floor((float64)(y/ts.tileWidth)))

	var t *tile = ts.aoi[xx][yy]
	var typ int32
	typ = t.entityIds[id]
	t.removeObj(id)
	switch typ {
	case common.Const_Castle:
		delete(WorldEntity.Castles, id)
		break;
	case common.Const_Monster:
		delete(WorldEntity.Monsters, id)
		break;
	case common.Const_Resource:
		delete(WorldEntity.Resources, id)
		break;
	}

}

func (ts *tiles) AddPlayer(id int32, x int32, y int32){
	if !ts.CheckPos(x, y){
		log.Debug("World aoi add player pos error:%d %d", x, y)
		return
	}

	xx := (int32)(math.Floor((float64)(x/ts.tileLong)))
	yy := (int32)(math.Floor((float64)(y/ts.tileWidth)))

	var t *tile = ts.aoi[xx][yy]
	t.addObj(id, common.Const_Player)

	var p *Player = new(Player)
	p.b.id = WorldEntity.GenUuid()
	WorldEntity.Players[p.b.id] = p
}

func (ts *tiles) RemovePlayer(id int32, x int32, y int32){
	if !ts.CheckPos(x, y){
		log.Debug("World aoi remove player pos error:%d %d", x, y)
		return
	}

	xx := (int32)(math.Floor((float64)(x/ts.tileLong)))
	yy := (int32)(math.Floor((float64)(y/ts.tileWidth)))

	var t *tile = ts.aoi[xx][yy]
	t.removeObj(id)
	delete(WorldEntity.Players, id)
}

var(
	WorldAoi = new(tiles)
)