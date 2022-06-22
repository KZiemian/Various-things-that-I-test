package main

import "fmt"

type Entity interface {
	UnmarshalHTTP(*http.Request) error
}

func GetEntity(r *http.Request, v Entity) error {
	return v.UnmarshalHTTP(r)
}

func (u *User) UnmarshalHTTP(r *http.Request) error {
	//
}

func main() {
	var u User
	if err := GetEntity(req, &u); err != nil {
		//
	}
}
