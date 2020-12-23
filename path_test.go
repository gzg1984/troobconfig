package troobconfig

import (
	"testing"
)

func TestInitGlobalBasePath(t *testing.T) {
	InitGlobalBasePath()

	t.Logf("All Projects Are in : %v\n", globalProjectBasePath)
	t.Logf("All Index Are in : %v\n", globalIndexBasePath)
}
