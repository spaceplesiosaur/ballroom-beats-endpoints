package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
//This import gin, is like express, it's a library for handling http requests and responses
	"github.com/jinzhu/gorm"
	//gorm is an ORM, and object relationship manager.  It is kind of like knex. This builds your database
	_ "github.com/lib/pq"
		//this is our import to use postgres
)


const (
	host = "localhost"
	port = 5432
	dbname = "bbpractice"
)

type App struct {
  DB *gorm.DB
}

func (a *App) Initialize(dbType string, dbInfo string) {

	var err error

	a.DB, err = gorm.Open(dbType, dbInfo)

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	a.DB.AutoMigrate(&songModel{})
}

func main() {

	dbInfo := fmt.Sprintf("host=%s port=%d "+
		"dbname=%s sslmode=disable",
		host, port, dbname)

	a := &App{}

	a.Initialize("postgres", dbInfo)
	router := a.MakeRouter()
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

func (a *App) MakeRouter() *gin.Engine{
	router := gin.Default()

	songs := router.Group("/api/v1/songs")
	{
		songs.POST("/", a.AddSong)
		songs.GET("/", a.FetchAllSongs)
		songs.GET("/:id", a.FetchSong)
		songs.DELETE("/:id", a.RemoveSong)
	}

	return router
}

//AddSong posts a new songs

func (a *App) AddSong(context *gin.Context) {
  var body songInput

  context.BindJSON(&body)

	if body.Title == "" || body.SpotifyId == "" ||  body.URL == "" {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "message": "Has a missing or malformed string property"})
		return
	}

	if body.Delay == 0 || body.AvBarDuration == 0 ||  body.Duration == 0 || body.Tempo == 0 || body.TimeSignature == 0 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "message": "Missing or malformed numerical property"})
		return
	}

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

	a.DB.Save(&song)

	context.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Song created successfully!", "resourceId": song.ID})
}

//FetchAllSongs fetches every song in the database

func (a *App) FetchAllSongs(context *gin.Context) {

	var songs []songModel

	a.DB.Find(&songs)

	if len(songs) == 0 {
		context.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No songs found!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": songs})
}

// FetchSong fetches a single song

func (a *App) FetchSong(context *gin.Context) {
	var song songModel
	songID := context.Param("id")

	a.DB.First(&song, songID)

	if song.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No song found!"})
		return
	}

	_song := transformedSong{ID: song.ID, Title: song.Title, SpotifyId: song.SpotifyId, URL: song.URL, Delay: song.Delay, AvBarDuration: song.AvBarDuration, Duration: song.Duration, Tempo: song.Tempo, TimeSignature: song.TimeSignature}

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _song})
}

//RemoveSong deletes a song

func (a *App) RemoveSong(context *gin.Context) {
	var song songModel

	songID := context.Param("id")

	a.DB.First(&song, songID)

	if song.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No song found!"})
		return
	}

	a.DB.Delete(&song)

	context.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "Successfully deleted!"})

}
