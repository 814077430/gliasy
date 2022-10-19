package service

import (
	"server/conf"

	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/db/mongodb"
)

type Person struct {
	Uid		int
	Name  string
	Phone string
}

var mongoDB *mongodb.DialContext

func initMongo() {
	if conf.Server.DBMaxConnNum <= 0 {
		conf.Server.DBMaxConnNum = 100
	}

	db, err := mongodb.Dial(conf.Server.DBUrl, conf.Server.DBMaxConnNum)
	if err != nil {
		log.Fatal("dial mongodb error: %v", err)
	}
	mongoDB = db

	s := db.Ref()
	defer db.UnRef(s)
	client := s.DB("game").C("user")
	_ = client
}

func MongoDBDestroy() {
	mongoDB.Close()
	mongoDB = nil
}