package template

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo"

	"github.com/acm-uiuc/core/config"
)

type Template struct {
	template *template.Template
}

func (tmpl *Template) Render(w io.Writer, name string, data interface{}, ctx echo.Context) error {
	return tmpl.template.ExecuteTemplate(w, name, data)
}

var tmpl *Template

func New() (*Template, error) {
	if tmpl != nil {
		return tmpl, nil
	}

	templateSrcs, err := config.GetConfigValue("TEMPLATE_SRCS")
	if err != nil {
		return nil, fmt.Errorf("failed to get template srcs: %w", err)
	}

	rawTemplate, err := template.ParseGlob(templateSrcs)
	if err != nil {
		return nil, fmt.Errorf("failed to parse templates: %w", err)
	}

	tmpl = &Template{
		template: rawTemplate,
	}

	return tmpl, nil
}
