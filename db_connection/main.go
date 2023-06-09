package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	db "example.com/go_progresif/config"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<h1>About</h1>
	<h1>About Me</h1>`)
}

func getAlbums(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	albums, err := db.AlbumsAll()
	if err != nil {
		panic(err.Error())
	}
	jsonData, err := json.Marshal(albums)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprint(w, string(jsonData))
}

func main() {

	// HTTP router
	// Handle url path requests
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	// Establish Database connection
	db.Connect()

	// API Endpoint - RESTful API endpoint
	// GET
	http.HandleFunc("/get/albums", getAlbums)

	// router := gin.Default()
	// router.GET("/albums", getAlbums)
	// router.Run("localhost:3000")

	// CRUD
	// Create
	albID, err := db.AddAlbum(db.Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)

	// Read (multiple records)
	albums, err := db.AlbumsByArtist("John Coltrane")
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Albums found: %v\n", albums)

	// Read (singular record)
	alb, err := db.AlbumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)

	// Update (singular record)
	updID, err := db.UpdateAlbumByID(5, db.Album{
		Title:  "Updated",
		Artist: "Updated",
		Price:  00.00,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of updated album: %v\n", updID)

	//Delete (singular record)
	delID, err := db.DeleteAlbumByID(albID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of deleted album: %v\n", delID)

	// Start local http server @ PORT:3000
	fmt.Println("Server Started at PORT:3000")
	http.ListenAndServe(":3000", nil)
}
