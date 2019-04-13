package api

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/noppefoxwolf/amg/applemusic/models"

	"github.com/stretchr/testify/assert"
)

func TestStorefrontsService_GetAll(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v1/storefronts/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"l": "xx", "limit": "100", "offset": "yy"}, r)
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, `{"data":[{"id":"ai","type":"storefronts","href":"/v1/storefronts/ai","attributes":{"defaultLanguageTag":"en-GB","name":"Anguilla","supportedLanguageTags":["en-GB"],"explicitContentPolicy":"allowed"}}]}`)
	})

	client := NewClient(httpClient)
	storefronts, _, err := client.Storefronts.GetAll(&GetAllStorefrontsParams{
		L:      "xx",
		Limit:  100,
		Offset: "yy",
	})

	expected := &models.StorefrontResponse{
		Data: []models.Storefront{
			{
				Id:   "ai",
				Href: "/v1/storefronts/ai",
				Attributes: models.StorefrontAttributes{
					DefaultLanguageTag: "en-GB",
					Name:               "Anguilla",
					SupportedLanguageTags: []string{
						"en-GB",
					},
					ExplicitContentPolicy: "allowed",
				},
			},
		},
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, storefronts)
}

func TestStorefrontsService_GetMultipleStorefronts(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v1/storefronts/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"l": "xx", "ids": "ja,tw"}, r)
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, `{"data":[{"id":"ai","type":"storefronts","href":"/v1/storefronts/ai","attributes":{"defaultLanguageTag":"en-GB","name":"Anguilla","supportedLanguageTags":["en-GB"],"explicitContentPolicy":"allowed"}}]}`)
	})

	client := NewClient(httpClient)
	storefronts, _, err := client.Storefronts.GetMultipleStorefronts(&GetMultipleStorefrontsParams{
		L:   "xx",
		Ids: []string{"ja", "tw"},
	})

	expected := &models.StorefrontResponse{
		Data: []models.Storefront{
			{
				Id:   "ai",
				Href: "/v1/storefronts/ai",
				Attributes: models.StorefrontAttributes{
					DefaultLanguageTag: "en-GB",
					Name:               "Anguilla",
					SupportedLanguageTags: []string{
						"en-GB",
					},
					ExplicitContentPolicy: "allowed",
				},
			},
		},
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, storefronts)
}

func TestStorefrontsService_GetStorefront(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v1/storefronts/jp", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, `{"data":[{"id":"jp","type":"storefronts","href":"/v1/storefronts/jp","attributes":{"defaultLanguageTag":"ja","supportedLanguageTags":["ja","en-US"],"explicitContentPolicy":"allowed","name":"Japan"}}]}`)
	})

	client := NewClient(httpClient)
	storefronts, _, err := client.Storefronts.GetStorefront(&GetStorefrontParams{
		Id: "jp",
	})

	expected := &models.StorefrontResponse{
		Data: []models.Storefront{
			{
				Id:   "jp",
				Href: "/v1/storefronts/jp",
				Attributes: models.StorefrontAttributes{
					DefaultLanguageTag: "ja",
					Name:               "Japan",
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
