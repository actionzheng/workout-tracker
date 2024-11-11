package app

import (
	"strconv"

	"github.com/jovandeginste/workout-tracker/pkg/database"
	"github.com/labstack/echo/v4"
)

func (a *App) adminAddUser(data map[string]any, c echo.Context) error {
	user, err := a.getUser(c)
	if err != nil {
		return err
	}

	if user == nil {
		return ErrUserNotFound
	}

	data["user"] = user

	return nil
}

func (a *App) getUser(c echo.Context) (*database.User, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, err
	}

	u, err := database.GetUserByID(a.db, id)
	if err != nil {
		return nil, err
	}

	return u, nil
}
