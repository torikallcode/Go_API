package main

import (
	"html/template"
	"net/http"
)

type Product struct {
	Name          string
	Price         float64
	InStock       bool
	Discount      float64
	AfterDiscount float64
}

type PageData struct {
	Title    string
	Products []Product
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("products.html")
	if err != nil {
		http.Error(w, "Gagal parsing template", http.StatusInternalServerError)
		return
	}

	products := []Product{
		{Name: "Laptop", Price: 1000000, InStock: true, Discount: 0.1},
		{Name: "Mouse", Price: 300000, InStock: false, Discount: 0.1},
		{Name: "Keyboard", Price: 500000, InStock: true, Discount: 0.1},
	}

	for n := range products {
		if products[n].InStock {
			products[n].AfterDiscount = float64((products[n].Price)) * float64((1 - products[n].Discount))
		} else {
			products[n].AfterDiscount = products[n].Price
		}
	}

	data := PageData{
		Title:    "Daftar Products",
		Products: products,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handleProducts)
	http.ListenAndServe(":8080", nil)
}
