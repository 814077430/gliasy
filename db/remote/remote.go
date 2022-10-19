package remote
import (
	"strconv"

	"server/common"
	"server/msg"
	"server/db/service"

	"github.com/name5566/leaf/log"
	"github.com/golang/protobuf/proto"
)

func G2D_PlayerIncrId(args []interface{}) {
	m := args[1].(msg.G2D_PlayerIncrId)

	sql := "update t_general set IncrId = "+ strconv.Itoa((int)(m.GetIncrId())) + " where id = 1;";
	_, err := service.Db.Exec(sql)
	if err != nil {
		log.Debug("G2D_PlayerIncrId update error:", err, m.GetIncrId())
		return
	}
}

func G2D_CreatePlayer(args []interface{}) {
	m := args[1].(msg.G2D_CreatePlayer)
	cb := args[2].(func(int32))

	acc := m.GetData().GetAccount()
	uid := strconv.Itoa((int)(m.GetData().GetUid()))
	roleData, _ := proto.Marshal(m.GetData().GetRoleData())

	sql := "insert into t_user set " +
			"Acc = ?, " +
			"Uid = ?, " +
			"RoleData = ?" +
			";"

	_, err := service.Db.Exec(sql, acc, uid, roleData)
	if err != nil {
		log.Debug("G2D_CreatePlayer error: %v", err)
		cb(common.Error_CreatePlayerFailed)
	} else {
		cb(common.Error_SUCCES)
	}
}