package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// In Memory data store
var categories = []Category{
	{ID: 1, Name: "Peralatan Rumah", Description: "Peralatan rumah tangga"},
	{ID: 2, Name: "Furnitur", Description: "Perabot rumah dan kantor"},
	{ID: 3, Name: "Kecantikan", Description: "Produk kosmetik dan perawatan"},
}

func main() {
	http.HandleFunc("/categories/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getCategoryByID(w, r)

		} else if r.Method == http.MethodPut {
			updateCategory(w, r)

		}else if r.Method == http.MethodDelete {
			deleteCategory(w, r)
		}else {
			http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getCategories(w, r)

		} else if r.Method == http.MethodPost {
			createCategory(w, r)

		} else {
			http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server running di localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("gagal running server:", err)
	}
}

func getCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func createCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	newCategory.ID = len(categories) + 1
	categories = append(categories, newCategory)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCategory)
}

func updateCategory(w http.ResponseWriter, r *http.Request) {
	//GET id dulu
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	if idStr == "" {
		http.Error(w, "ID Tidak boleh kosong", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr) //-> diubah ke Int
	if err != nil {
		http.Error(w, "Invalid Kategori ID", http.StatusBadRequest)
		return
	}

	// get data dari request
	var updateCategory Category
	err = json.NewDecoder(r.Body).Decode(&updateCategory)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	//loop kategori , cari id , ganti sesuai data dari request
	for i := range categories {
		if categories[i].ID == id {
			updateCategory.ID = id
			categories[i] = updateCategory

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateCategory)
			return
		}
	}
	http.Error(w, "Kategori belum ada", http.StatusNotFound)
}

func getCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	if idStr == "" {
		http.Error(w, "ID Tidak boleh kosong", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr) //-> diubah ke Int
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	for _, p := range categories {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json") // -> Set biar jadi JSON
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	http.Error(w, "Kategori belum ada", http.StatusNotFound)
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	if idStr == "" {
		http.Error(w, "ID Tidak boleh kosong", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr) //-> diubah ke Int
	if err != nil {
		http.Error(w, "Invalid Kategori ID", http.StatusBadRequest)
		return
	}

	for i, p := range categories {
		if p.ID == id {
			categories = append(categories[:i], categories[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Kategori belum ada", http.StatusNotFound)	
}

