package troobconfig

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

/*DefaultDBConfigFileName is used to set the DB host IP */
//var DefaultDBConfigFileName = flag.String("config", "database.properties", "Config File Name or Path")
var DefaultDBConfigFileName = flag.String("config", "jdbc.properties", "Config File Name or Path")

/*SearchRoot is used to search config file */
var SearchRoot = flag.String("search", "/opt", "Config File Name or Path")

func getDBConfigFile() string {
	flag.Parse()
	fmt.Printf("Searching for %v\n", *DefaultDBConfigFileName)
	if isExist(*DefaultDBConfigFileName) {
		return *DefaultDBConfigFileName
	}
	return searchDBConfigInPath(*SearchRoot)
}

// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func isExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func searchDBConfigInPath(path string) string {
	err := filepath.Walk(path, setGlobleConfigPath)
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
		return ""
	}
	return globalDBConfigPath
}

var globalDBConfigPath string //获取文件列表
func setGlobleConfigPath(path string, f os.FileInfo, err error) error {
	if f.Name() != *DefaultDBConfigFileName {
		return nil
	}

	globalDBConfigPath = path

	fmt.Println(path) //list the file

	return nil
}
