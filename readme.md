# Golang Health Exercise Program (HEP) app
A small backend service to support a HEP web app

## Getting the service setup & running
If you don't already have Go installed you can download it and install it from here [Download Golang](https://go.dev/doc/install)

Note this application requires the latest Go version `1.20` but supports back to `1.17` if you already have a previous version
installed.  You will have to update the Go module to require a later version if you choose not to update.  This can be done by 
running: 
```sh
go mod edit -go=1.MY_VERSION
```
## Start the db
Start the Postgres database service with docker compose:
```sh
docker compose up
```
Run the `schema.sql` and `seeder.sql` files located [here](dao/)

## Start the server
In the root directory run:
```sh
go run main.go
```

## Database Design
Database [structure](docs/database/med_bridge_db.png)

## App Design
In deapth database and app design can be found [here](https://miro.com/app/board/uXjVML3AcEE=/?share_link_id=886445894903)

## Using the api
This application offers services for patients, thereapists and exercise programs.

Here's an example for retreiving a patient program:
```
curl http://localhost:8080/v1/patient-program/1
```

Sample response:
```json
{
    "id": 1,
    "name": "Program 1",
    "description": "This is program 1",
    "patient_id": 1,
    "therapist": {
        "id": 1,
        "first_name": "Michael",
        "last_name": "Scott",
        "email": "mscott@dundermifflin.com"
    },
    "exercises": [
        {
            "id": 1,
            "url": null,
            "name": "Exercise 1",
            "description": "This is exercise 1 for program 1"
        },
        {
            "id": 2,
            "url": null,
            "name": "Exercise 2",
            "description": "This is exercise 2 for program 1"
        }
    ],
    "created_at": "2023-05-07T21:50:51.057318Z"
}
```

## Todos:
- Add authentication
- User roles for scoping api requests
- Add services for patients, results and therapists
- Exercise program versioning needs to be implemented for changes after n sessions
- ORM could be updated
- Implement tests for each handler
- Because this program may have a daily or even hourly routine we'll need to implement scheduling and scoring 
based on multiple repitions of the exercise program
