package actions

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"buffalo_crud/helpers"
	"buffalo_crud/models"
)

// AuthCreate attempts to log the user in with an existing account.
func AuthCreate(c buffalo.Context) error {

	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return errors.WithStack(err)

	}

	tx := c.Value("tx").(*pop.Connection)

	// find a user with the email
	err := tx.Where("email = ?", strings.ToLower(strings.TrimSpace(u.Email))).First(u)

	// helper function to handle bad attempts
	bad := func() error {
		verrs := validate.NewErrors()
		verrs.Add("email", "invalid email/password")
		c.Set("errors", verrs)
		c.Set("user", u)

		return c.Render(http.StatusOK, r.Auto(c, verrs))
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
	user_id := u.ID

	//check user has non expired token

	access_token := &models.AccessToken{}
	err = tx.Where("user_id = ?", user_id).Where("expires_at > NOW()").First(access_token)
	if err != nil {
		fmt.Print("ERROR!\n")
		fmt.Printf("%v\n", err)
	}

	token := helpers.String(40)
	access_token.AccessToken = token
	tx.Update(access_token)

	return c.Render(http.StatusOK, r.JSON(access_token))
}
