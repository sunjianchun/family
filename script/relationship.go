package script

import (
	"family/conf"
	"family/db"
	"fmt"
	"regexp"
	"strings"
)

//TrimNameSpace 去除name里面的空格
func TrimNameSpace() {
	sql := "select name, id from person;"
	newdb := db.NewDB(sql)
	result := newdb.Do(conf.Query)
	for _, v := range result {
		flag := false
		tmp := ""
		for k1, v1 := range v {
			var id interface{}
			if k1 == "name" {
				str := v1.(string)
				if strings.Contains(str, " ") {
					flag = true
					list := strings.Split(str, " ")
					for _, v2 := range list {
						if v2 != "" {
							tmp += v2
						}
					}
				}
			} else if k1 == "id" && flag == true {
				id = v1
				sql = "update person set name=? where id=?"
				newdb = db.NewDB(sql)
				newdb.Do(conf.Update, tmp, id)
			}
		}
	}
}

//TrimDadFieldSpace 去除dad列里面的空格
func TrimDadFieldSpace() {
	sql := "select dad, id from person;"
	newdb := db.NewDB(sql)
	result := newdb.Do(conf.Query)
	for _, v := range result {
		flag := false
		tmp := ""
		for k1, v1 := range v {
			var id interface{}
			if k1 == "dad" {
				str := v1.(string)
				if strings.Contains(str, " ") {
					flag = true
					list := strings.Split(str, " ")
					for _, v2 := range list {
						if v2 != "" {
							tmp += v2
						}
					}
				}
			} else if k1 == "id" && flag == true {
				id = v1
				sql = "update person set dad=? where id=?"
				newdb = db.NewDB(sql)
				newdb.Do(conf.Update, tmp, id)
			}
		}
	}
}

//ReplaceDadNameToID 替换父亲的名字为id 便于后续计算关系
func ReplaceDadNameToID() []map[string]interface{} {
	var response = []map[string]interface{}{}
	sql := "select dad, id from person;"
	newdb := db.NewDB(sql)
	result := newdb.Do(conf.Query)
	count := 1
	for _, v := range result {
		count++
		flag := false
		var dadID interface{}
		var id interface{}
		content := make(map[string]interface{}, 1)
		str := ""
		for k1, v1 := range v {
			if k1 == "dad" {
				str = v1.(string)
				if m, _ := regexp.MatchString("^[0-9]+$", str); !m {
					sql = "select id from person where name=?;"
					newdb = db.NewDB(sql)
					res := newdb.Do(conf.Query, str)
					if len(res) == 1 {
						flag = true
						dadID = res[0]["id"]
					} else if len(res) > 1 {
						content[str] = "重名"
					} else if len(res) <= 0 {
						content[str] = "未找到"
					}
					response = append(response, content)
				}
			} else if k1 == "id" {
				id = v1
			}
		}

		if flag == true {
			sql = "update person set dad=? where id=?"
			newdb = db.NewDB(sql)
			newdb.Do(conf.Update, dadID, id)
			content[str] = "正确处理"
			response = append(response, content)
		}
	}
	fmt.Println(count, len(result))
	return response
}

func FindChildren(name string) []map[string]interface{} {
	var result = []map[string]interface{}{}
	sql := "select * from person where dad =? order by birthday;"
	newdb := db.NewDB(sql)
	result = newdb.Do(conf.Query, name)
	return result
}

func FindAllPosterity(name string) []map[string]interface{} {
	sql := "select name, selfIntroduce as bio, selfImageURL as image, children, dad, id from person where dad=? order by birthday;"
	newdb := db.NewDB(sql)
	result := newdb.Do(conf.Query, name)

	if len(result) > 0 {
		for _, value := range result {
			value["name"] = value["name"].(string) + "         "
			for innerkey, innervalue := range value {
				if innerkey == "id" {
					str := innervalue.(string)
					res := FindAllPosterity(str)
					value["children"] = res
				}
			}
		}
	} else {
		return nil
	}
	return result
}

func Tree(name string) map[string]interface{} {
	var result = make(map[string]interface{}, 1)
	sql := "select name, selfIntroduce as bio, selfImageURL as image, dad, id, children from person where id=?;"
	newdb := db.NewDB(sql)
	response := newdb.Do(conf.Query, name)
	if len(response) == 1 {
		result = response[0]
		result["name"] = result["name"].(string) + "        "
		result["children"] = FindAllPosterity(name)
	} else {
		return nil
	}

	return result
}
