package gamedata

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/recordfile"
	"reflect"
)

func readRf(st interface{}) *recordfile.RecordFile {
	rf, err := recordfile.New(st)
	if err != nil {
		log.Fatal("%v", err)
	}
	fn := reflect.TypeOf(st).Name() + ".txt"
	err = rf.Read("gamedata/" + fn)
	if err != nil {
		log.Fatal("%v: %v", fn, err)
	}

	return rf
}

type Test struct {
	IndexInt int "index"
	IndexStr string "index"
	Number  int32
	Str      string
	Arr1     [2]int
	Arr2     [3][2]int
	Arr3     []int
	St       struct {
		Name string "name"
		Num  int    "num"
	}
	M map[string]int
}

var RfTest *recordfile.RecordFile = readRf(Test{})

/*
	r := gamedata.RfTest.Index(1)

	if r != nil {
		row := r.(*(gamedata.Test))
		_ = row
		log.Debug("%v %v %v", row.IndexInt, row.IndexStr, row.Number)
	}
*/