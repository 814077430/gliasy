package gate

import (
	"server/game"
	"server/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.C2S_Login{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_Logout{}, game.ChanRPC)
}
