package main
import (

      "github.com/labstack/echo"
    )
func main() {

  e := echo.New()

  e.File("/favicon.ico", "assets/favicon.ico")
  e.File("/Lato-Medium.woff2", "assets/Lato-Medium.woff2")
  e.File("/", "assets/index.html")

  e.Logger.Fatal(e.Start(":1324"))
}
