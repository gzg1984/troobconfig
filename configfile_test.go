package troobconfig

import "testing"

func TestDBConfig(t *testing.T) {
	configfile := getDBConfigFile()
	if len(configfile) == 0 {
		t.Log("Cannot Find Config File")
		t.FailNow()
	}
	t.Logf("%v", configfile)

}
