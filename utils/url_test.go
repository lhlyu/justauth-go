package utils

import (
	"testing"
)

func TestUrlBuilder(t *testing.T) {
	ub := NewUrlBuilder("http://www.baidu.com")
	t.Log(ub.Build())
	ub.QueryParam("a", 1)
	ub.QueryParam("b", "xxx")
	t.Log(ub.Build())
}
