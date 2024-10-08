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

Test_DB =  'postgresql://track:jFXmrOm5VE0eSl5xFpo9BKJRcN1zZkj9@dpg-crhk8g3v2p9s73bc36c0-a.oregon-postgres.render.com/track_vk5b'

How to run app:
* go run cmd/main.go

Postman Collections: 
* 
