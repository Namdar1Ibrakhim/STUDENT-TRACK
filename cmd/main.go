package main

import (
	"os"

	track "github.com/Namdar1Ibrakhim/student-track-system"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/handler"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/repository"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter)) //Логгирование

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error()) //Проверка на инициализацию Конфигураций в функции ниже
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error()) //Проверка на инициализацию .env файла
	}

	db, err := repository.NewPostgresDB(repository.Config{    //создаем обьект NewPostgresDB реопзитория и инициализируем туда Конфигурации
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),

		//из .env файлы
		Password: os.Getenv("DB_PASSWORD"),
		//
		
		DBName:  viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db:  %s", err.Error()) //Проверка на инициализацию этого конфига
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(track.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

}

func initConfig() error { //Функция добавления конфигурации
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
