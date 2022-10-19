package internal

import (
	"reflect"
	"server/msg"
	msgHandler "server/game/handler"
)

func init() {
	handler(&msg.C2S_Login{}, msgHandler.C2S_Login)
	handler(&msg.C2S_Logout{}, msgHandler.C2S_Logout)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}
