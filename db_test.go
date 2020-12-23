package troobconfig

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDB(t *testing.T) {
	err := getDBConfig()
	if err != nil {
		t.Logf("%v", err)
		t.FailNow()
	}
}

func TestInitGlobalDBManager(t *testing.T) {
	err := InitGlobalDBManager()
	if err != nil {
		t.Logf("InitGlobalDBManager error:%v", err)
		t.FailNow()
	}
	err = LocateIndex()
	if err != nil {
		t.Logf("LocateIndex error:%v", err)
		t.FailNow()
	}
}
