package api

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/noppefoxwolf/amg/applemusic/models"
	"github.com/stretchr/testify/assert"
)

func TestCatalogService_GetAlbum(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v1/catalog/us/albums/310730204", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"l": "xx"}, r)
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, `{
    "data": [
        {
            "attributes": {
                "artistName": "Bruce Springsteen",
                "artwork": {
                    "bgColor": "ffffff",
                    "height": 1500,
                    "textColor1": "0c0b09",
                    "textColor2": "2a2724",
                    "textColor3": "3d3c3a",
                    "textColor4": "555250",
                    "url": "https://example.mzstatic.com/image/thumb/Music3/v4/2d/02/4a/2d024aaa-4547-ca71-7ba1-b8f5e1d98256/source/{w}x{h}bb.jpeg",
                    "width": 1500
                },
                "copyright": "\u2117 1975 Bruce Springsteen",
                "editorialNotes": {
                    "short": "Springsteen's third album was the one that broke it all open for him.",
                    "standard": "Springsteen's third album was the one that broke it all open for him, planting his tales of Jersey girls, cars, and nights spent sleeping on the beach firmly in the Top Five. He shot for an unholy hybrid of Orbison, Dylan and Spector \u2014 and actually reached it. \"Come take my hand,\" he invited in the opening lines. \"We're ridin' out tonight to case the Promised Land.\" Soon after this album, he'd discover the limits of such dreams, but here, it's a wide-open road: Even the tales of petty crime (\"Meeting Across the River\") and teen-gang violence (\"Jungleland\") are invested with all the wit and charm you can handle. Bruce's catalog is filled with one-of-a-kind albums from <i>The Wild, The Innocent and the E Street Shuffle</i> to <i>The Ghost of Tom Joad</i>. Forty years on, <i>Born to Run</i> still sits near the very top of that stack."
                },
                "genreNames": [
                    "Rock",
                    "Music",
                    "Arena Rock",
                    "Rock & Roll",
                    "Pop",
                    "Pop/Rock"
                ],
                "isComplete": true,
                "isMasteredForItunes": true,
                "isSingle": false,
                "name": "Born to Run",
                "playParams": {
                    "id": "310730204",
                    "kind": "album"
                },
                "recordLabel": "Columbia",
                "releaseDate": "1975-08-25",
                "trackCount": 8,
                "url": "https://itunes.apple.com/us/album/born-to-run/310730204"
            },
            "href": "/v1/catalog/us/albums/310730204",
            "id": "310730204",
            "relationships": {
                "artists": {
                    "data": [
                        {
                            "href": "/v1/catalog/us/artists/178834",
                            "id": "178834",
                            "type": "artists"
                        }
                    ],
                    "href": "/v1/catalog/us/albums/310730204/artists"
                },
                "tracks": {
                    "data": [
                        {
                            "attributes": {
                                "artistName": "Bruce Springsteen",
                                "artwork": {
                                    "bgColor": "ffffff",
                                    "height": 1500,
                                    "textColor1": "0c0b09",
                                    "textColor2": "2a2724",
                                    "textColor3": "3d3c3a",
                                    "textColor4": "555250",
                                    "url": "https://example.mzstatic.com/image/thumb/Music3/v4/2d/02/4a/2d024aaa-4547-ca71-7ba1-b8f5e1d98256/source/{w}x{h}bb.jpeg",
                                    "width": 1500
                                },
                                "composerName": "Bruce Springsteen",
                                "discNumber": 1,
                                "durationInMillis": 289186,
                                "genreNames": [
                                    "Rock",
                                    "Music",
                                    "Arena Rock",
                                    "Rock & Roll",
                                    "Pop",
                                    "Pop/Rock"
                                ],
                                "isrc": "USSM19904335",
                                "name": "Thunder Road",
                                "playParams": {
                                    "id": "310730206",
                                    "kind": "song"
                                },
                                "previews": [
                                    {
                                        "url": "https://example.itunes.apple.com/apple-assets-us-std-000001/Music4/v4/a2/4f/a7/a24fa727-36e4-c870-973b-00704c80187a/mzaf_7083782255622091857.plus.aac.p.m4a"
                                    }
                                ],
                                "releaseDate": "1975-08-25",
                                "trackNumber": 1,
                                "url": "https://itunes.apple.com/us/album/thunder-road/310730204?i=310730206"
                            },
                            "href": "/v1/catalog/us/songs/310730206",
                            "id": "310730206",
                            "type": "songs"
                        }
                    ],
                    "href": "/v1/catalog/us/albums/310730204/tracks"
                }
            },
            "type": "albums"
        }
    ]
}`)
	})

	client := NewClient(httpClient)
	storefronts, _, err := client.Catalog.GetAlbum(&GetAlbumParams{
		Storefront: "us",
		Id:         "310730204",
		L:          "xx",
	})

	locale, _ := time.LoadLocation("UTC")
	expected := &models.VariableResponse{
		Data: []models.VariableResource{
			{
				Type: "albums",
				Album: &models.Album{
					Id:   "310730204",
					Href: "/v1/catalog/us/albums/310730204",
					Attributes: models.AlbumAttributes{
						ArtistName: "Bruce Springsteen",
						Artwork: models.Artwork{
							BgColor:    "ffffff",
							Width:      1500,
							Height:     1500,
							TextColor1: "0c0b09",
							TextColor2: "2a2724",
							TextColor3: "3d3c3a",
							TextColor4: "555250",
							Url:        "https://example.mzstatic.com/image/thumb/Music3/v4/2d/02/4a/2d024aaa-4547-ca71-7ba1-b8f5e1d98256/source/{w}x{h}bb.jpeg",
						},
						Copyright: "℗ 1975 Bruce Springsteen",
						EditorialNotes: models.EditorialNotes{
							Short:    "Springsteen's third album was the one that broke it all open for him.",
							Standard: "Springsteen's third album was the one that broke it all open for him, planting his tales of Jersey girls, cars, and nights spent sleeping on the beach firmly in the Top Five. He shot for an unholy hybrid of Orbison, Dylan and Spector — and actually reached it. \"Come take my hand,\" he invited in the opening lines. \"We're ridin' out tonight to case the Promised Land.\" Soon after this album, he'd discover the limits of such dreams, but here, it's a wide-open road: Even the tales of petty crime (\"Meeting Across the River\") and teen-gang violence (\"Jungleland\") are invested with all the wit and charm you can handle. Bruce's catalog is filled with one-of-a-kind albums from <i>The Wild, The Innocent and the E Street Shuffle</i> to <i>The Ghost of Tom Joad</i>. Forty years on, <i>Born to Run</i> still sits near the very top of that stack.",
						},
						GenreNames: []string{
							"Rock",
							"Music",
							"Arena Rock",
							"Rock & Roll",
							"Pop",
							"Pop/Rock",
						},
						IsComplete: true,
						Name:       "Born to Run",
						PlayParams: models.PlayParameters{
							Id:        "310730204",
							Kind:      "album",
							IsLibrary: false,
						},
						RecordLabel: "Columbia",
						ReleaseDate: models.ReleaseDate{
							Time: time.Date(1975, 8, 25, 0, 0, 0, 0, locale),
						},
						TrackCount:          8,
						Url:                 "https://itunes.apple.com/us/album/born-to-run/310730204",
						IsMasteredForItunes: true,
					},
					Relationships: models.AlbumRelationships{
						Artists: models.ArtistRelationship{
							Data: []models.Artist{
								{
									Href: "/v1/catalog/us/artists/178834",
									Id:   "178834",
								},
							},
						},
						Tracks: models.TrackRelationship{
							Data: []models.SongOrMusicVideoResource{
								{
									Type: "songs",
									Song: &models.Song{
										Href: "/v1/catalog/us/songs/310730206",
										Id:   "310730206",
										Type: "songs",
										Attributes: models.SongAttributes{
											ArtistName:   "Bruce Springsteen",
											ComposerName: "Bruce Springsteen",
											Artwork: models.Artwork{
												BgColor:    "ffffff",
												Height:     1500,
												Width:      1500,
												TextColor1: "0c0b09",
												TextColor2: "2a2724",
												TextColor3: "3d3c3a",
												TextColor4: "555250",
												Url:        "https://example.mzstatic.com/image/thumb/Music3/v4/2d/02/4a/2d024aaa-4547-ca71-7ba1-b8f5e1d98256/source/{w}x{h}bb.jpeg",
											},
											DiscNumber:       1,
											DurationInMillis: 289186,
											EditorialNotes:   models.EditorialNotes{},
											GenreNames: []string{
												"Rock",
												"Music",
												"Arena Rock",
												"Rock & Roll",
												"Pop",
												"Pop/Rock",
											},
											Isrc: "USSM19904335",
											Name: "Thunder Road",
											PlayParams: models.PlayParameters{
												Id:        "310730206",
												Kind:      "song",
												IsLibrary: false,
											},
											Previews: []models.Preview{
												{
													Url: "https://example.itunes.apple.com/apple-assets-us-std-000001/Music4/v4/a2/4f/a7/a24fa727-36e4-c870-973b-00704c80187a/mzaf_7083782255622091857.plus.aac.p.m4a",
												},
											},
											ReleaseDate: models.ReleaseDate{
												Time: time.Date(1975, 8, 25, 0, 0, 0, 0, locale),
											},
											TrackNumber: 1,
											Url:         "https://itunes.apple.com/us/album/thunder-road/310730204?i=310730206",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, storefronts)
}

func TestCatalogService_GetAlbumRelationship(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v1/catalog/us/albums/310730204/artists", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"l": "xx", "limit": "100"}, r)
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, `{
    "data": [
        {
            "attributes": {
                "genreNames": [
                    "Rock"
                ],
                "name": "Bruce Springsteen",
                "url": "https://itunes.apple.com/us/artist/bruce-springsteen/178834"
            },
            "href": "/v1/catalog/us/artists/178834",
            "id": "178834",
            "type": "artists"
        }
    ]
}`)
	})

	client := NewClient(httpClient)
	storefronts, _, err := client.Catalog.GetAlbumRelationship(&GetAlbumRelationshipParams{
		Id:           "310730204",
		Storefront:   "us",
		Relationship: "artists",
		L:            "xx",
		Limit:        100,
	})

	expected := &models.VariableResponse{
		Data: []models.VariableResource{
			{
				Type: "artists",
				Artist: &models.Artist{
					Href: "/v1/catalog/us/artists/178834",
					Id:   "178834",
					Attributes: models.ArtistAttributes{
						GenreNames: []string{"Rock"},
						Name:       "Bruce Springsteen",
						Url:        "https://itunes.apple.com/us/artist/bruce-springsteen/178834",
					},
				},
			},
		},
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, storefronts)
}

func TestCatalogService_GetMultipleAlbums(t *testing.T) {

	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v1/catalog/us/albums", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"ids": "310730204,19075891"}, r)
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, `{
    "data": [
        {
            "type": "albums"
        }
	]}`)
	})

	client := NewClient(httpClient)
	storefronts, _, err := client.Catalog.GetMultipleAlbums(&GetMultipleAlbumsParams{
		Storefront: "us",
		Ids:        []string{"310730204", "19075891"},
	})

	expected := &models.VariableResponse{
		Data: []models.VariableResource{
			{
				Type:  "albums",
				Album: &models.Album{},
			},
		},
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, storefronts)
}
