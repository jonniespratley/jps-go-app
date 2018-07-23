package actions

import (
	"fmt"
	"os"
	"log"

	"github.com/gobuffalo/buffalo"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
  )

func init() {
	gothic.Store = App().SessionStore
	callback := fmt.Sprintf("%s%s", App().Host, "/auth/github/callback")
	log.Printf("callback", callback)
	goth.UseProviders(
	github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), callback),
	)
}

func AuthCallback(c buffalo.Context) error {
	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(401, err)
	}
	log.Printf("user", user)
	// Do something with the user, maybe register them/sign them in
	return c.Render(200, r.JSON(user))
}
