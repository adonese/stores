package main

import (
	"github.com/jinzhu/gorm"
)

// Post respresent each post in the list
type Post struct {
	gorm.Model
	Name        string
	Price       float32
	Description string
	Pictures    []Picture
}

type Picture struct {
	Filename string
	Mimetype string
}

type Likes struct {
	Count int
}

type Hits struct {
	Count int
}

type errResponse struct {
	Message string
	Code    string
}

type response struct {
}
