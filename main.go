package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "He is Habib, He is Youtuber")
}
func getProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "please Give Me GET Request", 400)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(productlist)
}

type Product struct {
	ID          int
	Title       string
	Description string
	Price       float64
	ImgUrl      string
}

var productlist []Product

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/getProducts", getProducts)
	fmt.Println("Server is running on : 3000 port")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting the Server: ", err)
		return
	}

}
func init() {
	prod1 := Product{
		ID:          1,
		Title:       "Orrange",
		Description: "Orrange is red. I love orrange",
		Price:       180,
		ImgUrl:      "xxxxxxxxxxxxxx1"
	}
	prod2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is good. I love it",
		Price:       220,
		ImgUrl:      "xxxxxxxxxxxxxx2",
	}
	prod3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is best for health",
		Price:       120,
		ImgUrl:      "xxxxxxxxxxxxxx3",
	}
	productlist = append(productlist, prod1)
	productlist = append(productlist, prod2)
	productlist = append(productlist, prod3)

}
