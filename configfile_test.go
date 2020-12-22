package troobconfig

import "testing"

func TestConfig(t *testing.T) {
	configfile := getConfigFile()
	t.Logf("%v", configfile)
}
