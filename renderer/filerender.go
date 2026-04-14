package renderer

import (
	"context"
	"fmt"
	"text/template"

	"github.com/aarondl/authboss/v3"
)

type FileRenderer struct{}

func NewFileRenderer() *FileRenderer {

	return &FileRenderer{}
}

func (f FileRenderer) Load(names ...string) error { return nil }

func (f FileRenderer) Render(ctx context.Context, page string, data authboss.HTMLData) ([]byte, string, error) {

	path := ""

	if page == "login" {

		path = "static/login/login.html"

	} else if page == "register" {

		path = "static/register/register.html"

	}

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return nil, "", fmt.Errorf("views/%s.html not found: %v", page, err)
	}

	var buf []byte
	w := &writerProxy{b: &buf}

	if err := tmpl.Execute(w, data); err != nil {
		return nil, "", err
	}
	return buf, "text/html", nil
}

type writerProxy struct{ b *[]byte }

func (w *writerProxy) Write(p []byte) (int, error) {
	*w.b = append(*w.b, p...)
	return len(p), nil
}
