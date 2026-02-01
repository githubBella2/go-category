package main

import (
	// "encoding/json"
	"fmt"
	"net/http"
	// "strconv"
	"strings"

	"categories-bella/database"
	"categories-bella/handlers"
	"categories-bella/repositories"
	"categories-bella/services"
	"github.com/spf13/viper"		
	"log"
	"os"

)


// In Memory data store
// var categories = []Category{
// 	{ID: 1, Name: "Peralatan Rumah", Description: "Peralatan rumah tangga"},
// 	{ID: 2, Name: "Furnitur", Description: "Perabot rumah dan kantor"},
// 	{ID: 3, Name: "Kecantikan", Description: "Produk kosmetik dan perawatan"},
// }
type Config struct {
	Port    string `mapstructure:"PORT"`
	DB_CONN string `mapstructure:"DB_CONN"`
}

// func main() {
    
//     viper.AutomaticEnv()
//     viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
//     if _, err := os.Stat(".env"); err == nil {
//         viper.SetConfigFile(".env")
//         _ = viper.ReadInConfig()
//     }
    
//     config := Config{
//         Port:    viper.GetString("PORT"),
//         DB_CONN: viper.GetString("DB_CONN"),
//     }
    
//     // Set default port kalo kosong
//     if config.Port == "" {
//         config.Port = "8080"
//     }
    
//     // Setup database
//     db, err := database.InitDB(config.DB_CONN)
//     if err != nil {
//         log.Fatal("Failed to initialize database:", err)
//     }
//     defer db.Close()
    
//     categoryRepo := repositories.NewCategoryRepository(db)
//     CategoryService := services.NewCategoryService(categoryRepo)
//     CategoryHandler := handlers.NewCategoryHandler(CategoryService)
    
//     //setup route
//     http.HandleFunc("/api/categories",  CategoryHandler.HandleCategories)
//     http.HandleFunc("/api/categories/", CategoryHandler.HandleCategoryByID)
    
//     // DEBUG
//     fmt.Println("PORT =", config.Port)
    
//     if config.DB_CONN == "" {
//         log.Fatal("DB_CONN KOSONG — .env tidak terbaca")
//     }
    
//     fmt.Printf("Server running on port %s\n", config.Port)
//     err = http.ListenAndServe(":"+config.Port, nil)
//     if err != nil {
//         log.Fatal("Gagal running server:", err)
//     }
// }

func main() {
    
    viper.AutomaticEnv()
    viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
    if _, err := os.Stat(".env"); err == nil {
        viper.SetConfigFile(".env")
        _ = viper.ReadInConfig()
    }
    
    config := Config{
        Port:    viper.GetString("PORT"),
        DB_CONN: viper.GetString("DB_CONN"),
    }
    
    // Set default port kalo kosong
    if config.Port == "" {
        config.Port = "8080"
    }
    
    // Setup database
    db, err := database.InitDB(config.DB_CONN)
    if err != nil {
        log.Fatal("Failed to initialize database:", err)
    }
    defer db.Close()
    
    categoryRepo := repositories.NewCategoryRepository(db)
    CategoryService := services.NewCategoryService(categoryRepo)
    CategoryHandler := handlers.NewCategoryHandler(CategoryService)
    
    //setup route
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Categories API is running!"))
    })
    http.HandleFunc("/api/categories",  CategoryHandler.HandleCategories)
    http.HandleFunc("/api/categories/", CategoryHandler.HandleCategoryByID)
    
    // DEBUG
    fmt.Println("PORT =", config.Port)
    fmt.Println("✅ Routes registered:")
    fmt.Println("  - / (health check)")
    fmt.Println("  - /api/categories")
    fmt.Println("  - /api/categories/")
    
    if config.DB_CONN == "" {
        log.Fatal("DB_CONN KOSONG — .env tidak terbaca")
    }
    
    fmt.Printf("Server running on port %s\n", config.Port)
    err = http.ListenAndServe(":"+config.Port, nil)
    if err != nil {
        log.Fatal("Gagal running server:", err)
    }
}

// func getCategories(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(categories)
// }

// func createCategory(w http.ResponseWriter, r *http.Request) {
// 	var newCategory Category
// 	err := json.NewDecoder(r.Body).Decode(&newCategory)
// 	if err != nil {
// 		http.Error(w, "Invalid Request", http.StatusBadRequest)
// 		return
// 	}

// 	newCategory.ID = len(categories) + 1
// 	categories = append(categories, newCategory)

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(newCategory)
// }

// func updateCategory(w http.ResponseWriter, r *http.Request) {
// 	//GET id dulu
// 	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
// 	if idStr == "" {
// 		http.Error(w, "ID Tidak boleh kosong", http.StatusBadRequest)
// 		return
// 	}

// 	id, err := strconv.Atoi(idStr) //-> diubah ke Int
// 	if err != nil {
// 		http.Error(w, "Invalid Kategori ID", http.StatusBadRequest)
// 		return
// 	}

// 	// get data dari request
// 	var updateCategory Category
// 	err = json.NewDecoder(r.Body).Decode(&updateCategory)
// 	if err != nil {
// 		http.Error(w, "Invalid request", http.StatusBadRequest)
// 		return
// 	}

// 	//loop kategori , cari id , ganti sesuai data dari request
// 	for i := range categories {
// 		if categories[i].ID == id {
// 			updateCategory.ID = id
// 			categories[i] = updateCategory

// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(updateCategory)
// 			return
// 		}
// 	}
// 	http.Error(w, "Kategori belum ada", http.StatusNotFound)
// }

// func getCategoryByID(w http.ResponseWriter, r *http.Request) {
// 	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
// 	if idStr == "" {
// 		http.Error(w, "ID Tidak boleh kosong", http.StatusBadRequest)
// 		return
// 	}

// 	id, err := strconv.Atoi(idStr) //-> diubah ke Int
// 	if err != nil {
// 		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
// 		return
// 	}

// 	for _, p := range categories {
// 		if p.ID == id {
// 			w.Header().Set("Content-Type", "application/json") // -> Set biar jadi JSON
// 			json.NewEncoder(w).Encode(p)
// 			return
// 		}
// 	}
// 	http.Error(w, "Kategori belum ada", http.StatusNotFound)
// }

// func deleteCategory(w http.ResponseWriter, r *http.Request) {
// 	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
// 	if idStr == "" {
// 		http.Error(w, "ID Tidak boleh kosong", http.StatusBadRequest)
// 		return
// 	}

// 	id, err := strconv.Atoi(idStr) //-> diubah ke Int
// 	if err != nil {
// 		http.Error(w, "Invalid Kategori ID", http.StatusBadRequest)
// 		return
// 	}

// 	for i, p := range categories {
// 		if p.ID == id {
// 			categories = append(categories[:i], categories[i+1:]...)
// 			w.WriteHeader(http.StatusNoContent)
// 			return
// 		}
// 	}
// 	http.Error(w, "Kategori belum ada", http.StatusNotFound)	
// }

