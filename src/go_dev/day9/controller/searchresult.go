package controller

import (
	"go_dev/day9/view"
	"net/http"
	"strconv"
	"strings"
)

type SearchResultHandle struct {
	view view.SearchResultView

}


func CreateSearchResultHandler(template string)SearchResultHandle{

}




func (h SearchResultHandle) ServeHTTP(w http.ResponseWriter,req *http.Request){
	g := strings.TrimSpace(req.FormValue("q"))
	form,err := strconv.Atoi(req.FormValue("from"))
}
