package actions

import (
	"buffalo_crud/constants/messages"
	"buffalo_crud/helpers"
	"buffalo_crud/models"
	"log"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/pkg/errors"
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

	return c.Render(http.StatusOK, r.Auto(c, verrs))
}

func currentUser(c buffalo.Context) error {
	header_token := c.Request().Header.Get("AppAuth")
	token := &models.AccessToken{}

	tx := c.Value("tx").(*pop.Connection)
	if err := tx.Where("access_token = ?", header_token).Eager("User").First(token); err != nil {
		log.Printf("err: %v", err)
		return errors.WithStack(err)
	}

	return c.Render(http.StatusOK, r.Auto(c, token))
}

// Authorize require a user be logged in before accessing a route
func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		header_token := c.Request().Header.Get("AppAuth")
		tx := c.Value("tx").(*pop.Connection)
		access_token := &models.AccessToken{}
		err := tx.Where("access_token = ?", header_token).First(access_token)
		resp := helpers.NewServerResponse()
		resp.Code = http.StatusForbidden
		resp.Message = messages.NotAllowedToAccess

		if err != nil {
			return c.Render(http.StatusForbidden, r.JSON(resp))
		}
		return next(c)
	}
}
