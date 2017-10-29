package conf

import (
	"sync"
)

//BaseConfig 数据库redis等信息
type BaseConfig struct {
	Filename string
	Data     map[string]map[string]string
	Offset   int64
	sync.RWMutex
}

//Person 数据
type Person struct {
	ID              int64       `json:"id"`
	Name            string      `json:"name" form:"name" binding:"required"`
	Password        string      `json:"-" form:"password" binding:"required"`
	FellowRank      int         `json:"fellowRank" form:"fellowRank"`
	Sex             string      `json:"sex" form:"sex"`
	CompatriotRank  int         `json:"compatriotRank" form:"compatriotRank"`
	Phone           int64       `json:"phone" form:"phone"`
	IDCard          int64       `json:"idCard" form:"idCard"`
	Age             int         `json:"age" form:"age"`
	Birthday        string      `json:"birthday" form:"birthday"`
	SelfImageURL    string      `json:"selfImageURL" form:"selfImageURL"`
	SelfIntroduce   string      `json:"selfIntroduce" form:"selfIntroduce"`
	SpouseImageURL  string      `json:"spouseImageURL" form:"spouseImageURL"`
	SpouseIntroduce string      `json:"spouseIntroduce" form:"spouseIntroduce"`
	Dad             interface{} `json:"dad" form:"dad"`
	Mom             string      `json:"mom" form:"mom"`
	Brothers        interface{} `json:"brothers" form:"brothers"`
	Sisters         interface{} `json:"sisters" form:"sisters"`
	Children        interface{} `json:"children" form:"children"`
	Status          bool        `json:"status" form:"status"`
	Generations     int         `json:"generations" form:"generations"`
	Remark          string      `json:"remark" form:"remark"`
}

var (
	DLock sync.Mutex
	BC    = &BaseConfig{
		string(""),
		make(map[string]map[string]string),
		int64(1),
		sync.RWMutex{},
	}
	Query = "query"
)
