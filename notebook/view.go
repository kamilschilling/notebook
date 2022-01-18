package notebook

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type view struct{}

func RegisterRoutes(echo *echo.Echo) {
	v := view{}
	echo.GET("", v.getMainPage)
}

func (v *view) getMainPage(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "main.html", nil)
}
