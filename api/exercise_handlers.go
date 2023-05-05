package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/awilson506/med-bridge-hep/dao"
	"github.com/awilson506/med-bridge-hep/dao/queries"
	"github.com/gorilla/mux"
)

// GetExercises get all the exercises
func (c *Client) GetExercises() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := c.db.Query(queries.GetExercises)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		exercises := []dao.Exercise{}
		for rows.Next() {
			exercise := dao.Exercise{}
			err := rows.Scan(&exercise.Id, &exercise.Name, &exercise.Description)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			exercises = append(exercises, exercise)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(exercises)
	}
}

// GetExercise get a single exercise
func (c *Client) GetExercise() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		row := c.db.QueryRow(queries.GetExerciseById, id)

		exercise := dao.Exercise{}
		err = row.Scan(&exercise.Id, &exercise.Name, &exercise.Description)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(exercise)
	}
}

// CreateExercise create a exercise
func (c *Client) CreateExercise() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		exercise := dao.Exercise{}
		err := json.NewDecoder(r.Body).Decode(&exercise)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		row := c.db.QueryRow(queries.CreateExercise, exercise.Name, exercise.Description)

		err = row.Scan(&exercise.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(exercise)
	}
}

// UpdateExercise update a exercise record
func (c *Client) UpdateExercise() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		exercise := dao.Exercise{}
		err = json.NewDecoder(r.Body).Decode(&exercise)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = c.db.Exec(queries.UpdateExercise, exercise.Name, exercise.Description, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		exercise.Id = id
		json.NewEncoder(w).Encode(exercise)
	}
}
