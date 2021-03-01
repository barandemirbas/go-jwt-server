package models

type User struct {
	_id      string `json:"id,omitempty"`
	Name     string `json:"name,omitempty" validate:"gte=3 & lte=25 & format=alnum"`
	Password string `json:"password,omitempty" validate:"gte=8 & lte=256"`
}
