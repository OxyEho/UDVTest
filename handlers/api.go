package handlers

import (
	"UDVTest/db"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

const (
	PageNum  = 1
	PageSize = 10
)

type Api struct {
	repo db.IRepo
}

func (a *Api) Run() {
	router := mux.NewRouter()
	router.Path("/api/book/{id:[0-9]+}").HandlerFunc(a.bookHandler).Methods("GET")
	router.Path("/api/book/{id:[0-9]+}/items").HandlerFunc(a.bookItemsHandler).Methods("GET")
	router.
		Path("/api/book").
		HandlerFunc(a.booksHandler).
		Queries("page_number", "{page_number:[0-9]+}", "size", "{size:[0-9]+}").
		Methods("GET")
	router.Path("/api/book").HandlerFunc(a.booksHandler).Methods("GET")
	http.Handle("/", router)
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatal(err)
	}
}

func (a *Api) bookHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	book, err := a.repo.GetBook(uint(id))
	if err != nil {
		sendErr(w, http.StatusNotFound, "Not found")
		return
	}
	sendJson(w, book)
}

func (a *Api) bookItemsHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	bookItems, err := a.repo.GetBookItems(uint(id))
	if err != nil {
		sendErr(w, http.StatusNotFound, "Not found")
		return
	}
	sendJson(w, bookItems)
}

func (a *Api) booksHandler(w http.ResponseWriter, r *http.Request) {
	pageNum, errPage := getQueryInt(r.URL.Query().Get("page_number"), PageNum)
	size, errSize := getQueryInt(r.URL.Query().Get("size"), PageSize)
	if errPage != nil || errSize != nil {
		sendErr(w, 400, "Bad request")
		return
	}
	var limit, offset int
	limit = size
	offset = (pageNum - 1) * size
	books, err := a.repo.GetBooks(limit, offset)
	if err != nil || len(books) == 0 {
		sendErr(w, http.StatusNotFound, "Not found")
		return
	}
	sendJson(w, books)
}

func NewApi(repo db.IRepo) Api {
	return Api{repo: repo}
}

func sendJson(w http.ResponseWriter, v interface{}) {
	jsonResp, jsonErr := json.Marshal(v)
	if jsonErr != nil {
		sendErr(w, 500, "Bad request")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func sendErr(w http.ResponseWriter, code int, msg string) {
	errMsg := make(map[string]string)
	errMsg["msg"] = msg
	jsonResp, _ := json.Marshal(errMsg)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonResp)
}

func getQueryInt(s string, n int) (int, error) {
	if s != "" {
		var err error
		n, err = strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		if n <= 0 {
			return 0, errors.New("bad num")
		}
	}
	return n, nil
}
