package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/oxks/myauth/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
