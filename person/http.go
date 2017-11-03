package person

import (
	"encoding/json"
	"family/conf"
	"family/db"
	"family/script"
	"family/util"
	"net/http"
	"os/exec"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
		sql = "select * from person where " + sql + " order by birthday;"
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

func Import(c *gin.Context) {
	path := c.DefaultQuery("path", "")
	if path == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "输入有误"})
		return
	}
	script := "/Users/sunjianchun/work/code/go/src/family/script/readexcelsjc.py"
	out, err := exec.Command(script, path).Output()
	util.Dealerr(err, util.Return)

	person := make(map[string]interface{})

	var f interface{}
	_ = json.Unmarshal([]byte(string(out)), &f)
	m := f.(map[string]interface{})
	for k, v := range m {
		switch k {
		case "name":
			person["name"] = v.(string)
		case "age":
			age := v.(string)
			if age != "" {
				temp, err := strconv.Atoi(age)
				util.Dealerr(err, util.Return)
				person["age"] = temp
			}
		case "children":
			person["children"] = v.(string)
		case "birthday":
			person["birthday"] = v.(string)
		case "brothers":
			person["brothers"] = v.(string)
		case "phone":
			person["phone"] = v.(string)
		case "compatriotRank":
			compatriotRank := v.(string)
			if compatriotRank != "" {
				temp, err := strconv.Atoi(compatriotRank)
				util.Dealerr(err, util.Return)
				person["compatriotRank"] = temp
			}
		case "dad":
			person["dad"] = v.(string)
		case "mom":
			person["mom"] = v.(string)
		case "fellowRank":
			fellowRank := v.(string)
			if fellowRank != "" {
				temp, err := strconv.Atoi(v.(string))
				util.Dealerr(err, util.Return)
				person["fellowRank"] = temp
			}
		case "selfIntroduce":
			person["selfIntroduce"] = v.(string)
		case "sisters":
			person["sisters"] = v.(string)
		case "spouseIntroduce":
			person["spouseIntroduce"] = v.(string)
		}
	}
	sql := "insert into person(name, age, children, birthday, brothers, phone, compatriotRank, dad, mom, fellowRank, selfIntroduce, sisters, spouseIntroduce) values (?,?,?,?,?,?,?,?,?,?,?,?,?);"
	newdb := db.NewDB(sql)
	newdb.Do(conf.Insert, person["name"], person["age"], person["children"], person["birthday"], person["brothers"], person["phone"], person["compatriotRank"], person["dad"], person["mom"], person["fellowRank"], person["selfIntroduce"], person["sisters"], person["spouseIntroduce"])
	c.JSON(http.StatusOK, person)
}

func Flush(c *gin.Context) {
	go script.TrimNameSpace()
	go script.TrimDadFieldSpace()
	res := script.ReplaceDadNameToID()
	c.JSON(http.StatusNotFound, res)
}

func GetChildren(c *gin.Context) {
	name := c.DefaultQuery("name", "")

	if name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "name不能为空"})
		return
	}
	result := script.FindChildren(name)
	c.JSON(http.StatusOK, result)
}

func GetAllPosterity(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	if name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "name不能为空"})
		return
	}
	result := script.Tree(name)
	c.JSON(http.StatusOK, result)
}
