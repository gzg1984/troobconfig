package troobconfig

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"gopkg.in/ini.v1"
)

/*
func initOnce() {
	initlocker.Do(func() {
		flag.Parse()

		cfg, err := ini.Load(*ConfigFile)
		if err != nil {
			log.Printf("Cannot Find Config file[%v], use default values", *ConfigFile)
		}

		err = InitNameServiceWithConfig(cfg)
		if err != nil {
			log.Printf("InitNameServiceWithConfig err:%v", err)
		}
	})

}*/

func getDBConfig() error {
	cfg, err := ini.Load("database.properties")
	if err != nil {
		log.Printf("Cannot Find Config file, use default values\n")
		return err
	}

	if cfg != nil {
		temphost := cfg.Section("").Key("hikaricp.url").String()
		if len(temphost) == 0 {
			return fmt.Errorf("Get DB Host From Config File Error")
		}
		log.Printf("URL:%v\n", temphost)

	}
	return nil
}

/*DBHost is used to set the DB host IP */
var DBHost = flag.String("host", "127.0.0.1", "Host IP of name service Database")

/*DBPort is used to set the DB host Port */
var DBPort = flag.Int("P", 3476, "Port used to login name service Database")

/*DBUser is used to set the DB host User */
var DBUser = flag.String("u", "itil", "user used to login name service Database")

/*DBPassword is used to set the DB host Password */
var DBPassword = flag.String("p", "itil", "Password to use when connecting to name service Database server.")

/*DBName is used to set monitor router config database name */
var DBName = flag.String("D", "proc_conf", "name service Database to use.")

var db *sql.DB

/*InitNameService should call before use all name service routers
func InitNameService() error {
	var err error
	connectComman := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		*DBUser, *DBPassword, *DBHost, *DBPort, *DBName)
	db, err = sql.Open("mysql", connectComman)
	if err != nil {
		fmt.Printf("DB Open Error %v\n", err)
		db = nil
		return err
	}

	err = CheckTableInfo()
	if err != nil {
		return err
	}
	go func() {
		for {
			time.Sleep(10 * time.Second)
			err := CheckTableInfo()
			if err != nil {
				log.Printf("CheckTableInfo Failed: %v\n", *DBHost)
			}
		}

	}()
	return nil
}*/

/*InitNameServiceWithConfig should be called before using every routers
func InitNameServiceWithConfig(cfg *ini.File) error {
	if cfg != nil {
		temphost := cfg.Section("mysql").Key("db_host").String()
		if len(temphost) == 0 {
			return fmt.Errorf("Get DB Host From Config File Error")
		}
		*DBHost = temphost
		log.Printf("Get DB Host From Config File: %v", *DBHost)

		tempname := cfg.Section("mysql").Key("db_name").String()
		if len(tempname) == 0 {
			return fmt.Errorf("GetDB Name From Config File Error")
		}
		*DBName = tempname
		log.Printf("Get DB Name  From Config File: %v", *DBName)

		tempUser := cfg.Section("mysql").Key("db_user").String()
		if len(tempUser) == 0 {
			return fmt.Errorf("Get DB User From Config File Error")
		}
		*DBUser = tempUser
		log.Printf("Get DB User From Config File: %v", *DBUser)

		tempPasswd := cfg.Section("mysql").Key("db_password").String()
		if len(tempPasswd) == 0 {
			return fmt.Errorf("Get DB Password From Config File Error")
		}
		*DBPassword = tempPasswd
		log.Printf("Get DB Password From Config File: %v", *DBPassword)

		tempPort, err := cfg.Section("mysql").Key("db_port").Int()
		if err != nil {
			return fmt.Errorf("Get Key From Config File Error: %v", err)
		}
		*DBPort = tempPort
		log.Printf("Get DB Port From Config File: %v", *DBPort)

	}
	return InitNameService()
}
*/

/*GetGlobalDBName return DBName from globle variable*/
func GetGlobalDBName() string {
	return *DBName
}
