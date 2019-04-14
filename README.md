# amg
Apple Music API client written in golang.

# Usage

```
func (t *oauthTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("Authorization", "Bearer xxxxx")
	return http.DefaultTransport.RoundTrip(r)
}

c := &http.Client{Transport: &oauthTransport{}}
c2 := applemusic.NewClient(c)
params := &applemusic.GetAllStorefrontsParams{}
res, req, err := c2.StorefrontsAndLocalization.GetAllStorefronts(params)
```
