package app

import (
	"cmp"
	"errors"
	"net/http"

	"github.com/a-h/templ"
	"github.com/jovandeginste/workout-tracker/pkg/geocoder"
	appviews "github.com/jovandeginste/workout-tracker/views"
	"github.com/labstack/echo/v4"
)

var ErrUserNotFound = errors.New("user not found")

func (a *App) redirectWithError(c echo.Context, target string, err error) error {
	if err != nil {
		a.setError(c, "Something went wrong: "+err.Error())
	}

	return c.Redirect(http.StatusFound, target)
}

func (a *App) statisticsHandler(c echo.Context) error {
	data := a.defaultData(c)
	data["since"] = cmp.Or(c.QueryParam("since"), "1 year")
	data["per"] = cmp.Or(c.QueryParam("per"), "month")

	return c.Render(http.StatusOK, "user_statistics.html", data)
}

func (a *App) dashboardHandler(c echo.Context) error {
	data := a.defaultData(c)

	u := a.getCurrentUser(c)
	if u == nil {
		return a.redirectWithError(c, a.echo.Reverse("user-signout"), ErrUserNotFound)
	}

	if err := a.addWorkouts(u, data); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-signout"), err)
	}

	if err := a.addUsers(data); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-signout"), err)
	}

	if err := a.addRecentWorkouts(data); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-signout"), err)
	}

	data["user"] = u

	return c.Render(http.StatusOK, "user_show.html", data)
}

func (a *App) userLoginHandler(c echo.Context) error {
	data := a.defaultData(c)

	return c.Render(http.StatusOK, "user_login.html", data)
}

func (a *App) lookupAddressHandler(c echo.Context) error {
	a.setContext(c)

	q := c.FormValue("location")

	results, err := geocoder.Search(q)
	if err != nil {
		a.setError(c, "Something went wrong: "+err.Error())
	}

	return Render(c, http.StatusOK, appviews.AddressResults(results))
}

func (a *App) heatmapHandler(c echo.Context) error {
	data := a.defaultData(c)

	u := a.getCurrentUser(c)
	if u == nil {
		return a.redirectWithError(c, a.echo.Reverse("user-signout"), ErrUserNotFound)
	}

	if err := a.addWorkouts(u, data); err != nil {
		return a.redirectWithError(c, a.echo.Reverse("user-signout"), err)
	}

	return c.Render(http.StatusOK, "heatmap.html", data)
}

func (a *App) testHandler(c echo.Context) error {
	a.setContext(c)

	a.setError(c, "Something went wrong!")

	return Render(c, http.StatusOK, appviews.Test())
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
