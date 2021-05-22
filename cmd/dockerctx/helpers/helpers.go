package helpers

import (
	"io/ioutil"
	"os"
)

func getFilePath() string {
	homedir, _ := os.UserHomeDir()

	return homedir + "/.docker/dockerctx"
}

func SaveContext(prevContext string) error {
	file := getFilePath()
	data := []byte(prevContext)
	err := ioutil.WriteFile(file, data, 0644)

	return err
}

func ReadContext() (string, error) {
	file := getFilePath()
	b, err := ioutil.ReadFile(file)

	return string(b), err
}
