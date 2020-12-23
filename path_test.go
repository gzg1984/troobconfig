package troobconfig

import (
	"testing"
)

func TestInitGlobalBasePath(t *testing.T) {
	err := InitGlobalBasePath()
	if err != nil {
		t.Logf("InitGlobalBasePath error:%v", err)
		t.FailNow()
	}
	t.Logf("All Projects Are in : %v\n", globalProjectBasePath)
	t.Logf("All Index Are in : %v\n", globalIndexBasePath)
}
