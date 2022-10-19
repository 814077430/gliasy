package msg

import (
	"github.com/name5566/leaf/network/protobuf"
)

var Processor = protobuf.NewProcessor()

func init() {
	Processor.Register(&C2S_Login{})
	Processor.Register(&C2S_Logout{})

	Processor.Register(&S2C_Login{})
	Processor.Register(&S2C_Logout{})
}
