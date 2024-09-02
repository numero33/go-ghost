// Package ghost provides the binding for Ghost APIs
package ghost

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	URL       string
	Key       string
	Version   string
	GhostPath string
	UserAgent string
	client    *http.Client
}

// defaultHTTPTimeout is the default http.Client timeout.
const defaultHTTPTimeout = 10 * time.Second

// NewClient creates a new Ghost API client. URL and Key are explained in
// the Ghost API docs: https://ghost.org/docs/api/v2/content/#authentication
func NewClient(url, key string) *Client {
	httpClient := &http.Client{Timeout: defaultHTTPTimeout}
	return &Client{
		URL:       url,
		Key:       key,
		Version:   "v3.0",
		GhostPath: "ghost",
		client:    httpClient,
	}
}

// generateJWT follows the Token generation alogrithm outlined in Ghost token authentication.
//
// https://ghost.org/docs/admin-api/#token-authentication
func (c *Client) generateJWT() (string, error) {
	keyParts := strings.Split(c.Key, ":")
	if len(keyParts) != 2 {
		return "", fmt.Errorf("Invalid Client.Key format")
	}

	secret := make([]byte, hex.DecodedLen(len(keyParts[1])))
	_, err := hex.Decode(secret, []byte(keyParts[1]))
	if err != nil {
		return "", err
	}

	now := time.Now()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"aud": "/admin/",
		"exp": now.Add(5 * time.Minute).Unix(),
		"iat": now.Unix(),
	})

	token.Header["kid"] = keyParts[0]

	return token.SignedString(secret)
}

func (c *Client) NewRequest(method, path string, data interface{}) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", c.URL, path)

	var body io.Reader
	if data != nil {
		b := &bytes.Buffer{}
		if err := json.NewEncoder(b).Encode(data); err != nil {
			return nil, fmt.Errorf("Request: %v", err)
		}
		body = b
	}

	r, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("buildRequest: %v", err)
	}

	ua := c.UserAgent
	if ua == "" {
		ua = "go-ghost v1"
	}
	r.Header.Add("User-Agent", ua)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept-Version", c.Version)
	if c.Key != "" {
		token, err := c.generateJWT()
		if err != nil {
			return nil, err
		}
		r.Header.Add("Authorization", "Ghost "+token)
	}

	return r, nil
}

// Request makes an API request to Ghost with the given HTTP method, HTTP Path,
// and data marshaled to JSON in the body. The JSON object needs to follow the
// structure of all Ghost Request/Response objects.
//
// https://ghost.org/docs/admin-api/#json-format
func (c *Client) Request(method, path string, data interface{}) (*http.Response, error) {

	r, err := c.NewRequest(method, path, data)
	if err != nil {
		return nil, err
	}

	return c.client.Do(r)
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}

// Endpoint is a HTTP Path slug generator for a Ghost resource. e.g. "admin", "post", "abc123".
func (c *Client) Endpoint(api, resource string) string {
	return fmt.Sprintf("/%s/api/%s/%s/", c.GhostPath, api, resource)
}

// EndpointForID is a HTTP Path slug generator for a Ghost resource with an ID. e.g. "admin", "post", "abc123".
func (c *Client) EndpointForID(api, resource, id string) string {
	return c.Endpoint(api, resource) + fmt.Sprintf("%s/", id)
}

// EndpointForSlug is a HTTP Path slug generator for a Ghost resource with a slug. e.g. "admin", "post", "abc123".
func (c *Client) EndpointForSlug(api, resource, slug string) string {
	return c.Endpoint(api, resource) + fmt.Sprintf("slug/%s/", slug)
}

// String returns a pointer to the string value passed in.
func String(v string) *string {
	return &v
}

// Bool returns a pointer to the bool value passed in.
func Bool(v bool) *bool {
	return &v
}
