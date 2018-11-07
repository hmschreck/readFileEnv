package readFileEnv

import (
	"io/ioutil"
	"strings"
)

func MakeEnvFileReader(dir string) func(file string) string {
	return func(file string) string {
		fileData, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err.Error())
		}
		return strings.TrimSuffix(string(fileData), "\n")
	}
}


