package models

import "cesargdd/rest-api/userspb"

type AuthResponse struct {
	Token string        `json:"token"`
	User  *userspb.User `json:"user"`
}
