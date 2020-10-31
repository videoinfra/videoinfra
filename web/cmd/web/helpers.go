package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, name string,
	td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("The template %s does not exist", name))
		return
	}

	buf := new(bytes.Buffer)

	err := ts.Execute(buf, app.addDefaultData(td, r))
	if err != nil {
		app.serverError(w, err)
		return
	}

	buf.WriteTo(w)
}

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}
	td.CurrentYear = time.Now().Year()

	// Add the flash message to the template data, if one exists.
	td.Flash = app.session.PopString(r, "flash")
	return td
}

