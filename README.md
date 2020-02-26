# ballroom-beats-service

This API was built for the [Ballroom Beats UI](https://github.com/Asilo5/Ballroom-Beats-UI).

## Created by:

- [Allison McCarthy](https://github.com/spaceplesiosaur)
- [Jonathan Patterson](https://github.com/Jonpatt92)

## View Ballroom Beats locally in your computer

Make sure you have Go installed on your machine.  If you don't, run `brew install go` in your terminal, or go to https://golang.org/ and install it there.  

Next, run `go env gopath` in terminal to see where your gopath is expected to be.  Make a go directory in that place, and make a src file in that directory.  Clone down this repo into your src file (` git clone https://github.com/Jonpatt92/ballroom-beats-service.git` ), and install the following packages into the src file as well:

`go get -u github.com/jinzhu/gorm`
`go get -u github.com/gin-gonic/gin`
`go get -u github.com/lib/pq`

`` cd `` into Ballroom Beats Service.

And to run on your local browser:

``$ go run main.go``

### Endpoints

## GET

``Get`` requests are used to read the API.  You will be adding endpoints `https://ballroom-blitz.herokuapp.com/`.  So, for instance, if you wanted to get all songs available from this API, you would fetch `https://ballroom-blitz.herokuapp.com/api/v1/songs/`.

#### GET all possible songs

Endpoint:

 `https://ballroom-blitz.herokuapp.com/api/v1/songs`

Example reply:

        {
          data: [
            {
              ID: 6,
              CreatedAt: "2020-02-24T17:23:48.84244Z",
              UpdatedAt: "2020-02-24T17:23:48.84244Z",
              DeletedAt: null,
              title: "Beyond the Sea",
              spotify_id: "",
              url: "beyond-the-sea",
              delay: 0.167,
              avg_bar_duration: 0,
              duration: 172.47,
              tempo: 136.482,
              time_signature: 4
            },
            {
              ID: 7,
              CreatedAt: "2020-02-24T21:31:04.952961Z",
              UpdatedAt: "2020-02-24T21:31:04.952961Z",
              DeletedAt: null,
              title: "Deja Vu",
              spotify_id: "",
              url: "deja-vu",
              delay: 0,
              avg_bar_duration: 0,
              duration: 196,
              tempo: 124,
              time_signature: 4
            },
            {
              ID: 8,
              CreatedAt: "2020-02-24T21:33:32.778773Z",
              UpdatedAt: "2020-02-24T21:33:32.778773Z",
              DeletedAt: null,
              title: "Game of Thrones",
              spotify_id: "",
              url: "game-of-thrones",
              delay: 0,
              avg_bar_duration: 0,
              duration: 106,
              tempo: 170,
              time_signature: 3
            },
            {
              ID: 9,
              CreatedAt: "2020-02-24T21:34:48.337236Z",
              UpdatedAt: "2020-02-24T21:34:48.337236Z",
              DeletedAt: null,
              title: "Melancolia Tropical",
              spotify_id: "",
              url: "melancolia-tropical",
              delay: 0,
              avg_bar_duration: 0,
              duration: 226,
              tempo: 121,
              time_signature: 4
            }
          ],
          status: 200
        }

#### GET a single song from list

Endpoint:

 `https://ballroom-blitz.herokuapp.com/api/v1/songs/:id`

Example reply:

        {
          data: {
            id: 6,
            title: "Beyond the Sea",
            spotify_id: "",
            url: "beyond-the-sea",
            delay: 0.167,
            avg_bar_duration: 0,
            duration: 172.47,
            tempo: 136.482,
            time_signature: 4
          },
          status: 200
        }

## POST
To add data to endpoints, you will need to use post.  Make sure that your options object includes the object you are posting in the body and `application/json` in the Content-Type header.

#### POST a new project to the list

Endpoint:

    `https://palette-picker-ac.herokuapp.com/api/v1/projects`

Request body:

    `{name: <string>}`

Example response:

    {
      "id": [
         14
        ]
    }

#### POST a new palette list

Endpoint:

    `https://ballroom-blitz.herokuapp.com/api/v1/songs`

Request body:

        {
            "title": "Melancolia Tropical",
            "spotifyid": "",
            "url": "melancolia-tropical",
            "delay": 0,
            "avbarduration": 0,
            "duration": 226,
            "tempo": 121,
            "time_signature": 4
        }

Example response:

    {
      "message": "Song created successfully!",
      "resourceId": 10,
      "status": 201
    }

## DELETE

Endpoint:

 `https://ballroom-blitz.herokuapp.com/api/v1/songs/:id`

Example response:

    None, expect a 204 status code.

## Built With:
- Golang
- Gin
- Gorm
- PostgreSQL

