package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/jonniespratley/jps_go_app/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
