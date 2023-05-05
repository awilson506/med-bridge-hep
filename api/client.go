package api

import (
	"database/sql"
)

// Client struct to hold db and other clients
type Client struct {
	db *sql.DB
}

// Get an instance of the api client
func NewClient(db *sql.DB) *Client {
	return &Client{
		db: db,
	}
}
