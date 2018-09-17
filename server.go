package main

import (
	"html/template"
	"io"	
	"net/http"

	"github.com/labstack/echo"
)

type (
	TemplateRenderer struct {
		templates *template.Template
	}
)


// Render a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("/app/**/*.html")),
	}
	e.Renderer = renderer

	//----------
	// Handlers
	//----------
	e.Static("/assets", "/app/public/assets/")
	e.File("/favicon.ico", "/app/public/assets/images/favicon.png")

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", map[string]map[string]interface{}{
			"data": {
				"current": "home",
			},
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}