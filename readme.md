# Go Web Application with http

This is a repository with an example of a http web application with the [Go](https://golang.org/) programming language.

Author: [Kevin Gleeson](https://github.com/kevgleeson78)

Third year student at:[GMIT](http://gmit.ie) Galway

## Cloning, compiling and running the application.

1: Download [git](https://git-scm.com/downloads) to your machine if not already installed.

1.1: Download [go](https://golang.org/dl/) if not already installed.

2: Open git bash and cd to the folder you wish to hold the repository.
Alternatively you can right click on the folder and select git bash here.
This will open the git command prompt in the folder selected.
 
 3: To clone the repository type the following command in the terminal making sure you are in the folder needed for the repository.
```bash
>git clone https://github.com/kevgleeson78/go-guessing-game.git
```
4: To compile the application cd to the folder and type the following 
```bash
> go build 
```
This will compile and create an executable file from the .go file and give it the name of the folder.

5: To run the application ensure you cd to folder the application is held.
Type the following command
```bash
>./go-guessing-game
```
6: Go to the following URL to view the application in your browser
[http://127.0.0.1:8080](http://127.0.0.1:8080).

# Using Curl to see response result

1: Download [curl](https://curl.haxx.se/download.html) if not already on your machine.

2: Make sure the application is running as per step five above.

3: In command prompt type the following to view the response
```bash
>curl -v http://127.0.0.1:8080
```
This will print out all of the responses that are requested the command -v stands for verbose which will print out all of the hearder information.


