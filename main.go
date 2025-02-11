package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

    "hugo.mardbrink.se/internal/handlers"
)

type Template struct {
    template *template.Template
}

func newTemplate() *Template {
    return &Template{
        template: template.Must(template.ParseGlob("views/*.html")),
    }
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.template.ExecuteTemplate(w, name, data)
}

func main() {
    e := echo.New()

    e.Renderer = newTemplate()
    e.Use(middleware.Logger())

    e.Static("/resources", "resources")
    e.Static("/css", "css")

    handlers.RegisterRoutes(e)

    e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}
