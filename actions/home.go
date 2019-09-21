package actions

import (
	"log"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {

	userSID := c.Session().Get("current_user_id")

	if userSID != nil {

		log.Println("We are on the index page")
	}

	return c.Render(200, r.HTML("index.plush.html"))
}
