package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Products struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Price       string `json:"price"`
}

func main() {
	http.HandleFunc("/", handleroot)
	http.HandleFunc("/AddProduct", handlepostrequest)
	http.HandleFunc("/GetProduct", handleGet)
	http.HandleFunc("/PutProduct", handleput)
	http.HandleFunc("/DeleteProduct", handleDelete)
	http.ListenAndServe(":8000", nil)

}

func handleroot(w http.ResponseWriter, request *http.Request) {
	w.Write([]byte("Welcome to Ecommerce Website Program"))
}
func handlepostrequest(w http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		var user Products
		var Add []Products
		err := json.NewDecoder(request.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid request: %v", err)
			return
		}
		Add = append(Add, user)
		w.WriteHeader(http.StatusCreated)
		file, _ := json.MarshalIndent(Add, "", " ")

		f, err := os.OpenFile("data.json", os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println(err)
		}

		_, err = io.WriteString(f, string(file))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Item was succesfully added")

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Println("method not found")
	}

}

func handleGet(w http.ResponseWriter, request *http.Request) {
	var Get []Products
	file, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Print(err)
	}
	_ = json.Unmarshal(file, &Get)

	if request.Method == http.MethodGet {
		json.NewEncoder(w).Encode(Get)
		fmt.Println("Product Sent Succesfully")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Invalid method")
	}
}
func handleput(w http.ResponseWriter, request *http.Request) {
	var Edit []Products
	if request.Method == http.MethodPut {
		var user Products
		err := json.NewDecoder(request.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid request: %v", err)
			return
		}
		Edit = append(Edit, user)
		w.WriteHeader(http.StatusCreated)
		file, _ := json.MarshalIndent(Edit, "", " ")

		data, _ := os.Create("data.json")
		if err != nil {
			fmt.Println(err)
		}

		_, err = io.WriteString(data, string(file))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Item was succesfully Edited")

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Println("method not found")
	}

}

func handleDelete(w http.ResponseWriter, request *http.Request) {
	var Delete []Products
	if request.Method == http.MethodDelete {
		var user Products
		err := json.NewDecoder(request.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid request: %v", err)
			return
		}
		Delete = append(Delete, user)
		w.WriteHeader(http.StatusCreated)
		file, _ := json.MarshalIndent(Delete, "", " ")

		data, _ := os.Create("data.json")
		if err != nil {
			fmt.Println(err)
		}

		_, err = io.WriteString(data, string(file))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Item was succesfully Deleted")

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Println("method not found")
	}

}
