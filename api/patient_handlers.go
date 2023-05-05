package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/awilson506/med-bridge-hep/dao"
	"github.com/awilson506/med-bridge-hep/dao/queries"
	"github.com/gorilla/mux"
)

func (c *Client) GetPatientExerciseProgram() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: this is a bit clunky, should use something like gorm to clean up the queries
		id, err := strconv.Atoi(mux.Vars(r)["programId"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		row := c.db.QueryRow(queries.GetProgramById, id)
		program := dao.Program{}

		err = row.Scan(&program.Id, &program.Name, &program.Description, &program.PatientID, &program.TherapistId, &program.CreatedAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		therapist := c.db.QueryRow(queries.GetTherapistById, &program.TherapistId)
		err = therapist.Scan(&program.Therapist.Id, &program.Therapist.FirstName, &program.Therapist.LastName, &program.Therapist.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		exercises, err := c.db.Query(queries.GetExercisesByProgramId, &program.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//TODO add session data to the kitchen sink here
		for exercises.Next() {
			exercise := dao.Exercise{}
			err := exercises.Scan(&exercise.Id, &exercise.Name, &exercise.Description)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			program.Exercises = append(program.Exercises, &exercise)
		}
		json.NewEncoder(w).Encode(program)
	}
}
