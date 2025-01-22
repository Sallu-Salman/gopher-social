package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"sallu.com/internal/store"
)

func (a *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	post := &store.Post{}
	ReadJson(r, post)

	log.Printf("Received post: %v\n", post)
	s := &a.store.Posts

	if err := Validate.Struct(post); err != nil {
		WriteJsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := s.Create(post); err != nil {
		WriteJsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("Created post: %v\n", post)
	WriteJson(w, http.StatusCreated, post)
}

func (a *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "postId")
	id, err := strconv.Atoi(postId)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, err)
		return
	}

	s := a.store.Posts
	post, err := s.GetById(id)
	if err != nil {
		WriteJsonError(w, http.StatusInternalServerError, err.Error())
	}
	log.Printf("Got post: %v\n", post)

	WriteJson(w, http.StatusOK, post)
}
