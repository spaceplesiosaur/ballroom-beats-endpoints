package main

import (
	"fmt"
	"net/http"
	// "strconv" // Was removing automatically so I commented it out
	// "reflect" // Was removing automatically so I commented it out
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	dbname = "ballroom_beats_development"
)

var db *gorm.DB

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d "+
		"dbname=%s sslmode=disable",
		host, port, dbname)
	var err error

	db, err = gorm.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// db.DropTable(&songModel{}) // We need to refactor our migrations to account for 'rollbacks' where we drop our tables and rebuild.
	db.AutoMigrate(&songModel{})
}

func main() {
	router := gin.Default()
	songs := router.Group("/api/v1/songs")
	{
		songs.GET("/", fetchAllSongs)
	}
	router.Run()
}

type (
	songModel struct {
		gorm.Model
		Title          string  `json:"title"`
		SpotifyID      string  `json:"spotify_id"`
		URL            string  `json:"url"`
		Delay          float64 `json:"delay"`
		AvgBarDuration float64 `json:"avg_bar_duration"`
		Duration       float64 `json:"duration"`
		Tempo          float64 `json:"tempo"`
		TimeSignature  uint    `json:"time_signature"`
	}

	transformedSong struct {
		ID             uint    `json:"id"`
		Title          string  `json:"title"`
		SpotifyID      string  `json:"spotify_id"`
		URL            string  `json:"url"`
		Delay          float64 `json:"delay"`
		AvgBarDuration float64 `json:"avg_bar_duration"`
		Duration       float64 `json:"duration"`
		Tempo          float64 `json:"tempo"`
		TimeSignature  uint    `json:"time_signature"`
	}
)

// Proof of concept, will need to be refactored according to
func fetchAllSongs(context *gin.Context) {
	var songs []songModel
	db.Find(&songs)

	if len(songs) <= 0 {
		context.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No songs found!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": songs})
}
