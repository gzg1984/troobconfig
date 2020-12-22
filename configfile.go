package troobconfig

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

/*DefaultConfigFileName is used to set the DB host IP */
var DefaultConfigFileName = flag.String("config", "database.properties", "Config File Name or Path")

/*SearchRoot is used to search config file */
var SearchRoot = flag.String("search", "/opt", "Config File Name or Path")

func getConfigFile() string {

	if isExist(*DefaultConfigFileName) {
		return *DefaultConfigFileName
	}
	return searchInPath(*SearchRoot)
}

// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func isExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func searchInPath(path string) string {
	err := filepath.Walk(path, setGlobleConfigPath)
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
		return ""
	}
	return globalConfigPath
}

var globalConfigPath string //获取文件列表
func setGlobleConfigPath(path string, f os.FileInfo, err error) error {
	if f.Name() != *DefaultConfigFileName {
		return nil
	}

	var strRet string
	strRet, _ = os.Getwd()

	strRet += "/"

	if f == nil {
		return err
	}
	if f.IsDir() {
		return nil
	}

	strRet += path

	globalConfigPath = strRet

	fmt.Println(strRet) //list the file

	return nil
}
