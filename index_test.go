package troobconfig

import (
	"flag"
	"testing"
)

var TestProject = flag.String("project", "spdk", "project title to search")

/*TestGetIndexPath will give the index path from the project name*/
func TestGetIndexPath(t *testing.T) {
	r := GetIndexPath(*TestProject)
	t.Logf("Result:%v\n", r)
}
