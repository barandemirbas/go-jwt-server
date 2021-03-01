package models

type Book struct {
	_id    string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
	Rating int    `json:"rating,omitempty"`
}
