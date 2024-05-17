// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/hawa130/computility-cloud/ent"
)

type LoginInput struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginPayload struct {
	Token *string   `json:"token,omitempty"`
	User  *ent.User `json:"user,omitempty"`
}
