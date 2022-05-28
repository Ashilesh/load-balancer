package utils_test

import (
	"os"

	"testing"

	"github.com/Ashilesh/load-balancer/utils"
)

var (
	testCmdData = [...][3]string{
		{"field1", "args1", "correct"},
		{"field2", "", "err"},
	}
)

func TestGetCmdArgs(t *testing.T) {
	for _, val := range testCmdData {
		os.Args = []string{val[0], val[1]}
		arg, err := utils.GetCmdArgs(val[0])
		if err != nil && val[2] != "err" {
			t.Errorf("Expected to get argument but got error")
		} else if arg != val[1] {
			t.Errorf("Expected to get %s but got %s", val[1], arg)
		}
	}
}
