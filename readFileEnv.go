package readFileEnv

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func MakeEnvFileReader(dir string) func(file string) string {
	return func(file string) string {
		fileData, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", dir, file))
		if err != nil {
			panic(err.Error())
		}
		return strings.TrimSuffix(string(fileData), "\n")
	}
}


