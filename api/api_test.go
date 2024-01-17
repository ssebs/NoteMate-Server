package api

import (
	"testing"

	"github.com/ssebs/padpal-server/util"
)

func TestAPIFoo(t *testing.T) {
	got := APIFoo()
	want := "APIFoo"
	util.GotWantTest[string](got, want, t)
}
