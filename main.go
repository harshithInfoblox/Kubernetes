package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Define Model Struct
type Usern struct {
	ID       uint
	Name     string
	Age      int
	Birthday time.Time
	CreatedAt time.Time // Add CreatedAt field
}

var err error
var db *gorm.DB

func main() {
	// Connect to PostgreSQL Database

	//db, err = gorm.Open(postgres.Open("postgres://postgres:@localhost:5432/newprojectdb?sslmode=disable"), &gorm.Config{})
	//db, err = gorm.Open(postgres.Open("postgres://hp:Harshith@postgres:5432/newprojectdb?sslmode=disable"), &gorm.Config{})
	time.Sleep(10 * time.Second)
	dsn := "host=postgres user=postgres password=Harshith dbname=prithvi port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open("postgres", dsn)

	// url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
	// pql.conf.Postgres.User,
	// pql.conf.Postgres.Password,
	// pql.conf.Postgres.Host,
	// pql.conf.Postgres.Port,
	// pql.conf.Postgres.DB)
	//db,err = gorm.Open(postgres.Open("postgres://postgres:@localhost:5432/newprojectdb?sslmode=disable"),&gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// err = db.AutoMigrate(&User{})
	// if err != nil {
	// 	log.Fatal("failed to migrate database", err)
	// }
	if err := db.AutoMigrate(&Usern{}).Error; err != nil {
		// Handle the error here
		log.Fatal(err)
	}
	// Initialize HTTP server with Gorilla Mux
	r := mux.NewRouter()
	fmt.Println("correct till here")
	r.HandleFunc("/users", CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", GetUserByIDHandler).Methods("GET")

	// Start HTTP server
	fmt.Println("HTTP server listening on port 8080")
	http.ListenAndServe(":8080", r)
}

// HTTP Handler functions
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Implement logic to create user using HTTP
	// Example:
	user := Usern{Name: "Harsh", Age: 50, Birthday: time.Now()}
	result := db.Create(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created successfully!")
}

func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Implement logic to retrieve user by ID using HTTP
	// Example:
	vars := mux.Vars(r)
	userID := vars["id"]
	var user Usern
	result := db.First(&user, userID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	// Return user details as JSON response
	json.NewEncoder(w).Encode(user)
}
