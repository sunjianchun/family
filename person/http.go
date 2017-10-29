package person

import (
	"encoding/json"
	"family/conf"
	"family/db"
	"family/util"
	"net/http"
	"os/exec"
	"strconv"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	ids := c.DefaultQuery("ids", "")
	name := c.DefaultQuery("name", "")

	var persons Persons
	var sql string = ""
	if ids != "" {
		sql += "id in (" + ids + ");"
	}
	if name != "" {
		sql += "name like'%" + name + "%';"
	}
	if sql != "" {
		sql = "select * from person where " + sql
		newDB := db.NewDB(sql, &persons)
		newDB.Do(conf.Query)
		c.JSON(http.StatusOK, persons)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "未找到族人"})
	}

}
func Get(c *gin.Context) {
	id := c.DefaultQuery("id", "")

	var person Person
	var sql string = ""
	if id != "" {
		sql += "select * from person where id=" + id + ";"
	}
	if sql != "" {
		newDB := db.NewDB(sql, &person)
		newDB.Do(conf.Query)
		c.JSON(http.StatusOK, person)
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
	script := "/Users/sunjianchun/work/code/go/src/family/script/readexcel.py"
	out, err := exec.Command(script, path).Output()
	util.Dealerr(err, util.Return)

	person := &Person{}

	var f interface{}
	_ = json.Unmarshal([]byte(string(out)), &f)
	m := f.(map[string]interface{})
	for k, v := range m {
		switch k {
		case "name":
			person.Name = v.(string)
		case "age":
			temp, err := strconv.Atoi(v.(string))
			util.Dealerr(err, util.Return)
			person.Age = temp
		case "children":
			person.Children = v.(string)
		case "birthday":
			person.Birthday = v.(string)
		case "brothers":
			person.Brothers = v.(string)
		case "phone":
			temp, err := strconv.ParseInt(v.(string), 10, 64)
			util.Dealerr(err, util.Return)
			person.Phone = temp
		case "compatriotRank":
			temp, err := strconv.Atoi(v.(string))
			util.Dealerr(err, util.Return)
			person.CompatriotRank = temp
		case "dad":
			person.Dad = v.(string)
		case "mom":
			person.Mom = v.(string)
		case "fellowRank":
			temp, err := strconv.Atoi(v.(string))
			util.Dealerr(err, util.Return)
			person.FellowRank = temp
		case "selfIntroduce":
			person.SelfIntroduce = v.(string)
		case "sisters":
			person.Sisters = v.(string)
		case "spouseIntroduce":
			person.SpouseIntroduce = v.(string)
		}
	}
	sql := "insert into person values(" + person.Password + ", " + person.FellowRank + ", " + person.CompatriotRank + ", " + person.Phone + ", " + person.IDCard + ", " + person.Age + ", " + person.Sex + ", " + person.Birthday + ", " + person.SelfImageURL + ", " + person.Status + ", " + person.SelfIntroduce + ", " + person.SpouseIntroduce + ", " + person.Dad + ", " + person.Mom + ", " + person.Remark + ", " + person.Password + ", " + person.Brothers + ", " + person.Sisters + ", " + person.Children + ", " + person.Generations + ");"
	c.JSON(http.StatusOK, *person)
}
