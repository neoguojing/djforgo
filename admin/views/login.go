package views

import (
	l4g "github.com/alecthomas/log4go"
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, req *http.Request) {
	l4g.Info(req)
	
	if req.Method != http.MethodPost {
		html, err := template.ParseFiles("./admin/templates/login.html")
		if err != nil {
			l4g.Error("Login parse html faild", err)
			return
		}
	
		html.Execute(w, nil)
		return
	}
	
	
	
	
}
