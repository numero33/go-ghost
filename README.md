# go-ghost
Simple `ghost` api client for golang

> Heavily inspired by [philips/go-ghost](https://github.com/philips/go-ghost)

## Installation
```bash
go get -u github.com/numero33/go-ghost
```

## Usage
```go
package main

import (
	"github.com/numero33/go-ghost
)

func main() {
	client := ghost.NewClient("<URL>", "<API_KEY>")

	resp, err := client.Request(http.MethodGet, client.Endpoint("admin", "posts"), nil)
	if err != nil {
		log.Fatalf("get: %v", err)
	}
	defer resp.Body.Close()

	pr := &ghost.PostRequest{}
	if err = json.NewDecoder(resp.Body).Decode(pr); err != nil {
		log.Fatalf("decode: %v", err)
	}
	for _, post := range pr.Posts {
		fmt.Printf("title: %v Updated At: %v\n", *post.Title, post.UpdatedAt.String())
	}
}

```

### Admin API
<details>
  <summary>Get Posts</summary>

```go
package main

import (
	"github.com/numero33/go-ghost
)

func main() {
	client := ghost.NewClient("<URL>", "<API_KEY>")

	resp, err := client.Request(http.MethodGet, client.Endpoint("admin", "posts"), nil)
	if err != nil {
		log.Fatalf("get: %v", err)
	}
	defer resp.Body.Close()

	pr := &ghost.PostRequest{}
	if err = json.NewDecoder(resp.Body).Decode(pr); err != nil {
		log.Fatalf("decode: %v", err)
	}
	for _, post := range pr.Posts {
		fmt.Printf("title: %v Updated At: %v\n", *post.Title, post.UpdatedAt.String())
	}
}

```
</details>

<details>
  <summary>Get Posts with Parameters</summary>

```go
package main

import (
	"github.com/numero33/go-ghost
)

func main() {
	client := ghost.NewClient("<URL>", "<API_KEY>")

	req, err := client.NewRequest(http.MethodGet, client.Endpoint("admin", "posts"), nil)
	if err != nil {
		log.Fatalf("get: %v", err)
	}

	// https://ghost.org/docs/content-api/#parameters
	q := req.URL.Query()
    
	// https://ghost.org/docs/content-api/#formats
	q.Add("formats", "plaintext,html")

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("get: %v", err)
	}
	defer resp.Body.Close()

	pr := &ghost.PostRequest{}
	if err = json.NewDecoder(resp.Body).Decode(pr); err != nil {
		log.Fatalf("decode: %v", err)
	}
	for _, post := range pr.Posts {
		fmt.Printf("title: %v Updated At: %v\n", *post.Title, post.UpdatedAt.String())
	}
}

```
</details>

<details>
  <summary>Create Post</summary>

```go
package main

import (
	"github.com/numero33/go-ghost
)

func main() {
	client := ghost.NewClient("<URL>", "<API_KEY>")

	md := `This is a **test post** made with the [go-ghost](https://github.com/numero33/go-ghost) library.`

	mobiledoc := `{ "version": "0.3.1", "atoms": [], "cards": [["markdown", { "markdown": "` + md + `" }]], "markups": [], "sections": [[10, 0]] }`

	presp := ghost.PostRequest{
		Posts: []ghost.Post{
			ghost.Post{
				Title:     ghost.String("Test Post"),
				Mobiledoc: ghost.String(mobiledoc),
			}},
	}

	resp, err := client.Request(http.MethodPost, client.Endpoint("admin", "posts"), presp)
	if err != nil {
		log.Fatalf("post: %v", err)
	}

	defer resp.Body.Close()

	pr := &ghost.PostRequest{}
	if err = json.NewDecoder(resp.Body).Decode(pr); err != nil {
		log.Fatalf("decode: %v", err)
	}
	for _, post := range pr.Posts {
		fmt.Printf("title: %v Updated At: %v\n", *post.Title, post.UpdatedAt.String())
	}
}

```
</details>



