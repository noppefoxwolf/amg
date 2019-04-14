package main

import (
	"log"
	"net/http"

	"github.com/noppefoxwolf/amg/applemusic"
)

type oauthTransport struct{}

func (t *oauthTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("Authorization", "Bearer xxxxx")
	return http.DefaultTransport.RoundTrip(r)
}

func main() {
	c := &http.Client{Transport: &oauthTransport{}}
	c2 := applemusic.NewClient(c)
	params := &applemusic.GetAllStorefrontsParams{}
	res, req, err := c2.StorefrontsAndLocalization.GetAllStorefronts(params)
	log.Print(req)
	if err != nil {
		log.Print(err)
	} else {
		log.Print(res.Data[0].Type)
		log.Print(res.Data[0].Storefront)
	}
}
