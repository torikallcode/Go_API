package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.New("form").Parse(`
<!DOCTYPE html>
<html>
	<head>
		<title>Form Page</title>
	</head>
	<body>
		<h1>submit your detail</h1>
		<form method="POST" actin="/form">
			<label for="name">Name:</label>
			<input type="text" id="name" name="name">
			<br><br>
			<label for="email">Email:</label>
			<input type="email" id="email" name="email">
			<br><br>
			<input type="submit" value="Submit">
		</form>
	</body>
</html>
`))

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Fatal()
		}

		name := r.FormValue("name")
		email := r.FormValue("email")

		fmt.Fprintf(w, "Name: %s\n", name)
		fmt.Fprintf(w, "Emaail: %s\n", email)
	} else {
		http.Error(w, "Metode tidak didukung", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/form", formHandler)
	fmt.Println("Server berjalan di http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error saat menjalankan server", err)
	}
}
