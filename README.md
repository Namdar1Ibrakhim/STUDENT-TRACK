# STUDENT-TRACK
Back-End

Installed libraries:
* go get -u github.com/gin-gonic
* go get -u gorm.io/gorm
* go get -u github.com/spf13/viper
* go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

Make migrations: 
* For up:
*  * migrate -path ./schema -database ${YOUR_DB} up

How to run app:
* go run cmd/main.go

