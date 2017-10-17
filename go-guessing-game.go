/*
*Go-guessing-game
*@Author: Kevin Gleeson
*Date: 15/10/2017
*Source: https://github.com/data-representation/go-echo
*Version: 1.0
*Sources: 
*https://golang.org/pkg/net/http/#SetCookie
*https://stackoverflow.com/questions/12130582/setting-cookies-in-golang-net-http
*https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/06.1.html
*https://astaxie.gitbooks.io/build-web-application-with-golang/en/07.4.html
*
*/

package main

import (
	"html/template"
	"net/http"
	"time"
	"fmt"
	"math/rand"
	"strconv"
)
//struct declaration
type TodoPageData struct {
	PageTitle string

}
func requestHandler(w http.ResponseWriter, r *http.Request) {
	// guessing game echoed out
	//fmt.Fprintln(w, "<h1>Guessing Game</h1>")
////////############################Cookie start#################################
	//condition to check if the cookie length is 0 (exists)
	//it will run this only once if it is 0 as the next time it 
	//will have avlue greater than  0.
	if len(r.Cookies()) == 0{
	// cookie will get expired after 1 year 
    expires := time.Now().AddDate(1, 0, 0)
	ck := http.Cookie{
		//set target for the random number
		Name: "target",
		Path: "/",
		Expires: expires,
	}
	//one single random value
	rand.Seed(time.Now().UnixNano())
	tarNum := rand.Intn(20-1) + 1
	// value of cookie converted to string
	//@todo convert back to integer when comparing the values
	ck.Value = strconv.Itoa(tarNum)

	// write the cookie to response
	http.SetCookie(w, &ck)
}
	// read cookie
	var cookie,err = r.Cookie("target")
	if err == nil {
		//test the value target on the terminal
		var cookievalue = cookie.Value
		fmt.Println(w, "<b>get cookie value is " + cookievalue + "</b>\n")
}

//////////########################Cookie end###############################

	//set the header content type to text/html
	w.Header().Set("Content-Type", "text/html")
	
   //struct set string for the template
	data := TodoPageData{
		PageTitle: "Pick a number between 1 and 20",
		}
	//parse the static folder for any template files
	tmpl := template.Must(template.ParseGlob("static/*"))
	//pass PageTile string to guess.tmpl
	tmpl.ExecuteTemplate(w,"guess.tmpl", data)
	
}


func main() {
	
	
	//store the directory where the html and template files are held
	fs := http.FileServer(http.Dir("static"))
	//Start at the root directory
	http.Handle("/", fs)
	//select the index.html file
	http.HandleFunc("/index", requestHandler)
	//select from the current directory
	http.Handle("./", fs)
	//handle the guess request
	http.HandleFunc("/guess", requestHandler)

	//Listen out for requests to the server
	http.ListenAndServe(":8080", nil)
	
}