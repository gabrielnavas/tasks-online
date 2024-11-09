package main

import (
	"api/controllers"
	"api/database"
	"api/repositories"
	"api/services"
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"

	"github.com/redis/go-redis/v9"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PATCH, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

var ctx = context.Background()

func main() {

	// envs
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var (
		DB_MYSQL_USERNAME      = os.Getenv("DB_MYSQL_USERNAME")
		DB_MYSQL_PASSWORD      = os.Getenv("DB_MYSQL_PASSWORD")
		DB_MYSQL_HOST          = os.Getenv("DB_MYSQL_HOST")
		DB_MYSQL_PORT          = os.Getenv("DB_MYSQL_PORT")
		DB_MYSQL_DATABASE_NAME = os.Getenv("DB_MYSQL_DATABASE_NAME")

		DB_REDIS_ADDR     = os.Getenv("DB_REDIS_ADDR")
		DB_REDIS_PASSWORD = os.Getenv("DB_REDIS_PASSWORD")
	)

	// cache
	rdb := redis.NewClient(&redis.Options{
		Addr:     DB_REDIS_ADDR,     //"localhost:6379",
		Password: DB_REDIS_PASSWORD, //"secretpass123", // set password
		DB:       0,                 // use default DB
	})
	// reseta o cache
	rdb.FlushAll(ctx)

	// databases
	dbMysql := database.OpenMysqlConnection(
		DB_MYSQL_USERNAME,
		DB_MYSQL_PASSWORD,
		DB_MYSQL_HOST,
		DB_MYSQL_PORT,
		DB_MYSQL_DATABASE_NAME,
	)

	// controllers, services, repositories
	tr := repositories.NewTaskRepository(dbMysql)
	ts := services.NewTaskService(rdb, tr)
	tc := controllers.NewTaskController(tr, ts)

	// gin
	router := gin.Default()

	// cors
	router.Use(CORSMiddleware())

	// routes
	tasksRouter := router.Group("/tasks")
	{
		tasksRouter.POST("", tc.CreateTask)
		tasksRouter.GET("", tc.FindTasks)
		tasksRouter.GET(":taskId", tc.FindTaskById)
		tasksRouter.PATCH(":taskId", tc.UpdateTask)
		tasksRouter.DELETE(":taskId", tc.DeleteTask)
	}

	router.Run(":8080")
}
