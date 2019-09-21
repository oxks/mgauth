package actions

import (
	"database/sql"
	"fmt"
	"strings"

	//"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
	"github.com/oxks/myauth/models"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	//"reflect"
)

// AuthNew loads the signin page
func AuthNew(c buffalo.Context) error {
	c.Set("user", models.User{})

	return c.Render(200, r.HTML("auth/new.html"))
}

// AuthCreate attempts to log the user in with an existing account.
func AuthCreate(c buffalo.Context) error {
	//c.Flash().Add("error", "test")
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return errors.WithStack(err)
	}

	tx := c.Value("tx").(*pop.Connection)

	// find a user with the email
	err := tx.Where("email = ?", strings.ToLower(strings.TrimSpace(u.Email))).First(u)

	// helper function to handle bad attempts
	bad := func() error {
		c.Set("user", u)
		verrs := validate.NewErrors()
		verrs.Add("email", "invalid email/password")
		c.Set("errors", verrs)
		return c.Render(422, r.HTML("auth/new.html"))
	}

	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			// couldn't find an user with the supplied email address.
			return bad()
		}
		return errors.WithStack(err)
	}

	// confirm that the given password matches the hashed password from the db
	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(u.Password))
	if err != nil {
		return bad()
	}
	c.Session().Set("current_user_id", u.ID)

	c.Flash().Add("success", "Welcome Back to Buffalo!")

	redirectURL := "/"
	if redir, ok := c.Session().Get("redirectURL").(string); ok {
		redirectURL = redir
	}

	return c.Redirect(302, redirectURL)
}

// AuthDestroy clears the session and logs a user out
func AuthDestroy(c buffalo.Context) error {
	c.Session().Clear()
	c.Flash().Add("success", "You have been logged out!")
	return c.Redirect(302, "/")
}

func init() {
	gothic.Store = App().SessionStore

	// other providers supported by Goth can be added here, the format is following:
	// set global variables for private data (GOOGLE_KEY, GITHUB_KEY etc)
	goth.UseProviders(
		google.New("GOOGLE_KEY", "GOOGLE_SECRET", fmt.Sprintf("%s%s", App().Host, "/auth/google/callback")),
		github.New("GITHUB_KEY", "GITHUB_SECRET", fmt.Sprintf("%s%s", App().Host, "/auth/github/callback")),
	)

}

// AuthCallback function, here you get provider's authorisation data
func AuthCallback(c buffalo.Context) error {
	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())

	if err != nil {
		return c.Error(401, err)
	}
	// Do something with the user, maybe register them/sign them in

	c.Set("user", user)

	return c.Render(200, r.JSON(user))
}
