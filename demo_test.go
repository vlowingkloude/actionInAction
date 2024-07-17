package actionInAction

import "testing"

func TestGetMagicNumber(t *testing.T) {
	if GetMagicNumber() != 42 {
		t.Error("Now we have a bigger problem...")
	}
}
