package models

type Man struct {
	Id int64 `json:"id,omitempty"`

	Age int8 `json:"age,omitempty"`

	Name string `json:"name,omitempty"`

	Verified bool `json:"verified,omitempty"`
}
