package troobconfig

import "testing"

func TestDB(t *testing.T) {
	err := getDBConfig()
	t.Logf("%v", err)
}
