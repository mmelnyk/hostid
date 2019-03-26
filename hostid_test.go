package hostid

import "testing"

func TestSimpleCall(t *testing.T) {
	id, err := ID()
	if err != nil {
		t.Fail()
	}
	t.Log(id)
}
