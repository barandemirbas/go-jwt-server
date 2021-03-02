package models

type User struct {
	_id      string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email 	 string `json:"email, omitempty"`
	Password string `json:"password,omitempty"`
}
