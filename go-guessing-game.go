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
*https://stackoverflow.com/questions/26189523/go-represent-path-without-query-string
*https://stackoverflow.com/questions/20320549/how-can-you-delete-a-cookie-in-an-http-response
 */

 
//@toDo place conditionals for dynamic result templates
//@toDO form input validation 
package main

import (
	//"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

//struct declaration
type TodoPageData struct {
	//name of template variables to be used in .tmpl file
	PageTitle   string
	GuessTmpl   string
	ResultTmpl  string
	CorrectTmpl string
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	// guessing game echoed out
	//fmt.Fprintln(w, "<h1>Guessing Game</h1>")

	//condition to check if the cookie length is 0 (exists)
	//it will run this only once if it is 0 as the next time it
	//will have avlue greater than  0.

	if len(r.Cookies()) == 0 {
		// cookie will expire after 1 year
		expires := time.Now().AddDate(1, 0, 0)
		ck := http.Cookie{
			//set target for the random number
			Name:    "target",
			Path:    "/",
			Expires: expires,
		}
		//one single random value
		rand.Seed(time.Now().UnixNano())
		tarNum := rand.Intn(20-1) + 1
		// value of cookie converted to string
		
		ck.Value = strconv.Itoa(tarNum)

		// write the cookie to response
		http.SetCookie(w, &ck)
	}

	//Initialise string for use outside of conditional
	guessStr := ""
	tmplGuess := ""
	
	if len(r.URL.Query().Get("guess")) >= 0 {
		r.ParseForm()
		//@ToDo form validation for integer value
		//create variable guessStr and store the users guess from the client side
		guessStr = r.Form.Get("guess")
		tmplGuess = string(guessStr)
	}
	//Cast to string

	//test guess var in terminal when the submit button is pressed
	//fmt.Println(guess)
	//for result message template
	var result string
	var correct string
	// read cookie
	var cookie, err = r.Cookie("target")
	//if cookie is set
	if err == nil {
		//set strings tmplGuess and cookie.Value to int
		guessInt, _ := strconv.Atoi(tmplGuess)
		tarInt, _ := strconv.Atoi(cookie.Value)
		//input validation for numbers between 1 and 20 only.
		if guessInt > 0 && guessInt < 21{ 
			//conditional to check for parity between the two numbers
			if guessInt < tarInt {
				result = "you have guessed too low!!!"
			} else if guessInt > tarInt {
				result = "You have guessed too high!!!"
			} else { //reset the cookie length to 0 if the number has been guessed
				correct = "Well done you have guessed the correct number!!!!!!!"
				expires := time.Now().AddDate(1, 0, 0)
				ck := http.Cookie{
					//set target for the random number
					Name:    "target",
					Path:    "/",
					Expires: expires,
				}
				
				//one single random value
				rand.Seed(time.Now().UnixNano())
				tarNum := rand.Intn(20-1) + 1
				// value of cookie converted to string
				
				ck.Value = strconv.Itoa(tarNum)

				// write the cookie to response
				http.SetCookie(w, &ck)
			}
		}else{
			result = "That is not a valid number"
		}
			//Testing to console ouput 
			//fmt.Println(w, "<b>get cookie value is ", guessInt, tarInt, "</b>\n")
	}
	//test the value target on the terminal
	//struct set string for the template
	
	data := TodoPageData{
		PageTitle:   "Pick a number between 1 and 20",
		GuessTmpl:   tmplGuess,
		ResultTmpl:  result,
		CorrectTmpl: correct,
	}
	//parse the static folder for any template files
	tmpl := template.Must(template.ParseGlob("static/*"))
	//pass PageTile string to guess.tmpl
	tmpl.ExecuteTemplate(w, "guess.tmpl", data)

	//set the header content type to text/html
	w.Header().Set("Content-Type", "text/html")

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