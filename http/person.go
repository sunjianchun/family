package http

import (
	"family/conf"
	"family/db"
	"family/script"
	"family/util"
	"fmt"
	"net/http"
	"os/exec"
	"regexp"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func NotNullReturn(c *gin.Context, name string, value string) {
	if value == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": name + "不能为空"})
		return
	}
}

func List(c *gin.Context) {
	ids := c.DefaultQuery("ids", "")
	name := c.DefaultQuery("name", "")

	var sql string = ""
	if ids != "" {
		sql += "id in (" + ids + ")"
	}
	if name != "" {
		if sql != "" {
			sql += " and name like'%" + name + "%'"
		} else {
			sql = "name like'%" + name + "%'"
		}
	}
	if sql != "" {
		sql = "select * from person where " + sql + " order by generations, birthday;"
		newDB := db.NewDB(sql)
		result := newDB.Do(conf.Query)
		if len(result) != 0 {
			c.JSON(http.StatusOK, result)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"message": "未找到族人"})
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "未找到族人"})
	}

}
func Get(c *gin.Context) {
	id := c.DefaultQuery("id", "")

	var sql string = ""
	if id != "" {
		sql += "select * from person where id=" + id + ";"
	}
	if sql != "" {
		newDB := db.NewDB(sql)
		result := newDB.Do(conf.Query)
		if len(result) == 1 {
			c.JSON(http.StatusOK, result[0])
		} else if len(result) <= 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "未找到族人"})
		} else {
			c.JSON(http.StatusOK, result)
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "未找到族人"})
	}
}

func Add(c *gin.Context) {
	var person conf.Person
	err := c.ShouldBind(&person)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	sql := "select name from person where id=?;"
	newdb := db.NewDB(sql)
	result := newdb.Do(conf.Query, person.DadID)

	if len(result) == 1 {
		person.Dad = result[0]["name"].(string)
	} else {
		person.Dad = ""
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	files := form.File["image"]

	for _, file := range files {
		if err := c.SaveUploadedFile(file, file.Filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
	}

	sql = "insert into person(name, fellowRank, sex, compatriotRank, phone, idCard, age, birthday, selfIntroduce, spouseIntroduce, dadID, dad, mom, brothers, sisters, children, status, remark, email, password) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
	newdb = db.NewDB(sql)
	result = newdb.Do(conf.Insert, person.Name, person.FellowRank, person.Sex, person.CompatriotRank, person.Phone, person.IDCard, person.Age, person.Birthday, person.SelfIntroduce, person.SpouseIntroduce, person.DadID, person.Dad, person.Mom, person.Brothers, person.Sisters, person.Children, person.Status, person.Remark, person.Email, util.Md5Encode("zaqzaqzaq"))
	if len(result) == 1 {
		fmt.Println("in")
		lastInsertId := result[0]["lastInsertId"]
		script.GetGeneration(person.DadID, 1, lastInsertId.(string))
	}
	c.JSON(http.StatusOK, gin.H{"message": "新增成功"})
}

func Import(c *gin.Context) {
	path := "/tmp/file.txt1"
	_, err := exec.Command("cat", path).Output()
	util.Dealerr(err, util.Return)
	script.LoadConfig(path)
	res := make([]map[string]string, 273)
	for _, p := range conf.Persons {
		dict := make(map[string]string, 1)
		p1, _ := regexp.Compile("([0-9]{2,4})[^0-9]+([0-9]{1,2})[^0-9]+([0-9]{1,2})")
		list := p1.FindAllStringSubmatch(p.Birthday, -1)
		compile := false
		if len(list) > 0 {
			if len(list[0]) == 4 {
				tmp := fmt.Sprintf("%s-%s-%s", list[0][1], list[0][2], list[0][3])
				p.Birthday = tmp
				compile = true
			}
		}
		p1, _ = regexp.Compile("([0-9]{2,4})[^0-9]+([0-9]{1,2})")
		list = p1.FindAllStringSubmatch(p.Birthday, -1)
		if len(list) > 0 {
			if len(list[0]) == 3 && compile == false {
				tmp := fmt.Sprintf("%s-%s-%s", list[0][1], list[0][2], "01")
				p.Birthday = tmp
				compile = true
			}
		}

		p1, _ = regexp.Compile("([0-9]{2,4})[^0-9]*")
		list = p1.FindAllStringSubmatch(p.Birthday, -1)
		if len(list) > 0 {
			if len(list[0]) == 2 && compile == false {
				tmp := fmt.Sprintf("%s-%s-%s", list[0][1], "01", "01")
				p.Birthday = tmp
				compile = true
			}
		}
		if compile == false {
			dict["pre"] = p.Birthday
			p.Birthday = ""
			dict["post"] = p.Birthday
			res = append(res, dict)

		}
		p1, _ = regexp.Compile("([0-9]{11,13})")
		list = p1.FindAllStringSubmatch(p.Remark, -1)
		if len(list) > 0 {
			if len(list[0]) == 2 {
				p.Phone = list[0][1]
			}
		}
		if p.Birthday == "" {
			p.Birthday = "2200-01-01"
		}
		p1, _ = regexp.Compile("父[:： ]*([^ ：: ]+)[ ]*[生]?母[:： ]*(.*)")
		list = p1.FindAllStringSubmatch(p.Parents, -1)
		if len(list) > 0 {
			if len(list[0]) == 3 {
				p.Dad = list[0][1]
				p.Mom = strings.Trim(list[0][2], " ")
			}
		}

		sql := "insert into person(name, age, children, birthday, brothers, phone, compatriotRank, dad, mom, fellowRank, selfIntroduce, sisters, spouseIntroduce) values (?,?,?,?,?,?,?,?,?,?,?,?,?);"
		newdb := db.NewDB(sql)
		newdb.Do(conf.Insert, p.Name, p.Age, p.Children, p.Birthday, p.Brothers, p.Phone, p.CompatriotRank, p.Dad, p.Mom, p.FellowRank, p.SelfIntroduce, p.Sisters, p.SpouseIntroduce)
		sql = "update person set birthday = NULL where birthday='2200-01-01';"
		newdb = db.NewDB(sql)
		newdb.Do(conf.Update)
	}
	c.JSON(http.StatusOK, conf.Persons)
}

func Flush(c *gin.Context) {
	go script.TrimNameSpace()
	go script.TrimDadFieldSpace()
	res := script.InsertDadID()
	c.JSON(http.StatusNotFound, res)
	go script.InsertGeneration()
}

func GetChildren(c *gin.Context) {
	id := c.DefaultQuery("id", "")

	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "id不能为空"})
		return
	}
	result := script.FindChildren(id)
	c.JSON(http.StatusOK, result)
}

func GetAllPosterity(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "id不能为空"})
		return
	}
	result := script.Tree(id)
	c.JSON(http.StatusOK, result)
}

func GetUserInfo(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userID")
	if userID == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "no auth"})
	} else {
		sql := "select * from person where id=?;"
		newdb := db.NewDB(sql)
		result := newdb.Do(conf.Query, userID)
		if len(result) > 0 {
			c.JSON(http.StatusOK, result[0])
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "not found"})
		}
	}
}
