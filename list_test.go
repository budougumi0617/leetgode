package leetgode

import (
	"context"
	"io/ioutil"
	"testing"
)

func TestListCmd(t *testing.T) {
	cmd := &ListCmd{}
	if err := cmd.Run(context.TODO(), ioutil.Discard, []string{}); err != nil {
		t.Errorf("ListCmd() error = %v", err)
	}
}
