package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/pkg/errors"

	"buffalo_crud/models"
)

// UsersCreate registers a new user with the application.
func UsersCreate(c buffalo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return errors.WithStack(err)
	}

	tx := c.Value("tx").(*pop.Connection)
	verrs, err := u.Create(tx)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		c.Set("user", u)
		c.Set("errors", verrs)
	}

	c.Session().Set("current_user_id", u.ID)
	return c.Render(http.StatusOK, r.Auto(c, verrs))
}

// Authorize require a user be logged in before accessing a route
func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		header_token := c.Request().Header.Get("AppAuth")
		tx := c.Value("tx").(*pop.Connection)
		access_token := &models.AccessToken{}
		err := tx.Where("access_token = ?", header_token).First(access_token)
		if err != nil {
			return c.Render(http.StatusForbidden, r.JSON("ss"))
		}
		return next(c)
	}
}
