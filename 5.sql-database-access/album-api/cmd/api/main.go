package main

import (
	"log"

	"album-api/internal/album"
	"album-api/internal/database"
)

func main() {
	// Database Layer
	db, err := database.NewConnection()
	if err != nil {
		log.Fatalf("couldn't connnect to the database: %v", err)
	}
	log.Println("MySQL DB connected and ready for operation.")

	// Repository Layer
	albumRepo := album.NewRepository(db)

	// Service Layer
	albumService := album.NewService(albumRepo)

	// Handler(Controller) Layer
	albumHandler := album.NewHandler(albumService)

	// In a production grade application, you would start an http server here and pass this albumHandler to it.
	// Here we just simulate manual API calls, in reality these would be created by the client and would reach the controllers

	log.Println("\n--- Running Application ---")
	albumHandler.AddNewAlbum("Sajna", "Pujan Khunt", 399.31)
	albumHandler.GetAlbumByID(5)
	albumHandler.GetAlbumsByArtist("Pujan Khunt")
}
