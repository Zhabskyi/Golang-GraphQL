package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	movie, err := app.models.DB.Get(id)

	err = app.writeJSON(w, http.StatusOK, movie, "movie")

	app.checkErr(w, err)
}

func (app *application) getAllMovie(w http.ResponseWriter, r *http.Request) {
	movies, err := app.models.DB.All()
	if app.checkErr(w, err) {
		return
	}

	err = app.writeJSON(w, http.StatusOK, movies, "movies")
	app.checkErr(w, err)

}

func (app *application) getAllGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := app.models.DB.GenresAll()
	if app.checkErr(w, err) {
		return
	}

	err = app.writeJSON(w, http.StatusOK, genres, "genres")
	app.checkErr(w, err)

}

func (app *application) getAllMoviesByGenge(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	genreID, err := strconv.Atoi(params.ByName("genre_id"))
	if app.checkErr(w, err) {
		return
	}

	movies, err := app.models.DB.All(genreID)
	if app.checkErr(w, err) {
		return
	}

	err = app.writeJSON(w, http.StatusOK, movies, "movies")
	app.checkErr(w, err)
}

func (app *application) deleteMovie(w http.ResponseWriter, r *http.Request) {

}

func (app *application) insertMovie(w http.ResponseWriter, r *http.Request) {

}

func (app *application) updateMovie(w http.ResponseWriter, r *http.Request) {

}

func (app *application) searchMovies(w http.ResponseWriter, r *http.Request) {

}
