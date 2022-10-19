package main

import (
	"runtime"
	"server/conf"
	"server/db"
	"server/battle"
	"server/game"
	"server/world"
	"server/gate"

	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	cpuNum := runtime.NumCPU()
	runtime.GOMAXPROCS(cpuNum)

	leaf.Run(
		db.Module,
		battle.Module,
		game.Module,
		world.Module,
		gate.Module,
	)
}
