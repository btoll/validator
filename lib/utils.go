package lib

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
)

// Remove the unique suffix that was added by the migration tool to avoid the GVK error.
// https://github.com/kubernetes-sigs/kustomize/issues/4832#issuecomment-1344574518
var reServiceName = regexp.MustCompile(`^([a-z\-]*)-[a-z]*$`)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !errors.Is(err, os.ErrNotExist)
}

func CreateBuildDir(dir string) error {
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func CreateBuildFile(fileName string) (*os.File, error) {
	f, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func GetFileContents(filename string) ([]byte, error) {
	if CheckFileExists(filename) {
		return os.ReadFile(filename)
	}
	return nil, errors.New(fmt.Sprintf("File `%s` does not exist", filename))
}

func GetProperServiceName(name string) string {
	substring := reServiceName.FindStringSubmatch(name)
	return substring[1]
}
