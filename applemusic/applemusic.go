package applemusic

import "net/http"
import "github.com/dghubble/sling"

const applemusicAPI = "https://api.music.apple.com/v1/"

type Client struct {
	sling       *sling.Sling
	Storefronts *StorefrontsService
	Me          *MeService
}

func NewClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(applemusicAPI)
	return &Client{
		sling:       base,
		Storefronts: newStorefrontsService(base.New()),
		Me:          newMeService(base.New()),
	}
}
