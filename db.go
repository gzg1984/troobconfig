package troobconfig

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"strings"
	"sync"

	"gopkg.in/ini.v1"
)

func analyzeURL(url string) {
	/*jdbc:mysql://localhost:3306/lxr_prod?useUnicode=true&characterEncoding=utf-8*/
	/*url=jdbc:mysql://127.0.0.1/lxr_db*/

	/*step 1: sep from ? */
	cutArgs := strings.Split(url, "?")
	if len(cutArgs) < 1 {
		return
	}

	/* step 2: get target and DB name */
	mysqlBase := cutArgs[0]

	var targetAndDB string
	n, err := fmt.Sscanf(mysqlBase, "jdbc:mysql://%s",
		&targetAndDB)
	if err != nil || n != 1 {
		fmt.Printf("scan[%v] result is [%d] , want 1 , error is %v\n",
			mysqlBase, n, err)
		return
	}

	fmt.Printf("targetAndDB is %v\n",
		targetAndDB)

	/* step 2.5: get DB name */
	targetAndDBArray := strings.Split(targetAndDB, "/")
	if len(targetAndDBArray) < 2 {
		return
	}
	target := targetAndDBArray[0]
	*DBName = targetAndDBArray[1]

	/* step 3: get DB Host and port */
	hp := strings.Split(target, ":")
	if len(hp) < 1 {
		return
	} else if len(hp) == 1 {
		*DBHost = target
		return
	} else if len(hp) == 2 {
		var host string
		var port int
		fmt.Sscanf(target, "%s:%d",
			&host, &port)
		*DBHost = target
		*DBPort = port
	}

}

func initHostFromConfig(cfg *ini.File) {
	hikaricpurl := cfg.Section("").Key("hikaricp.url").String()
	if len(hikaricpurl) != 0 {
		analyzeURL(hikaricpurl)
		return
	}

	url := cfg.Section("").Key("url").String()
	if len(url) != 0 {
		analyzeURL(url)
		return
	}
}

func initUserFromConfig(cfg *ini.File) {
	hikaricpuser := cfg.Section("").Key("hikaricp.username").String()
	if len(hikaricpuser) != 0 {
		*DBUser = hikaricpuser
		return
	}

	user := cfg.Section("").Key("username").String()
	if len(user) != 0 {
		*DBUser = user
		return
	}
}

func initPasswordFromConfig(cfg *ini.File) {
	hikaricppassword := cfg.Section("").Key("hikaricp.password").String()
	if len(hikaricppassword) != 0 {
		*DBPassword = hikaricppassword
		return
	}

	password := cfg.Section("").Key("password").String()
	if len(password) != 0 {
		*DBPassword = password
		return
	}
}

func getDBConfig() error {
	configfile := getDBConfigFile()
	if len(configfile) == 0 {
		return nil
	}
	cfg, err := ini.Load(configfile)
	if err != nil {
		log.Printf("Cannot Find Config file, use default values\n")
		return err
	}

	if cfg != nil {
		initHostFromConfig(cfg)
		initUserFromConfig(cfg)
		initPasswordFromConfig(cfg)
	}
	return nil
}

/*DBHost is used to set the DB host IP */
var DBHost = flag.String("host", "127.0.0.1", "Host IP of name service Database")

/*DBPort is used to set the DB host Port */
var DBPort = flag.Int("P", 3306, "Port used to login name service Database")

/*DBUser is used to set the DB host User */
var DBUser = flag.String("u", "itil", "user used to login name service Database")

/*DBPassword is used to set the DB host Password */
var DBPassword = flag.String("p", "itil", "Password to use when connecting to name service Database server.")

/*DBName is used to set monitor router config database name */
var DBName = flag.String("D", "proc_conf", "name service Database to use.")

var globalDB *sql.DB
var dbOnce sync.Once

/*InitGlobalDBManager should call before use all other db feature*/
func InitGlobalDBManager() {
	dbOnce.Do(func() {
		err := getDBConfig()
		if err != nil {
			log.Printf("config file exist but Get DB From configfile failed:%v\n", err)
		}
		connectComman := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
			*DBUser, *DBPassword, *DBHost, *DBPort, *DBName)
		globalDB, err = sql.Open("mysql", connectComman)
		if err != nil {
			fmt.Printf("DB Open Error %v\n", err)
			globalDB = nil
		}
		fmt.Printf("Connect Success to : %v\n", connectComman)
	})

}

/*ListAllIndex is for all route*/
func ListAllIndex() error {
	queryString := fmt.Sprintf("SELECT index_path,title_en FROM tb_project_base;")
	tables, err := globalDB.Query(queryString)
	if err != nil {
		return fmt.Errorf("show tables failed:%v", err)

	}
	for tables.Next() {
		var indexPath string
		var title string

		err = tables.Scan(&indexPath, &title)
		if err != nil {
			return fmt.Errorf("Scan Error :%v", err)
		}

		fmt.Printf("indexPath is %v, title is %v\n", indexPath, title)

	}
	return nil
}

func searchIndex(project string) string {
	queryString := fmt.Sprintf("SELECT index_path FROM tb_project_base where title like '%%%v%%';",
		project)
	tables, err := globalDB.Query(queryString)
	if err != nil {
		fmt.Printf("show tables failed:%v", err)
		return ""
	}
	for tables.Next() {
		var indexPath string

		err = tables.Scan(&indexPath)
		if err != nil {
			fmt.Printf("Scan Error :%v", err)
			return ""
		}

		return indexPath
	}
	return ""
}
