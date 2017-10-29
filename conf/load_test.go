package conf

import (
	"fmt"
	"testing"
)

func Test_Load(t *testing.T) {
	LoadBaseConfig("/Users/sunjianchun/work/code/go/src/family/config/config.ini")
	fmt.Printf("%v\n", BC.Offset)
}
