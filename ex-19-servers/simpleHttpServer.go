package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)

	io.WriteString(
		res,
		`<DOCTYPE html>
            <html>
            <head>
                <title>Hello World</title>
            </head>
            <body>
                <script src="./assets/hello.js"></script>
                Hello World!
            </body>
        </html>`,
	)
}

func main() {
	http.Handle(
		"/assets/",
		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("assets")),
		),
	)
	http.HandleFunc("/tim", hello)
	http.ListenAndServe(":7000", nil)
}
