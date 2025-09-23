package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("q") != "" {
			fmt.Fprintf(w, `
				<body>
					<h1>HELLLOOOOOO, %s</h1>
				</body>
			`, r.FormValue("q"))
			return
		}
		fmt.Fprintf(w, `
			<body>
				<form action="/" method="GET">
					<label>SEU NOME</label>
					<input name="q">
					<button type="submit">Submit</button>
				</form>
			</body>
		`)
	})

	http.ListenAndServe(":80", nil)
}
