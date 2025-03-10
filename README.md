[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/colored.png)](#music-library-api)

# Music Library API


## Table of Contents
1. [Introduction](#-introduction)
2. [Main Features](#-main-features)
3. [Installation](#-installation)
2. [Project Structure](#-project-structure)
3. [Models](#-models)
4. [API Endpoints](#-api-endpoints)


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/colored.png)](#introduction)

## ➤ Introduction
This project is a RESTful API for managing a music library. The API allows adding, deleting, updating, and retrieving information about songs and their lyrics.


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/colored.png)](#main-features)

## ➤ Main Features

- **Add Song**: Add a new song with lyrics (or without).
- **Get List of Songs**: Retrieve a list of songs with filtering and pagination support.
- **Update Song**: Update information about a song.
- **Delete Song**: Delete a song.
- **Get Song**: Get concrete song and its associated lyrics.
- **Get Lyrics**: Get lyrics for a specific song.
- **Update Lyric**: Update the lyrics of a song.
- **Delete Lyric**: Delete lyrics for a song.
- **Add Lyric**: Add lyrics for a specific song.


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/colored.png)](#technologies)

## ➤ Technologies

- **Programming Language**: Go
- **Framework**: Gin
- **Database**: PostgreSQL
- **ORM**: GORM


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/colored.png)](#installation)

## ➤ Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/Darya-Tolmeneva/Music-Library.git
   cd Music-Library
   ```

2. **Install dependencies**:

   Make sure Go is installed. Then run:

   ```bash
   go mod download
   ```

3. **Configure the database**:

   Create a PostgreSQL database and configure the connection in the `config/config.yaml` file.

4. **Start the server**:

   ```bash
   go run ./cmd/app/main.go
   ```

   The server will be available at `http://localhost:8080`.


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/colored.png)](#models)

## ➤ Models

#### Model `Song`

| Field        | Type         | Description                                   |
|--------------|--------------|-----------------------------------------------|
| `ID`         | `uint`       | Unique identifier of the song                 |
| `Group`      | `string`     | The group that performed the song             |
| `Title`      | `string`     | The title of the song                         |
| `ReleaseDate`| `string`     | Release date                                  |
| `Link`       | `string`     | Link to the song                              |
| `CreatedAt`  | `time.Time`  | Date and time of creation                     |
| `UpdatedAt`  | `time.Time`  | Date and time of last update                  |
| `DeletedAt`  | `time.Time`  | Date and time of deletion (if applicable)     |

#### Model `Lyric`

| Field         | Type        | Description                               |
|---------------|-------------|-------------------------------------------|
| `ID`          | `uint`      | Unique identifier of the verse            |
| `SongID`      | `uint`      | ID of the song this verse is related to   |
| `Text`        | `string`    | Text of the verse                         |
| `VerseNumber` | `int`       | Number of verse                           |
| `CreatedAt`   | `time.Time` | Date and time of creation                 |
| `UpdatedAt`   | `time.Time` | Date and time of last update              |
| `DeletedAt`   | `time.Time` | Date and time of deletion (if applicable) |



[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/colored.png)](#project-structure)

## ➤ Project Structure

```
|_ cmd
     |_ app
         |_ main.go
|_ config
     |_ config.go
     |_ config.yaml
|_ internal
     |_ database
         |_ postgres
               |_ client.go
               |_ repository.go
     |_ models
         |_ moedls.go
     |_ router
         |_ router.go
     |_ transport
         |_ handlers
               |_ lyricHandlers.go
               |_ songHandlers.go
go.mod
README.md
```


[![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/colored.png)](#api-endpoints)

## ➤ API Endpoints

### Add Song

**Request**:

```bash
POST /songs

{
    "group": "The Beatles",
    "title": "Norwegian Wood",
    "releaseDate": "1965-12-03",
    "link": "https://example.com",
    "lyrics": [
        {
            "verse_number": 1,
            "text": "I once had a girl, or should I say, she once had me..."
        },
        {
            "verse_number": 2,
            "text": "She showed me her room, isn't it good, Norwegian wood?"
        }
    ]
}
```
Here you can skip lyrics.

**Response**:

```json
{
   "song": {
      "CreatedAt": "2025-03-10T18:04:34.552783942+03:00",
      "UpdatedAt": "2025-03-10T18:04:34.552783942+03:00",
      "DeletedAt": null,
      "ID": 12,
      "group": "The Beatles",
      "title": "Norwegian Wood",
      "releaseDate": "1965-12-03",
      "link": "https://example.com",
      "lyrics": [
         {
            "CreatedAt": "2025-03-10T18:04:34.562256089+03:00",
            "UpdatedAt": "2025-03-10T18:04:34.562256089+03:00",
            "DeletedAt": null,
            "ID": 21,
            "song_id": 12,
            "verse_number": 1,
            "text": "I once had a girl, or should I say, she once had me..."
         },
         {
            "CreatedAt": "2025-03-10T18:04:34.562256089+03:00",
            "UpdatedAt": "2025-03-10T18:04:34.562256089+03:00",
            "DeletedAt": null,
            "ID": 22,
            "song_id": 12,
            "verse_number": 2,
            "text": "She showed me her room, isn't it good, Norwegian wood?"
         }
      ]
   }
}
```

### Get List of Songs

**Request**:

```bash
GET /songs?group=The Beatles&offset=1&page_size=10
```
Here you can specify offset and page_size. Also group name, title, release date (in format "2005-01-01") and link.

**Response**:

```json
{
   "data": [
      {
         "CreatedAt": "2025-03-06T23:16:40.736127+03:00",
         "UpdatedAt": "2025-03-06T23:16:40.736127+03:00",
         "DeletedAt": null,
         "ID": 2,
         "group": "The Beatles",
         "title": "Norwegian wood",
         "releaseDate": "2005-01-01",
         "link": "meow",
         "lyrics": []
      },
      {
         "CreatedAt": "2025-03-07T09:46:29.29592+03:00",
         "UpdatedAt": "2025-03-07T09:46:29.29592+03:00",
         "DeletedAt": null,
         "ID": 9,
         "group": "The Beatles",
         "title": "Norwegian Wood",
         "releaseDate": "1965-12-03",
         "link": "https://example.com",
         "lyrics": [
            {
               "CreatedAt": "2025-03-07T09:46:29.297095+03:00",
               "UpdatedAt": "2025-03-07T09:46:29.297095+03:00",
               "DeletedAt": null,
               "ID": 6,
               "song_id": 9,
               "verse_number": 1,
               "text": "I once had a girl, or should I say, she once had me..."
            }
         ]
      },
      {
         "CreatedAt": "2025-03-07T09:47:01.348434+03:00",
         "UpdatedAt": "2025-03-07T09:47:01.348434+03:00",
         "DeletedAt": null,
         "ID": 10,
         "group": "The Beatles",
         "title": "Norwegian Wood",
         "releaseDate": "1965-12-03",
         "link": "https://example.com",
         "lyrics": []
      },
      {
         "CreatedAt": "2025-03-07T10:00:19.772603+03:00",
         "UpdatedAt": "2025-03-07T10:00:19.772603+03:00",
         "DeletedAt": null,
         "ID": 11,
         "group": "The Beatles",
         "title": "Norwegian Wood",
         "releaseDate": "1965-12-03",
         "link": "https://example.com",
         "lyrics": [
            {
               "CreatedAt": "2025-03-07T10:00:19.774201+03:00",
               "UpdatedAt": "2025-03-10T16:49:11.238136+03:00",
               "DeletedAt": null,
               "ID": 7,
               "song_id": 11,
               "verse_number": 1,
               "text": "meow"
            },
            {
               "CreatedAt": "2025-03-10T16:50:15.198218+03:00",
               "UpdatedAt": "2025-03-10T16:50:15.198218+03:00",
               "DeletedAt": null,
               "ID": 20,
               "song_id": 11,
               "verse_number": 2,
               "text": "meow"
            }
         ]
      },
      {
         "CreatedAt": "2025-03-10T18:04:34.552783+03:00",
         "UpdatedAt": "2025-03-10T18:04:34.552783+03:00",
         "DeletedAt": null,
         "ID": 12,
         "group": "The Beatles",
         "title": "Norwegian Wood",
         "releaseDate": "1965-12-03",
         "link": "https://example.com",
         "lyrics": [
            {
               "CreatedAt": "2025-03-10T18:04:34.562256+03:00",
               "UpdatedAt": "2025-03-10T18:04:34.562256+03:00",
               "DeletedAt": null,
               "ID": 21,
               "song_id": 12,
               "verse_number": 1,
               "text": "I once had a girl, or should I say, she once had me..."
            },
            {
               "CreatedAt": "2025-03-10T18:04:34.562256+03:00",
               "UpdatedAt": "2025-03-10T18:04:34.562256+03:00",
               "DeletedAt": null,
               "ID": 22,
               "song_id": 12,
               "verse_number": 2,
               "text": "She showed me her room, isn't it good, Norwegian wood?"
            }
         ]
      },
      {
         "CreatedAt": "2025-03-10T18:05:47.618165+03:00",
         "UpdatedAt": "2025-03-10T18:05:47.618165+03:00",
         "DeletedAt": null,
         "ID": 13,
         "group": "The Beatles",
         "title": "Norwegian Wood",
         "releaseDate": "1965-12-03",
         "link": "https://example.com",
         "lyrics": [
            {
               "CreatedAt": "2025-03-10T18:05:47.618779+03:00",
               "UpdatedAt": "2025-03-10T18:05:47.618779+03:00",
               "DeletedAt": null,
               "ID": 23,
               "song_id": 13,
               "verse_number": 1,
               "text": "I once had a girl, or should I say, she once had me..."
            },
            {
               "CreatedAt": "2025-03-10T18:05:47.618779+03:00",
               "UpdatedAt": "2025-03-10T18:05:47.618779+03:00",
               "DeletedAt": null,
               "ID": 24,
               "song_id": 13,
               "verse_number": 2,
               "text": "She showed me her room, isn't it good, Norwegian wood?"
            }
         ]
      }
   ],
   "pagination": {
      "offset": 1,
      "page_size": 10,
      "total": 7
   }
}
```

### Update Song

**Request**:

```bash
PUT /songs/{id}
```
```json
{
    "title": "Norwegian Wood (Updated)"
}
```
Here you can update one parameter or many.

**Response**:

```json
{
   "song": {
      "CreatedAt": "2025-03-07T10:00:19.772603+03:00",
      "UpdatedAt": "2025-03-10T22:09:40.25671+03:00",
      "DeletedAt": null,
      "ID": 11,
      "group": "The Beatlesss",
      "title": "Norwegian Wood (Updated)",
      "release_date": "1965-12-03",
      "link": "https://example.com",
      "lyrics": [
         {
            "CreatedAt": "2025-03-07T10:00:19.774201+03:00",
            "UpdatedAt": "2025-03-10T16:49:11.238136+03:00",
            "DeletedAt": null,
            "ID": 7,
            "song_id": 11,
            "verse_number": 1,
            "text": "meow"
         },
         {
            "CreatedAt": "2025-03-10T16:50:15.198218+03:00",
            "UpdatedAt": "2025-03-10T16:50:15.198218+03:00",
            "DeletedAt": null,
            "ID": 20,
            "song_id": 11,
            "verse_number": 2,
            "text": "meow"
         }
      ]
   }
}
```

### Delete Song

**Request**:

```bash
DELETE /songs/{id}
```

**Response**:
```json
{
    "deleted_song_id": 11
}
```

It is also delete Lyrics which have this song_id.

### Update Song

**Request**:

```bash
GET /songs/{id}
```

**Response**:
```json
{
    "song": {
        "CreatedAt": "2025-03-10T18:05:47.618165+03:00",
        "UpdatedAt": "2025-03-10T18:05:47.618165+03:00",
        "DeletedAt": null,
        "ID": 13,
        "group": "The Beatles",
        "title": "Norwegian Wood",
        "release_date": "1965-12-03",
        "link": "https://example.com",
        "lyrics": [
            {
                "CreatedAt": "2025-03-10T18:05:47.618779+03:00",
                "UpdatedAt": "2025-03-10T18:05:47.618779+03:00",
                "DeletedAt": null,
                "ID": 23,
                "song_id": 13,
                "verse_number": 1,
                "text": "I once had a girl, or should I say, she once had me..."
            },
            {
                "CreatedAt": "2025-03-10T18:05:47.618779+03:00",
                "UpdatedAt": "2025-03-10T18:05:47.618779+03:00",
                "DeletedAt": null,
                "ID": 24,
                "song_id": 13,
                "verse_number": 2,
                "text": "She showed me her room, isn't it good, Norwegian wood?"
            }
        ]
    }
}
```

### Get Lyric

**Request**:

```bash
GET /lyrics/{id}
```

**Response**:

```json
{
   "lyric": {
      "CreatedAt": "2025-03-10T18:05:47.618779+03:00",
      "UpdatedAt": "2025-03-10T18:05:47.618779+03:00",
      "DeletedAt": null,
      "ID": 23,
      "song_id": 13,
      "verse_number": 1,
      "text": "I once had a girl, or should I say, she once had me..."
   }
}
```

### Add Lyric

**Request**:

```bash
POST /lyrics
{
    "song_id": 1,
    "verse_number": 3,
    "text": "This is the third verse of the song."
}
```

**Response**:

```json
{
   "lyric": {
      "CreatedAt": "2025-03-10T22:21:37.587893086+03:00",
      "UpdatedAt": "2025-03-10T22:21:37.587893086+03:00",
      "DeletedAt": null,
      "ID": 25,
      "song_id": 13,
      "verse_number": 3,
      "text": "This is the third verse of the song."
   }
}
```

### Delete Lyric

**Request**:

```bash
DELETE /lyrics/{id}
```

**Response**:

```json
{
   "deleted_lyric_id": 25
}
```

### Update Lyric

**Request**:

```bash
PUT /lyrics/{id}
{
    "song_id": 12,
    "verse_number": 3,
    "text": "This is the third verse of the song."
}
```

**Response**:

```json
{
   "lyric": {
      "CreatedAt": "2025-03-10T18:05:47.618779+03:00",
      "UpdatedAt": "2025-03-10T22:23:36.303528+03:00",
      "DeletedAt": null,
      "ID": 24,
      "song_id": 12,
      "verse_number": 3,
      "text": "This is the third verse of the song."
   }
}
```
