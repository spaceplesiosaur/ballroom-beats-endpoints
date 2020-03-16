package main

import (
  "fmt"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setUp() *App {

  app := &App{}

  app.Initialize("sqlite3", ":memory:")

  return app
}

func tearDown(app *App) {
  app.DB.Close()
}

func TestAddSong(t *testing.T) {
  app := setUp()

  testBody :=  "{\"Title\": \"blue\", \"SpotifyId\": \"12345\", \"URL\": \"www.moose.com\", \"Delay\": 2, \"AvBarDuration\": 3, \"Duration\": 123, \"Tempo\": 4, \"TimeSignature\": 8}"

  stringifiedBody := strings.NewReader(testBody)

  //if I hadn't written this as a string already, I'd have had to use the .Encode method
  router := app.MakeRouter()

  req, err := http.NewRequest("POST", "/api/v1/songs/", stringifiedBody)
	if err != nil {
		t.Fatal(err)
	}
	// Our API expects a form body, so set the content-type header to make sure it's treated as one.
	req.Header.Add("Content-Type", "application/json")

	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

  // assert.Equal(t, http.StatusOK, rr.Code)
  //use with testify
  if res.Code != http.StatusCreated {
    t.Errorf("ouchie")
  }

  tearDown(app)
}

func TestAddSongMissingString(t *testing.T) {
  app := setUp()

  testBody :=  "{\"SpotifyId\": \"12345\", \"URL\": \"www.moose.com\", \"Delay\": 2, \"AvBarDuration\": 3, \"Duration\": 123, \"Tempo\": 4, \"TimeSignature\": 8}"

  stringifiedBody := strings.NewReader(testBody)

  //if I hadn't written this as a string already, I'd have had to use the .Encode method
  router := app.MakeRouter()

  req, err := http.NewRequest("POST", "/api/v1/songs/", stringifiedBody)
	if err != nil {
		t.Fatal(err)
	}
	// Our API expects a form body, so set the content-type header to make sure it's treated as one.
	req.Header.Add("Content-Type", "application/json")

	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

  // assert.Equal(t, http.StatusOK, rr.Code)
  //use with testify
  if res.Code != http.StatusUnprocessableEntity {
    t.Errorf("ouchie")
  }

  tearDown(app)
}

func TestAddSongMissingFloat(t *testing.T) {
  app := setUp()

  testBody :=  "{\"Title\": \"blue\", \"SpotifyId\": \"12345\", \"URL\": \"www.moose.com\", \"Delay\": 2, \"AvBarDuration\": 3, \"Tempo\": 4, \"TimeSignature\": 8}"

  stringifiedBody := strings.NewReader(testBody)

  //if I hadn't written this as a string already, I'd have had to use the .Encode method
  router := app.MakeRouter()

  req, err := http.NewRequest("POST", "/api/v1/songs/", stringifiedBody)
	if err != nil {
		t.Fatal(err)
	}
	// Our API expects a form body, so set the content-type header to make sure it's treated as one.
	req.Header.Add("Content-Type", "application/json")

	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

  // assert.Equal(t, http.StatusOK, rr.Code)
  //use with testify
  if res.Code != http.StatusUnprocessableEntity {
    t.Errorf("ouchie")
  }

  tearDown(app)
}

func TestFetchAllSongs(t *testing.T) {
  app := setUp()

  testSongs := []songModel{
      songModel{Title: "Blue da boo dee", SpotifyId: "12345", URL: "www.moose1.com", Delay: 2, AvBarDuration: 3, Tempo: 4, TimeSignature: 8},
      songModel{Title: "Yellow", SpotifyId: "123456", URL: "www.moose2.com", Delay: 1, AvBarDuration: 2, Tempo: 5, TimeSignature: 4},
    }

    for _, song := range testSongs {
      app.DB.Save(&song)
  	}

    router := app.MakeRouter()
    req, err := http.NewRequest("GET", "/api/v1/songs/", nil)
    if err != nil {
      t.Fatal(err)
    }

    res := httptest.NewRecorder()

  	router.ServeHTTP(res, req)

    if res.Code != http.StatusOK {
      t.Errorf("ouchie")
    }

    tearDown(app)
}

func TestFetchAllSongsNoSongs (t *testing.T) {
  app := setUp()

  router := app.MakeRouter()
  req, err := http.NewRequest("GET", "/api/v1/songs/", nil)
  if err != nil {
    t.Fatal(err)
  }

  res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

  if res.Code != http.StatusNotFound{

    t.Errorf("ouchie")
  }

  tearDown(app)
}

func TestFetchSong(t *testing.T) {
  app := setUp()

  testSongs := []songModel{
      songModel{Title: "Blue da boo dee", SpotifyId: "12345", URL: "www.moose1.com", Delay: 2, AvBarDuration: 3, Tempo: 4, TimeSignature: 8},
      songModel{Title: "Yellow", SpotifyId: "123456", URL: "www.moose2.com", Delay: 1, AvBarDuration: 2, Tempo: 5, TimeSignature: 4},
    }

    for _, song := range testSongs {
      app.DB.Save(&song)
  	}

    // testParam := "1"
    // stringifiedParam := strings.NewReader(testParam)

    var firstSong songModel

    app.DB.First(&firstSong)
    // testParam := "1"
    // stringifiedParam := strings.NewReader(testParam)

    router := app.MakeRouter()
    req, err := http.NewRequest("GET", fmt.Sprintf("/api/v1/songs/%d",
      firstSong.ID), nil)
    if err != nil {
      t.Fatal(err)
    }

    res := httptest.NewRecorder()

  	router.ServeHTTP(res, req)

    if res.Code != http.StatusOK {
      t.Errorf("ouchie")
    }

    tearDown(app)
}

func TestFetchSongNoSong(t *testing.T) {
  app := setUp()

  testSongs := []songModel{
      songModel{Title: "Blue da boo dee", SpotifyId: "12345", URL: "www.moose1.com", Delay: 2, AvBarDuration: 3, Tempo: 4, TimeSignature: 8},
      songModel{Title: "Yellow", SpotifyId: "123456", URL: "www.moose2.com", Delay: 1, AvBarDuration: 2, Tempo: 5, TimeSignature: 4},
    }

    for _, song := range testSongs {
      app.DB.Save(&song)
  	}

    // testParam := "1"
    // stringifiedParam := strings.NewReader(testParam)

    router := app.MakeRouter()
    req, err := http.NewRequest("GET", "/api/v1/songs/5000", nil)
    if err != nil {
      t.Fatal(err)
    }

    res := httptest.NewRecorder()

  	router.ServeHTTP(res, req)

    if res.Code != http.StatusNotFound {
      t.Errorf("ouchie")
    }

    tearDown(app)
}

func TestRemoveSong(t *testing.T) {
  app := setUp()

  testSongs := []songModel{
      songModel{Title: "Blue da boo dee", SpotifyId: "12345", URL: "www.moose1.com", Delay: 2, AvBarDuration: 3, Tempo: 4, TimeSignature: 8},
      songModel{Title: "Yellow", SpotifyId: "123456", URL: "www.moose2.com", Delay: 1, AvBarDuration: 2, Tempo: 5, TimeSignature: 4},
    }

    for _, song := range testSongs {
      app.DB.Save(&song)
  	}

    var firstSong songModel

    app.DB.First(&firstSong)
    // testParam := "1"
    // stringifiedParam := strings.NewReader(testParam)

    router := app.MakeRouter()
    req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/songs/%d",
  		firstSong.ID), nil)
    if err != nil {
      t.Fatal(err)
    }

    res := httptest.NewRecorder()

  	router.ServeHTTP(res, req)

    if res.Code != http.StatusNoContent {
      t.Errorf("ouchie")
    }

    tearDown(app)
}

func TestRemoveSongNoSong(t *testing.T) {
  app := setUp()

  testSongs := []songModel{
      songModel{Title: "Blue da boo dee", SpotifyId: "12345", URL: "www.moose1.com", Delay: 2, AvBarDuration: 3, Tempo: 4, TimeSignature: 8},
      songModel{Title: "Yellow", SpotifyId: "123456", URL: "www.moose2.com", Delay: 1, AvBarDuration: 2, Tempo: 5, TimeSignature: 4},
    }

    for _, song := range testSongs {
      app.DB.Save(&song)
  	}

    // testParam := "1"
    // stringifiedParam := strings.NewReader(testParam)

    router := app.MakeRouter()
    req, err := http.NewRequest("DELETE", "/api/v1/songs/5000", nil)
    //we expect this to never be a valid ID
    if err != nil {
      t.Fatal(err)
    }

    res := httptest.NewRecorder()

  	router.ServeHTTP(res, req)

    if res.Code != http.StatusNotFound {
      t.Errorf("ouchie")
    }

    tearDown(app)
}
