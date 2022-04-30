package main

import ("github.com/labstack/echo/v4"
"net/http")

func main() {

	api := echo.New()
	api.GET("home", hello)
	api.POST("home", hello)
	api.PUT("home", hello)

	api.Start(":9090")
	// localhost:9090
}
func hello(c echo.Context) error {
	    return c.JSON(200, "Tejaswini B K")
}
func anotherHandler(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}