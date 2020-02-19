package main

import (
	// "database/sql"
	"fmt"
	//format - functions related to input and output
	"net/http"
	// "strconv"
  // "reflect"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
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

	db.AutoMigrate(&songModel{})
}

func main() {

	router := gin.Default()

	songs := router.Group("/api/v1/songs")
	{
		songs.POST("/", addSong)
		songs.GET("/", fetchAllSongs)
		songs.GET("/:id", fetchSong)
		// v1.PUT("/:id", changeSong)
		// v1.DELETE("/:id", removeSong)
	}
	router.Run()
}

type (
	songModel struct {
		gorm.Model
		Title         string  `json:"title"`
		SpotifyId     string  `json:"spotifyid"`
		URL           string  `json:"url"`
		Delay         float64 `json:"delay"`
		AvBarDuration float64 `json:"avbarduration"`
		Duration      float64 `json:"duration"`
		Tempo         float64 `json:"tempo"`
		TimeSignature int64    `json: "timesignature"`
	}
	songInput struct {
		Title         string  `json:"title"`
		SpotifyId     string  `json:"spotifyid"`
		URL           string  `json:"url"`
		Delay         float64 `json:"delay"`
		AvBarDuration float64 `json:"avbarduration"`
		Duration      float64 `json:"duration"`
		Tempo         float64 `json:"tempo"`
		TimeSignature int64    `json: "timesignature"`
	}

	transformedSong struct {
		ID            uint    `json:"id"`
		Title         string  `json:"title"`
		SpotifyId     string  `json:"spotifyid"`
		URL           string  `json:"url"`
		Delay         float64 `json:"delay"`
		AvBarDuration float64 `json:"avbarduration"`
		Duration      float64 `json:"duration"`
		Tempo         float64 `json:"tempo"`
		TimeSignature int64    `json: "timesignature"`
	}
)

func addSong(context *gin.Context) {

  var body songInput
  context.BindJSON(&body)

	song := songModel{
    Title: body.Title,
    SpotifyId: body.SpotifyId,
    URL: body.URL,
    Delay: body.Delay,
    AvBarDuration: body.AvBarDuration,
    Duration: body.Duration,
    Tempo: body.Tempo,
    TimeSignature: body.TimeSignature,
  }

	db.Save(&song)

	context.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Song created successfully!", "resourceId": song.ID})
}

func fetchAllSongs(context *gin.Context) {

	var songs []songModel

	db.Find(&songs)

	if len(songs) <= 0 {
		context.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No songs found!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": songs})
}


func fetchSong(context *gin.Context) {
	var song songModel
	songID := context.Param("id")

	db.First(&song, songID)

	if song.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No song found!"})
		return
	}

	_song := transformedSong{ID: song.ID, Title: song.Title, SpotifyId: song.SpotifyId, URL: song.URL, Delay: song.Delay, AvBarDuration: song.AvBarDuration, Duration: song.Duration, Tempo: song.Tempo, TimeSignature: song.TimeSignature}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _song})
}

// changeSong update a song

// func changeSong(c *gin.Context) {
// 	var song songModel
// 	songID := c.Param("id")
//
// 	db.First(&song, songID)
//
// 	if song.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No song found!"})
// 		return
// 	}
//
// 	db.Model(&song).Update("title", c.PostForm("title"))
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Song updated successfully!"})
// }

// removeSong remove a song

// func removeSong(c *gin.Context) {
// 	var song songModel
// 	songID := c.Param("id")
//
// 	db.First(&song, songID)
//
// 	if song.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No song found!"})
// 		return
// 	}
//
// 	db.Delete(&song)
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Song deleted successfully!"})
// }
