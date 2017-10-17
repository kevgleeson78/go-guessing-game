/*
*App-Name: Go-guessing-game
*@Author:  Kevin Gleeson
*Date:     15/10/2017
*Version:  1.0
*Sources: 
*https://github.com/data-representation/go-echo
*https://golang.org/pkg/net/http/#SetCookie
*https://stackoverflow.com/questions/12130582/setting-cookies-in-golang-net-http
*https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/06.1.html
*https://astaxie.gitbooks.io/build-web-application-with-golang/en/07.4.html
*https://stackoverflow.com/questions/22593259/check-if-string-is-int-golang
*https://stackoverflow.com/questions/28159520/passing-a-query-parameter-to-the-go-http-request-handler-using-the-mux-package
*https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/06.1.html
*https://github.com/gowww/client/blob/master/response.go
*https://godoc.org/hkjn.me/googleauth
*https://golang.org/pkg/strconv/
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
	//name of template variables to be used in .tmpl file
	PageTitle string
	GuessTmpl string
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
	//@ToDo form validation for integer value
	//create variable guessStr and store the users guess from the client side
	 guessStr := r.URL.Query().Get("guess")
	//Cast to string
	 tmplGuess := string(guessStr)

	//test guess var in terminal when the submit button is pressed
	//fmt.Println(guess)
//////////########################Cookie end###############################

	//set the header content type to text/html
	w.Header().Set("Content-Type", "text/html")
	
   //struct set string for the template
	data := TodoPageData{
		PageTitle: "Pick a number between 1 and 20",
		GuessTmpl: tmplGuess,
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