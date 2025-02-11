package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"hugo.mardbrink.se/internal/models"
)

func RegisterRoutes(e *echo.Echo,) {
    e.GET("/", homePageHandler())
    e.GET("/projects", projectsPageHandler())
    e.GET("/articles", articlesPageHandler())
}

func homePageHandler() echo.HandlerFunc {
    return func(c echo.Context) error {
        isHtmx := c.Request().Header.Get("HX-Request") == "true"
        c.Request().Header.Add("Vary", "HX-Request")
        p := models.Page{
                Title: "Hugo Mårdbrink",
                Description: "Home page of Hugo Mårdbrinks personal website.", 
                IsHtmx: isHtmx,
                Breadcrumbs: models.Breadcrumbs{models.NewBreadcrumb("home", "/")}}
        hp := models.NewHomePage(p)
        
        return c.Render(http.StatusOK, "home", hp)
    }
}

func projectsPageHandler() echo.HandlerFunc {
    return func(c echo.Context) error {
        isHtmx := c.Request().Header.Get("HX-Request") == "true"
        c.Request().Header.Add("Vary", "HX-Request")
       
        p := models.Page{
                Title: "Hugo Mårdbrink - Projects",
                Description: "Hobby projects by Hugo Mårdbrink.", 
                IsHtmx: isHtmx,
                Breadcrumbs: models.Breadcrumbs{
                    models.NewBreadcrumb("Home", "/"),
                    models.NewBreadcrumb("projects", "/projects")}}
        pp := models.NewProjectPage(p)
        
        return c.Render(http.StatusOK, "projects", pp)
    }
}

func articlesPageHandler() echo.HandlerFunc {
    return func(c echo.Context) error {
        isHtmx := c.Request().Header.Get("HX-Request") == "true"
        c.Request().Header.Add("Vary", "HX-Request")
       
        p := models.Page{
                Title: "Hugo Mårdbrink - Articles",
                Description: "Articles by Hugo Mårdbrink.", 
                IsHtmx: isHtmx,
                Breadcrumbs: models.Breadcrumbs{
                    models.NewBreadcrumb("Home", "/"),
                    models.NewBreadcrumb("articles", "/articles")}}
        pp := models.NewProjectPage(p)
        
        return c.Render(http.StatusOK, "articles", pp)
    }
}
