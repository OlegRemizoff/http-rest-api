DROP TABLE users;

-- PGPASSWORD=postgres migrate -path migrations -database "postgres://postgres:postgres@localhost/restapi_dev?sslmode=disable" up
