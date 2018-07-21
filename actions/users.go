package actions

import "github.com/gobuffalo/buffalo"

// UsersShow default implementation.
func UsersShow(c buffalo.Context) error {
	return c.Render(200, r.HTML("users/show.html"))
}

// UsersIndex default implementation.
func UsersIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("users/index.html"))
}

// UsersCreate default implementation.
func UsersCreate(c buffalo.Context) error {
	return c.Render(200, r.HTML("users/create.html"))
}

// UsersMessage default implementation.
func UsersMessage(c buffalo.Context) error {
	return c.Render(200, r.HTML("users/message.html"))
}
