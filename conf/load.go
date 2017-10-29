package conf

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"path"
	"strings"
	"unicode"

	"family/util"
)

var (
	bEmpty  = []byte{}
	bEquel  = []byte{' ', '=', ' '}
	bQuote  = []byte{'"'}
	bSquare = []byte{'['}
)

func LoadBaseConfig(filename string) {
	file, err := os.Open(filename)
	util.Dealerr(err, util.Return)

	BC.Lock()
	defer BC.Unlock()
	defer file.Close()

	buf := bufio.NewReader(file)
	modu := ""
	offset := int64(1)
	for {
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		if bytes.Equal(line, bEmpty) {
			continue
		}
		offset += int64(len(line))
		if bytes.HasPrefix(line, bSquare) {
			line = bytes.TrimLeft(line, "[")
			line = bytes.TrimRight(line, "]")
			line = bytes.TrimLeftFunc(line, unicode.IsSpace)
			line = bytes.TrimRightFunc(line, unicode.IsSpace)
			modu = string(line)
		}
		if bytes.Contains(line, bEquel) {
			val := bytes.SplitN(line, bEquel, 2)
			if bytes.HasPrefix(val[1], bQuote) {
				val[1] = bytes.Trim(val[1], `"`)

				key := strings.TrimSpace(string(val[0]))
				if BC.Data[modu] == nil {
					temp := make(map[string]string)
					temp[key] = strings.TrimSpace(string(val[1]))
					BC.Data[modu] = temp
				} else {
					BC.Data[modu][key] = strings.TrimSpace(string(val[1]))
				}
			}
		}
		BC.Offset = offset
	}

}

func init() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		wd, _ := os.Getwd()
		configPath = path.Join(wd, "config/config.ini")
		LoadBaseConfig(configPath)
	}
}
