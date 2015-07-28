package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

var (
	oauthConf = &oauth2.Config{
		ClientID:     "ffffffffffffffffffff",
		ClientSecret: "ffffffffffffffffffffffffffffffffffffffff",
		Scopes:       []string{"user:email"},
		Endpoint:     githuboauth.Endpoint,
	}
	oauthStateString = "randomstring"
)

func UnauthedGitHubClient(transport http.RoundTripper) *github.Client {
	t := github.UnauthenticatedRateLimitedTransport{
		ClientID:     oauthConf.ClientID,
		ClientSecret: oauthConf.ClientSecret,
		Transport:    transport,
	}
	return github.NewClient(t.Client())
}

type ETagTransport struct {
	ETag string
}

func (t *ETagTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.ETag == "" {
		return nil, errors.New("t.ETag is empty")
	}

	// To set extra querystring params, we must make a copy of the Request so
	// that we don't modify the Request we were given. This is required by the
	// specification of http.RoundTripper.
	req = cloneRequest(req)
	req.Header.Add("If-None-Match", t.ETag)

	// Make the HTTP request.
	resp, err := http.DefaultTransport.RoundTrip(req)
	log.Println(resp.Header.Get("ETag"))
	return resp, err
}

// cloneRequest returns a clone of the provided *http.Request. The clone is a
// shallow copy of the struct and its Header map.
func cloneRequest(r *http.Request) *http.Request {
	// shallow copy of the struct
	r2 := new(http.Request)
	*r2 = *r
	// deep copy of the Header
	r2.Header = make(http.Header)
	for k, s := range r.Header {
		r2.Header[k] = s
	}
	return r2
}

func main() {
	t := &ETagTransport{ETag: "\"ca6a7c00fe71845258d264496393cbd3\""}
	c := UnauthedGitHubClient(t)
	es, _, err := c.Activity.ListEventsPerformedByUser("samertm", true, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(len(es))
}
