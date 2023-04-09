package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Heinrich/110_Go_Booking_App/pkg/config"
	"github.com/Heinrich/110_Go_Booking_App/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template packag
func NewTemplate(a *config.AppConfig){
	app = a
}

// RenderTemplate Renders Template using HTML template
func RenderTemplate(w http.ResponseWriter, tmpl string,td *models.TemplateData){

	var tc map[string] *template.Template
	if app.UseCache {
		tc = app.TemplateCache
	}else{
		tc,_ = CreateTemplateCache()
	}
	
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing to browser: ", err)
	}

	// _,err := RenderTemplateTest()
	// if err != nil {
	// 	fmt.Println("error getting template cache")
	// }
	// parsedTemplate,_ := template.ParseFiles("./templates/"+tmpl)

	// err = parsedTemplate.Execute(w,nil)

	// if err != nil {
	// 	 fmt.Println("Error parsing template",err)
	// 	 return
	// }
}
//  CreateTemplateCache creates TemplateCache as a map
func CreateTemplateCache() (map[string] *template.Template,error){

	myCache := map[string]*template.Template{}

	pages,err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache,err
	}

	for _,page := range pages {
		name := filepath.Base(page)

		ts,err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache,err
		}

		matches,err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return myCache,err
		}

		if len(matches) > 0{
			ts,err = ts.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				return myCache,err
			}
			
		}
		myCache[name] = ts
		
	}

	return myCache,nil

}