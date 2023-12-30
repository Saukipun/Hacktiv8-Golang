package main

import (
	commentDelivery "MYGRAM-API/comment/delivery/http"
	commentRepository "MYGRAM-API/comment/repository/postgres"
	commentUseCase "MYGRAM-API/comment/usecase"
	"MYGRAM-API/config/database"
	_ "MYGRAM-API/docs"
	photoDelivery "MYGRAM-API/photo/delivery/http"
	photoRepository "MYGRAM-API/photo/repository/postgres"
	photoUseCase "MYGRAM-API/photo/usecase"
	socialMediaDelivery "MYGRAM-API/socialmedia/delivery/http"
	socialMediaRepository "MYGRAM-API/socialmedia/repository/postgres"
	socialMediaUseCase "MYGRAM-API/socialmedia/usecase"
	userDelivery "MYGRAM-API/user/delivery/http"
	userRepository "MYGRAM-API/user/repository/postgres"
	userUseCase "MYGRAM-API/user/usecase"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MyGram API
// @version 1.0
// @description MyGram is a free photo sharing app written in Go. People can share, view, and comment photos by everyone. Anyone can create an account by registering an email address and creating a username.
// @termOfService http://swagger.io/terms/
// @contact.name sauki
// @contact.email saukiadillah@gmail.com
// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey  Bearer
// @in                          header
// @name                        Authorization
// @description					Description for what is this security definition being used
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	db := database.StartDB()

	routers := gin.Default()

	routers.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	})

	userRepository := userRepository.NewUserRepository(db)
	userUseCase := userUseCase.NewUserUseCase(userRepository)

	userDelivery.NewUserHandler(routers, userUseCase)

	photoRepository := photoRepository.NewPhotoRepository(db)
	photoUseCase := photoUseCase.NewPhotoUseCase(photoRepository)

	photoDelivery.NewPhotoHandler(routers, photoUseCase)

	commentRepository := commentRepository.NewCommentRepository(db)
	commentUseCase := commentUseCase.NewCommentUseCase(commentRepository)

	commentDelivery.NewCommentHandler(routers, commentUseCase, photoUseCase)

	socialMediaRepository := socialMediaRepository.NewSocialMediaRepository(db)
	socialMediaUseCase := socialMediaUseCase.NewSocialMediaUseCase(socialMediaRepository)

	socialMediaDelivery.NewSocialMediaHandler(routers, socialMediaUseCase)

	routers.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	port := os.Getenv("PORT")

	if len(os.Args) > 1 {
		reqPort := os.Args[1]

		if reqPort != "" {
			port = reqPort
		}
	}

	if port == "" {
		port = "8080"
	}

	routers.Run(":" + port)
}
