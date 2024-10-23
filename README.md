# STUDENT-TRACK
Back-End

Installed libraries:
* go get -u github.com/gin-gonic
* go get -u gorm.io/gorm
* go get -u github.com/spf13/viper
* go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

Make migrations: 
* For up: migrate -path ./schema -database ${YOUR_DB} up
* For down: migrate -path ./schema -database ${YOUR_DB} down

Test_DB =  'postgresql://postgres1:8OO5fENAktaKxaGWEZ7RuFB2GDCaQ3dC@dpg-cs5pdpt6l47c73f6q7k0-a.oregon-postgres.render.com/track_p0r8'

How to run app:
* go run cmd/main.go

Postman Collections: 

