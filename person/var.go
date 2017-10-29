package person

import (
	"family/conf"
	"family/db"
	"family/util"
)

type Persons []*conf.Person
type Person conf.Person

func (person *Person) Scan(rows db.Rows) {
	for rows.Next() {
		var dad string
		var brothers string
		var sisters string
		var children string
		err := rows.Scan(&person.ID, &person.Name, &person.Password, &person.FellowRank, &person.CompatriotRank, &person.Phone, &person.IDCard, &person.Age, &person.Sex, &person.Birthday, &person.SelfImageURL, &person.Status, &person.SelfIntroduce, &person.SpouseImageURL, &person.SpouseIntroduce, &dad, &person.Mom, &person.Remark, &brothers, &sisters, &children, &person.Generations)
		util.Dealerr(err, util.Return)

		person.Brothers = brothers
		person.Sisters = sisters
		person.Children = children

		person.Dad = dad
	}
}

func (p *Persons) Scan(rows db.Rows) {
	for rows.Next() {
		person := &conf.Person{}
		var dad string
		var brothers string
		var sisters string
		var children string
		err := rows.Scan(&person.ID, &person.Name, &person.Password, &person.FellowRank, &person.CompatriotRank, &person.Phone, &person.IDCard, &person.Age, &person.Sex, &person.Birthday, &person.SelfImageURL, &person.Status, &person.SelfIntroduce, &person.SpouseImageURL, &person.SpouseIntroduce, &dad, &person.Mom, &person.Remark, &brothers, &sisters, &children, &person.Generations)
		person.Brothers = brothers
		person.Sisters = sisters
		person.Children = children

		person.Dad = dad
		util.Dealerr(err, util.Return)
		*p = append(*p, person)
	}
}
