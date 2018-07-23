package actions

import (
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/packr"
)

var r *render.Engine
var assetsBox = packr.NewBox("../public")

func isLoggedIn(help plush.HelperContext) bool {
	return help.Value("current_user") != nil
  }

func init() {
	r = render.New(render.Options{

		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.html",

		// Box containing all of the templates:
		TemplatesBox: packr.NewBox("../templates"),
		AssetsBox:    assetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{
			"is_logged_in": isLoggedIn,

			// uncomment for non-Bootstrap form helpers:
			// "form":     plush.FormHelper,
			// "form_for": plush.FormForHelper,
		},
	})
}
