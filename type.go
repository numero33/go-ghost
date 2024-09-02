package ghost

import (
	"time"
)

type Author struct {
	ID              *string `json:"id"`
	Name            *string `json:"name"`
	Slug            *string `json:"slug"`
	ProfileImage    *string `json:"profile_image"`
	CoverImage      *string `json:"cover_image"`
	Bio             *string `json:"bio"`
	Website         *string `json:"website"`
	Location        *string `json:"location"`
	Facebook        *string `json:"facebook"`
	Twitter         *string `json:"twitter"`
	MetaTitle       *string `json:"meta_title"`
	MetaDescription *string `json:"meta_description"`
	URL             *string `json:"url"`
}

type PostRequest struct {
	Posts  []Post  `json:"posts,omitempty"`
	Errors []Error `json:"errors,omitempty"`
}

type Post struct {
	ID                 *string    `json:"id,omitempty"`
	UUID               *string    `json:"uuid,omitempty"`
	Title              *string    `json:"title,omitempty"`
	Slug               *string    `json:"slug,omitempty"`
	HTML               *string    `json:"html,omitempty"`
	Mobiledoc          *string    `json:"mobiledoc,omitempty"`
	CommentID          *string    `json:"comment_id,omitempty"`
	FeatureImage       *string    `json:"feature_image,omitempty"`
	Featured           *bool      `json:"featured,omitempty"`
	Page               *bool      `json:"page,omitempty"`
	MetaTitle          *string    `json:"meta_title,omitempty"`
	MetaDescription    *string    `json:"meta_description,omitempty"`
	CreatedAt          *time.Time `json:"created_at,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
	PublishedAt        *time.Time `json:"published_at,omitempty"`
	CustomExcerpt      *string    `json:"custom_excerpt,omitempty"`
	OGImage            *string    `json:"og_image,omitempty"`
	OGTitle            *string    `json:"og_title,omitempty"`
	OGDescription      *string    `json:"og_description,omitempty"`
	TwitterImage       *string    `json:"twitter_image,omitempty"`
	TwitterTitle       *string    `json:"twitter_title,omitempty"`
	TwitterDescription *string    `json:"twitter_description,omitempty"`
	CustomTemplate     *string    `json:"custom_template,omitempty"`
	PrimaryAuthor      *Author    `json:"primary_author,omitempty"`
	PrimaryTag         *Tag       `json:"primary_tag,omitempty"`
	URL                *string    `json:"url,omitempty"`
	Excerpt            *string    `json:"excerpt,omitempty"`
}

type Tag struct {
	ID              *string   `json:"id"`
	Name            *string   `json:"name"`
	Slug            *string   `json:"slug"`
	Description     *string   `json:"description"`
	FeaturedImage   *string   `json:"featured_image"`
	Visibility      *string   `json:"visibility"`
	MetaTitle       *string   `json:"meta_title"`
	MetaDescription *string   `json:"meta_description"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Parent          *Tag      `json:"parent"`
}

type Error struct {
	ID      *string `json:"id"`
	Message *string `json:"message"`
	Context *string `json:"context"`
	Type    *string `json:"type"`
	/*
		TODO: determine types of these properties
		Details  *string `json:"details"`
		Property *string `json:"property"`
		Help     *string `json:"help"`
		Code     *string `json:"code"`
	*/
}

func (e Error) Error() string {
	return *e.Message + ": " + *e.Context
}
