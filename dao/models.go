package dao

import "time"

type Patient struct {
	Id          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	Email       string `json:"email,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Address     string `json:"address,omitempty"`
}

type Therapist struct {
	Id          int     `json:"id"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Email       string  `json:"email"`
	Phone       *string `json:"phone,omitempty"`
	Credentials *string `json:"credentials,omitempty"`
}

type Exercise struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type Session struct {
	Id          int     `json:"id"`
	PatientID   int     `json:"patient_id"`
	TherapistID int     `json:"therapist_id"`
	ProgramID   int     `json:"program_id"`
	CreatedAt   string  `json:"date"`
	Notes       *string `json:"notes"`
}

type Program struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	PatientID   int         `json:"patient_id"`
	TherapistId int         `json:"-"`
	Therapist   Therapist   `json:"therapist"`
	Sessions    []*Session  `json:"sessions,omitempty"`
	Exercises   []*Exercise `json:"exercises"`
	CreatedAt   time.Time   `json:"created_at,omitempty"`
	UpdatedAt   time.Time   `json:"-"`
}

type PatientResult struct {
	Id int `json:"id"`
}
