package main

import (
	pb "github.com/Namdar1Ibrakhim/student-track-system/proto"
	"google.golang.org/grpc"
	"os"

	track "github.com/Namdar1Ibrakhim/student-track-system"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/handler"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/repository"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/service"
	"github.com/gin-gonic/gin"
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

	db, err := repository.NewPostgresDB(repository.Config{ //создаем обьект NewPostgresDB реопзитория и инициализируем туда Конфигурации
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

	mlAddress := viper.GetString("url")
	conn, err := grpc.Dial(mlAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logrus.Fatalf("Error connecting to ML: %s", err.Error())
	}
	defer conn.Close()

	mlClient := pb.NewPredictionServiceClient(conn)

	repos := repository.NewRepository(db)
	services := service.NewService(repos, mlClient)
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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
