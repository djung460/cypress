package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	models "github.com/djung460/cypress/models"
)

func Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome!"))
	})
}

func UserNuggetIndex(db models.DB) http.Handler {
	type (
		ret struct {
			Nuggets []string `json:"nuggets"`
		}
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("author_id"))

		if err != nil {
			log.Print("ERROR PANIC")
			log.Print(err)
			jsonErr(w, http.StatusUnprocessableEntity, err)
			return
		}

		nuggets, err := db.GetNuggetsByUser(id)

		if err != nil {
			log.Print("ERROR PANIC")
			log.Print(err)
			jsonErr(w, http.StatusInternalServerError, err)
			return
		}

		var jsonnuggets []string
		// convert the articles slice into strings
		for _, d := range nuggets {
			jsonnugget, err := d.MarshalJSON()
			if err != nil {
				log.Print("ERROR PANIC")
				log.Print(err)
				jsonErr(w, http.StatusInternalServerError, err)
				return
			}
			jsonnuggets = append(jsonnuggets, string(jsonnugget))
		}
		if err := json.NewEncoder(w).Encode(ret{Nuggets: jsonnuggets}); err != nil {
			log.Print(err)
			jsonErr(w, http.StatusInternalServerError, err)
		}
	})
}

func NuggetCreate(db models.DB) http.Handler {
	type (
		ret struct {
			Result bool `json:"results"`
		}
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var nugget models.Nugget

		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

		if err != nil {
			log.Print("ERROR PANIC")
			log.Print(err)
			jsonErr(w, http.StatusUnprocessableEntity, err)
			return
		}

		if err = r.Body.Close(); err != nil {
			log.Print("ERROR PANIC")
			log.Print(err)
			jsonErr(w, http.StatusUnprocessableEntity, err)
			return
		}

		if err = json.Unmarshal(body, &nugget); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusUnprocessableEntity)
			if err = json.NewEncoder(w).Encode(err); err != nil {
				log.Print("ERROR PANIC")
				log.Print(err)
				jsonErr(w, http.StatusUnprocessableEntity, err)
				return
			}
		}
		res, err := db.CreateNugget(nugget)

		if err != nil {
			log.Print("ERROR PANIC")
			log.Print(err)
			jsonErr(w, http.StatusUnprocessableEntity, err)
			return
		}

		if err := json.NewEncoder(w).Encode(ret{Result: res}); err != nil {
			log.Print(err)
			jsonErr(w, http.StatusInternalServerError, err)
		}

	})
}

func CategoryCreate(db models.DB) http.Handler {
	type (
		ret struct {
			Result bool `json:"results"`
		}
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var category models.Category

		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

		if err != nil {
			log.Print("ERROR PANIC")
			log.Print(err)
			jsonErr(w, http.StatusUnprocessableEntity, err)
			return
		}

		if err = r.Body.Close(); err != nil {
			log.Print("ERROR PANIC")
			log.Print(err)
			jsonErr(w, http.StatusUnprocessableEntity, err)
			return
		}

		if err = json.Unmarshal(body, &category); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusUnprocessableEntity)
			if err = json.NewEncoder(w).Encode(err); err != nil {
				log.Print("ERROR PANIC")
				log.Print(err)
				jsonErr(w, http.StatusUnprocessableEntity, err)
				return
			}
		}
		res, err := db.CreateCategory(category)

		if err != nil {
			log.Print("ERROR PANIC")
			log.Print(err)
			jsonErr(w, http.StatusUnprocessableEntity, err)
			return
		}

		if err := json.NewEncoder(w).Encode(ret{Result: res}); err != nil {
			log.Print(err)
			jsonErr(w, http.StatusInternalServerError, err)
		}
	})
}

func CategoryIndex(db models.DB) http.Handler {
	type (
		ret struct {
			Categories []string `json:"categories"`
		}
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		categories, err := db.GetCategories()

		if err != nil {
			log.Print("ERROR PANIC")
			log.Print(err)
			jsonErr(w, http.StatusInternalServerError, err)
			return
		}

		var jsoncategories []string
		// convert the articles slice into strings
		for _, d := range categories {
			jsoncategory, err := d.MarshalJSON()
			if err != nil {
				log.Print("ERROR PANIC")
				log.Print(err)
				jsonErr(w, http.StatusInternalServerError, err)
				return
			}
			jsoncategories = append(jsoncategories, string(jsoncategory))
		}
		if err := json.NewEncoder(w).Encode(ret{Categories: jsoncategories}); err != nil {
			log.Print(err)
			jsonErr(w, http.StatusInternalServerError, err)
		}
	})
}

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
