package applemusic

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestMeService_GetStorefront(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v1/me/storefront", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"l": "xx"}, r)
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, `{"data":[{"id":"jp","type":"storefronts","href":"/v1/storefronts/jp","attributes":{"name":"日本","explicitContentPolicy":"allowed","supportedLanguageTags":["ja","en-US"],"defaultLanguageTag":"ja"}}]}`)
	})

	client := NewClient(httpClient)
	storefronts, _, err := client.Me.GetStorefront(&GetMeStorefrontParams{
		L: "xx",
	})

	expected := &Storefronts{
		Data: []Storefront{
			{
				Id: "jp",
				Type: "storefronts",
				Href: "/v1/storefronts/jp",
				Attributes: StorefrontAttribute{
					DefaultLanguageTag: "ja",
					Name: "日本",
					SupportedLanguageTags: []string{
						"ja",
						"en-US",
					},
					ExplicitContentPolicy: "allowed",
				},
			},
		},
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, storefronts)
}


func TestMeService_GetLibraryRecentryAdded(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v1/me/library/recently-added", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"limit": "100", "offset" : "offset"}, r)
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, `{"next":"/v1/me/library/recently-added?offset=1","data":[{"id":"l.dVaffho","type":"library-albums","href":"/v1/me/library/albums/l.dVaffho","attributes":{"trackCount":11,"playParams":{"id":"l.dVaffho","kind":"album","isLibrary":true},"artistName":"Various Artists","name":"New Invoke vol.1","artwork":{"width":1200,"height":1200,"url":"https://is5-ssl.mzstatic.com/image/thumb/Music128/v4/5a/85/02/5a85027f-b2e6-56d0-935e-60d2b88340ef/CYCLC-0002_NEW_INVOKE_VOL.1.png/{w}x{h}bb.jpeg"}}}]}`)
	})

	client := NewClient(httpClient)
	storefronts, _, err := client.Me.GetLibraryRecentryAdded(&GetMeLibraryRecentryAddedParams{
		Limit: 100,
		Offset: "offset",
	})

	expected := &LibraryAlbums{
		Data: []LibraryAlbum{
			{
				Id: "l.dVaffho",
				Type: "library-albums",
				Href: "/v1/me/library/albums/l.dVaffho",
				Attributes: AlbumAttribute{
					TrackCount: 11,
					PlayParams: PlayParams{
						Id: "l.dVaffho",
						Kind: "album",
						IsLibrary: true,
					},
					ArtistName: "Various Artists",
					Name: "New Invoke vol.1",
					Artwork: Artwork{
						Width: 1200,
						Height: 1200,
						Url: "https://is5-ssl.mzstatic.com/image/thumb/Music128/v4/5a/85/02/5a85027f-b2e6-56d0-935e-60d2b88340ef/CYCLC-0002_NEW_INVOKE_VOL.1.png/{w}x{h}bb.jpeg",
					},
				},
			},
		},
		Next: "/v1/me/library/recently-added?offset=1",

	}

	assert.Nil(t, err)
	assert.Equal(t, expected, storefronts)
}
