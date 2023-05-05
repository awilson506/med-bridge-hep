package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/awilson506/med-bridge-hep/dao"
	"github.com/awilson506/med-bridge-hep/dao/queries"
	"github.com/gorilla/mux"
)

// GetSessions get all the sessions
func (c *Client) GetSessions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := c.db.Query(queries.GetSessions)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		sessions := []dao.Session{}
		for rows.Next() {
			session := dao.Session{}
			err := rows.Scan(&session.Id, &session.CreatedAt, &session.Notes)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			sessions = append(sessions, session)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(sessions)
	}
}

// GetSession get a single session
func (c *Client) GetSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		row := c.db.QueryRow(queries.GetSessionById, id)

		session := dao.Session{}
		err = row.Scan(&session.Id, &session.CreatedAt, &session.Notes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(session)
	}
}

// CreateSession create a session
func (c *Client) CreateSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session := dao.Session{}
		err := json.NewDecoder(r.Body).Decode(&session)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		row := c.db.QueryRow(queries.CreateSession, session.Notes)

		err = row.Scan(&session.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(session)
	}
}

// UpdateSession update a patient session details/notes
func (c *Client) UpdateSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		session := dao.Session{}
		err = json.NewDecoder(r.Body).Decode(&session)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = c.db.Exec(queries.UpdateSession, session.CreatedAt, session.Notes, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session.Id = id
		json.NewEncoder(w).Encode(session)
	}
}
