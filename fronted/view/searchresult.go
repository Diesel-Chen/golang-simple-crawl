package view

import (
	"html/template"
	"io"
)

type SearchResultView struct {
	template *template.Template
}

func CreateSearchResultView(filename string) SearchResultView {
	t := template.Must(template.ParseFiles(filename))
	return SearchResultView{template: t}
}

func (s SearchResultView) Render(wr io.Writer, data interface{}) error {
	return s.template.Execute(wr, data)
}
