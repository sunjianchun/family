package util

import (
	"fmt"
	"os"
)

//Dealerr 统一处理错误函数
func Dealerr(err error, flag string) {
	if err != nil {
		fmt.Println(err)
		if flag == Return {
			return
		} else if flag == Exit {
			os.Exit(1)
		} else {
			panic(err)
		}
	}
}
