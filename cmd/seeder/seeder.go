package main

import (
	"github.com/OGFris/FlixFlex/internal/config"
	"github.com/OGFris/FlixFlex/internal/models"
	"github.com/cyruzin/golang-tmdb"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	// Load application configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize the database connection
	db, err := gorm.Open(mysql.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Initialize the TMDB client with your API key
	tmdbClient, err := tmdb.Init(cfg.TmdbApiKey)
	if err != nil {
		log.Fatalf("Failed to initialize the tmdb client: %v", err)
	}

	// Fetch movies from TMDB
	movies, err := tmdbClient.GetMovieTopRated(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Insert movies into your database
	for _, tmdbMovie := range movies.Results {
		movie := models.Movie{
			Title:     tmdbMovie.Title,
			Synopsis:  tmdbMovie.Overview,
			PosterUri: tmdbMovie.PosterPath,
			TmdbID:    uint(tmdbMovie.ID),
		}

		if err := db.Create(&movie).Error; err != nil {
			log.Printf("Error inserting movie: %v", err)
		} else {
			log.Printf("Inserted movie: %s", movie.Title)
		}
	}
}
