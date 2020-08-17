package utils

import "testing"

func TestPassword(t *testing.T) {
	ps := "password"
	hs := GetHashedPassword(ps)
	if ps != hs {
		t.Logf("GetHashedPassword PASSED")
	} else {
		t.Errorf("GetHashedPassword FAILED, expected password to be different from hashed password.")
	}
}
