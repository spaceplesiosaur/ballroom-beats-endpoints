package main

import (
  "fmt"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
  "encoding/json"
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

  testBody := songInput {Title: "blue", SpotifyId: "12345", URL: "blue-song", Delay: 2, AvBarDuration: 3, Duration: 123, Tempo: 4, TimeSignature: 8}

  jsonTestBody, err := json.Marshal(testBody)
  stringTestBody := string(jsonTestBody)

  stringifiedBody := strings.NewReader(stringTestBody)

  router := app.MakeRouter()

  req, err := http.NewRequest("POST", "/api/v1/songs/", stringifiedBody)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")

	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

  if res.Code != http.StatusCreated {
    preferredStatus := http.StatusCreated
    actualStatus := res.Code
    errorMessage := fmt.Sprintf("Summary of failure: Expected %d, got %d", preferredStatus, actualStatus)
    t.Errorf(errorMessage)
  }

  tearDown(app)
}

func TestAddSongMissingString(t *testing.T) {
  app := setUp()

  testBody := songInput {SpotifyId: "12345", URL: "blue-song", Delay: 2, AvBarDuration: 3, Duration: 123, Tempo: 4, TimeSignature: 8}

  jsonTestBody, err := json.Marshal(testBody)
  stringTestBody := string(jsonTestBody)

  stringifiedBody := strings.NewReader(stringTestBody)

  router := app.MakeRouter()

  req, err := http.NewRequest("POST", "/api/v1/songs/", stringifiedBody)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

  if res.Code != http.StatusUnprocessableEntity {
    preferredStatus := http.StatusUnprocessableEntity
    actualStatus := res.Code
    errorMessage := fmt.Sprintf("Summary of failure: Expected %d, got %d", preferredStatus, actualStatus)
    t.Errorf(errorMessage)
  }

  tearDown(app)
}

func TestAddSongMissingFloat(t *testing.T) {
  app := setUp()

  testBody := songInput {Title: "blue", SpotifyId: "12345", URL: "blue-song", Delay: 2, AvBarDuration: 3, Tempo: 4, TimeSignature: 8}

  jsonTestBody, err := json.Marshal(testBody)
  stringTestBody := string(jsonTestBody)

  stringifiedBody := strings.NewReader(stringTestBody)

  router := app.MakeRouter()

  req, err := http.NewRequest("POST", "/api/v1/songs/", stringifiedBody)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

  if res.Code != http.StatusUnprocessableEntity {
    preferredStatus := http.StatusUnprocessableEntity
    actualStatus := res.Code
    errorMessage := fmt.Sprintf("Summary of failure: Expected %d, got %d", preferredStatus, actualStatus)
    t.Errorf(errorMessage)
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
      preferredStatus := http.StatusOK
      actualStatus := res.Code
      errorMessage := fmt.Sprintf("Summary of failure: Expected %d, got %d", preferredStatus, actualStatus)
      t.Errorf(errorMessage)
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

  if res.Code != http.StatusNotFound {
    preferredStatus := http.StatusNotFound
    actualStatus := res.Code
    errorMessage := fmt.Sprintf("Summary of failure: Expected %d, got %d", preferredStatus, actualStatus)
    t.Errorf(errorMessage)
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

    router := app.MakeRouter()
    req, err := http.NewRequest("GET", fmt.Sprintf("/api/v1/songs/%d",
      firstSong.ID), nil)
    if err != nil {
      t.Fatal(err)
    }

    res := httptest.NewRecorder()

  	router.ServeHTTP(res, req)

    if res.Code != http.StatusOK {
      preferredStatus := http.StatusOK
      actualStatus := res.Code
      errorMessage := fmt.Sprintf("Summary of failure: Expected %d, got %d", preferredStatus, actualStatus)
      t.Errorf(errorMessage)
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

    router := app.MakeRouter()
    req, err := http.NewRequest("GET", "/api/v1/songs/5000", nil)
    if err != nil {
      t.Fatal(err)
    }

    res := httptest.NewRecorder()

  	router.ServeHTTP(res, req)

    if res.Code != http.StatusNotFound {
      preferredStatus := http.StatusNotFound
      actualStatus := res.Code
      errorMessage := fmt.Sprintf("Summary of failure: Expected %d, got %d", preferredStatus, actualStatus)
      t.Errorf(errorMessage)
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

    router := app.MakeRouter()
    req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/songs/%d",
  		firstSong.ID), nil)
    if err != nil {
      t.Fatal(err)
    }

    res := httptest.NewRecorder()

  	router.ServeHTTP(res, req)

    if res.Code != http.StatusNoContent {
      preferredStatus := http.StatusNoContent
      actualStatus := res.Code
      errorMessage := fmt.Sprintf("Summary of failure: Expected %d, got %d", preferredStatus, actualStatus)
      t.Errorf(errorMessage)
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

    router := app.MakeRouter()
    req, err := http.NewRequest("DELETE", "/api/v1/songs/5000", nil)
    //we expect 5000 to never be a valid ID
    if err != nil {
      t.Fatal(err)
    }

    res := httptest.NewRecorder()

  	router.ServeHTTP(res, req)

    if res.Code != http.StatusNotFound {
      preferredStatus := http.StatusNotFound
      actualStatus := res.Code
      errorMessage := fmt.Sprintf("Summary of failure: Expected %d, got %d", preferredStatus, actualStatus)
      t.Errorf(errorMessage)
    }

    tearDown(app)
}
