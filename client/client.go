package client

import (
	"fmt"
	"net/http"

	h "github.com/arjendevos/instagram-client/helpers"
	"github.com/rs/zerolog/log"
)

type Client struct {
	mainUrl  string
	cookie   string
	xIgAppId string
}

func NewClient() *Client {
	cookie := h.GetEnv("COOKIE")
	xIgAppId := h.GetEnv("X_IG_APP_ID")
	return &Client{mainUrl: "https://i.instagram.com/api/v1", cookie: cookie, xIgAppId: xIgAppId}
}

func (c *Client) createUrl(path string) string {
	return c.mainUrl + path
}

func (c *Client) get(path string) (*http.Response, error) {
	client := http.Client{}
	url := c.createUrl(path)
	log.Info().Msg(fmt.Sprintf("GET %v", url))
	// fmt.Printf("GET %v\n", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"Accept":       {"*/*"},
		"cookie":       {c.cookie},
		"origin":       {"https://www.instagram.com"},
		"referer":      {"https://www.instagram.com/"},
		"authority":    {"i.instagram.com"},
		"x-ig-app-id":  {c.xIgAppId},
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
