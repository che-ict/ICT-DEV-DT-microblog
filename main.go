package main

import (
	"github.com/che-ict/DEV-DT-Microblog/handlers"
	"github.com/che-ict/DEV-DT-Microblog/repositories"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

func main() {

	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	e := echo.New()
	e.Debug = true
	e.Renderer = t

	e.Use(echoprometheus.NewMiddleware("microblog")) // adds middleware to gather metrics
	e.GET("/metrics", echoprometheus.NewHandler())   // adds route to serve gathered metrics

	e.GET("/login", handlers.LoginGetHandler)
	e.POST("/login", handlers.LoginPostHandler)
	e.GET("/register", handlers.RegisterGetHandler)
	e.POST("/register", handlers.RegisterPostHandler)
	e.GET("/feed", handlers.FeedHandler)
	e.GET("/user/:user", handlers.UserPageHandler)
	e.POST("/post", handlers.CreatePostHandler)
	e.DELETE("/post/:id", handlers.DeletePostHandler)

	repositories.DB()

	e.Start(":1323")
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
