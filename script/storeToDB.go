package script

import (
	"encoding/json"
	"family/conf"
	"family/util"
	"log"
	"os"
)

func LoadConfig(filename string) {
	file, err := os.Open(filename)
	defer file.Close()
	util.Dealerr(err, util.Return)
	if err != nil {
		log.Fatal("BaseConfig Not Found!")
	} else {
		err = json.NewDecoder(file).Decode(&conf.Persons)
		if err != nil {
			log.Fatal(err)
		}
	}
}
