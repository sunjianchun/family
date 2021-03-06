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
	ID              int64  `json:"id"`
	Name            string `json:"name" form:"name" binding:"required"`
	Password        string `json:"-" form:"password"`
	FellowRank      string `json:"fellowRank" form:"fellowRank"`
	Sex             string `json:"sex" form:"sex"`
	CompatriotRank  string `json:"compatriotRank" form:"compatriotRank"`
	Phone           string `json:"phone" form:"phone"`
	IDCard          string `json:"idCard" form:"idCard"`
	Age             string `json:"age" form:"age"`
	Birthday        string `json:"birthday" form:"birthday"`
	SelfImageURL    string `json:"selfImageURL" form:"selfImageURL"`
	SelfIntroduce   string `json:"selfIntroduce" form:"selfIntroduce"`
	SpouseImageURL  string `json:"spouseImageURL" form:"spouseImageURL"`
	SpouseIntroduce string `json:"spouseIntroduce" form:"spouseIntroduce"`
	DadID           string `json:"dadID" form:"dadID"`
	Dad             string `json:"dad" form:"dad"`
	Mom             string `json:"mom" form:"mom"`
	Brothers        string `json:"brothers" form:"brothers"`
	Sisters         string `json:"sisters" form:"sisters"`
	Children        string `json:"children" form:"children"`
	Status          bool   `json:"status" form:"status"`
	Generations     string `json:"generations" form:"generations"`
	Remark          string `json:"remark" form:"remark"`
	Parents         string `json:"parents" form:"parents"`
	Audit           string `json:"audit" form:"audit"`
	Email           string `json:"email" form:"email"`
}

var (
	DLock sync.Mutex
	BC    = &BaseConfig{
		string(""),
		make(map[string]map[string]string),
		int64(1),
		sync.RWMutex{},
	}
	Query   = "query"
	Insert  = "insert"
	Update  = "update"
	Delete  = "delete"
	Drop    = "drop"
	Persons = []*Person{}
)
