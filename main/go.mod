module rest-api/main

go 1.24.1

require (
	github.com/lib/pq v1.10.9
	rest-api/config v0.0.0-00010101000000-000000000000
)

require github.com/joho/godotenv v1.5.1 // indirect

replace rest-api/config => ../config
