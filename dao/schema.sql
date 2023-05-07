-- create therapists table
CREATE TABLE therapists (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    email VARCHAR(100) NOT NULL,
    phone VARCHAR(20),
    credentials VARCHAR(200),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- create patients table
CREATE TABLE patients (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    date_of_birth DATE NOT NULL,
    email VARCHAR(100) NOT NULL,
    phone VARCHAR(20),
    address VARCHAR(200),
    therapist_id INTEGER REFERENCES therapists(id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- create exercises table
CREATE TABLE exercises (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    url TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- create programs table
CREATE TABLE programs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    patient_id INTEGER REFERENCES patients(id),
    therapist_id INTEGER REFERENCES therapists(id),
    description TEXT,
    version INTEGER,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- create program exercises table
CREATE TABLE program_exercises (
    program_id INTEGER REFERENCES programs(id),
    exercise_id INTEGER REFERENCES exercises(id),
    PRIMARY KEY (program_id, exercise_id)
);

-- create sessions table
CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    patient_id INTEGER REFERENCES patients(id),
    therapist_id INTEGER REFERENCES therapists(id),
    program_id INTEGER REFERENCES programs(id),
    notes TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- create results table
CREATE TABLE patient_results (
    id SERIAL PRIMARY KEY,
    patient_id INTEGER REFERENCES patients(id),
    exercise_id INTEGER REFERENCES exercises(id),
    program_id INTEGER REFERENCES programs(id),
  	score INTEGER,
  	complete BOOLEAN,
  	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- create updated at trigger function
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- add updated at function to tables
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON therapists
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON patients
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON exercises
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON programs
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON sessions
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON patient_results
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
