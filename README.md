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
### Importing the Postman Collection
1. Download the [Postman Collection](./postman_collection.json) file.
2. Open Postman and go to **File > Import**.
3. Select the `postman_collection.json` file to import the collection.
OR
* google-drive: https://drive.google.com/file/d/1SJhDRBaYGG-BLm-KWjY9iXm0zu9fYZP1/view?usp=sharing

How to run with docker:
* docker-compose up (-d)

Project host: https://backend-track-1dxh.onrender.com
