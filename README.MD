# Mezink Assignment

Requirement for running the project : 
- docker desktop - latest version (>=4.24.0)

library used :
- https://github.com/go-chi/chi (golang simple router)
- https://github.com/lib/pq (golang postgres driver)

## Running the project
1. clone the project 
`git clone https://github.com/Skyezon/mezink-assignment.git`
2. `cd ./mezink-assignment`
3. make sure docker desktop is latest version and is running then
`docker-compose up`
4. project will run on `localhost:8080`

## Importing Postman collection & testing
1. postman collection can be found on : `postmancollection.json`
2. open postman import -> file -> find `postmancollection.json` or just drag an drop into it
3. `mezink-assignment` collection should show up and for testing the main task will be `get sum (main task)`

## Notes
 if you want to change database name/user/password you can modify it on environment in `docker-compose.yaml` both for db & service
 
 Database port is not exposed if you wanted to connect to database you will need to expose it first on `docker-compose.yaml` by adding `:[desired-port-number]` in db -> ports. e.g : `5432:5432`
 
 There are other endpoints to accommodate testing, other endpoint such as (check postman collection for more info) :
 - create record - to add more record into database
 - get all record - getting all the record without having to check database
 - get record - get single record by id
 
