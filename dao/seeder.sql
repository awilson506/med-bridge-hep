-- seed therapist table
INSERT INTO therapists (id, first_name, last_name, email)
VALUES (1, 'Michael', 'Scott', 'mscott@dundermifflin.com'),
       (2, 'Dwight', 'Schrute', 'dschrute@dundermifflin.com');

-- seed patient table
INSERT INTO patients (id, first_name, last_name, date_of_birth, email, therapist_id)
VALUES (1, 'Alice', 'Smith', '1999-01-08', 'alicesmith@example.com', 1),
       (2, 'Bob', 'Johnson', '1989-01-08', 'bobjohnson@example.com', 1),
       (3, 'Charlie', 'Brown', '1969-01-08', 'charliebrown@example.com', 2);

-- seed program table
INSERT INTO programs (id, name, description, patient_id, therapist_id)
VALUES (1, 'Program 1', 'This is program 1', 1, 1),
       (2, 'Program 2', 'This is program 2', 2, 2);

-- seed session table
INSERT INTO sessions (program_id, patient_id, therapist_id)
VALUES (1, 1, 1),
       (1, 2, 1),
       (2, 3, 2);

-- seed exercise table
INSERT INTO exercises (id, name, description)
VALUES (1, 'Exercise 1', 'This is exercise 1 for program 1'),
       (2, 'Exercise 2', 'This is exercise 2 for program 1'),
       (3, 'Exercise 1', 'This is exercise 1 for program 2');

-- seed program_exercises table
INSERT INTO program_exercises (program_id, exercise_id)
VALUES (1, 1),
       (1, 2),
       (2, 1),
       (2, 2),
       (2, 3);
       
-- seed patient_results table
INSERT INTO patient_results (id, patient_id, exercise_id, program_id, score, complete)
VALUES (1, 1, 1, 1, 95, true),
	   (2, 1, 2, 1, 80, true);
