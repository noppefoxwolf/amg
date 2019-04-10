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
