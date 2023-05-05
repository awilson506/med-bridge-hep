package queries

// prgram queries
const GetProgramById = `
	SELECT 
		id, 
		name, 
		description, 
		patient_id, 
		therapist_id, 
		created_at 
	FROM programs WHERE id = $1
`

const GetExercisesByProgramId = `
	SELECT
		e.id,
		e.name, 
		e.description
	FROM programs pr 
	JOIN program_exercises pe ON pr.id = pe.program_id
	JOIN exercises e ON e.id = pe.exercise_id
	WHERE pr.id = $1
`

const GetTherapistById = `
	SELECT id, 
		first_name, 
		last_name, 
		email 
	FROM therapists 
	WHERE id = $1
`

// exercise queries
const GetExercises = `
	SELECT 
		id, 
		name, 
		description 
	FROM exercises
`

const GetExerciseById = `
	SELECT 
		id, 
		name, 
		description 
	FROM exercises WHERE id = $1
`

const CreateExercise = `
	INSERT INTO exercises(name, description) VALUES($1, $2) RETURNING id
`
const UpdateExercise = `
	UPDATE exercises SET name = $1, description = $2 WHERE id = $3
`

// session queries
const GetSessions = `
	SELECT 
		id, 
		date, 
		notes 
	FROM sessions
`
const GetSessionById = `
	SELECT 
		id, 
		date, 
		notes 
	FROM sessions 
	WHERE id = $1
`
const CreateSession = `
	INSERT INTO sessions(notes) VALUES($1) RETURNING id
`
const UpdateSession = `
	UPDATE sessions SET date = $1, notes = $2 WHERE id = $4
`
