package utils

import (
	"testing"
)

func TestGetUUID(t *testing.T) {
	u := GetUUID()
	t.Log(u, len(u))
}
