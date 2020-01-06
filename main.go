package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("The ")
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		// handle the error here...
		e := errResponse{Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(&e)
		w.Write(res)
	}

	w.WriteHeader(http.StatusOK)
	id := r.Header.Get("id")
	if id == "" {
		id = "1"
	}

	var posts []Post
	db.Model("posts").Find(&posts)

	w.WriteHeader(http.StatusOK)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("The ")
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		// handle the error here...
		e := errResponse{Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(&e)
		w.Write(res)
	}

	w.WriteHeader(http.StatusOK)
	id := r.Header.Get("id")
	if id == "" {
		id = "1"
	}

	var post Post
	db.Model("posts").Where("id = ?", toInt(id)).First(&post)

	w.WriteHeader(http.StatusOK)
}
