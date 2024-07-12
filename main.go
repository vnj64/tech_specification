package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
	"runtime"
	"strings"
	v1 "tech/api/v1"
	"tech/connection/postgresql_driver"
	"tech/domain"
	"tech/domain/services"
	"tech/services/config"
	encryptorService "tech/services/encryptor"
	"time"
)

const (
	exitSqlStatus int    = 50
	exitApp       int    = 1
	errorsSQL     string = "APP SQL ERROR"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9999
// @BasePath

var (
	DB  *gorm.DB
	CTX domain.Context
)

func GinRouter() *gin.Engine {
	router := gin.New()

	// Custom Logger
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s |%s %d %s| %s |%s %s %s %s | %s | Body: %s | %s | %s\n",
			param.TimeStamp.Format(time.RFC1123),
			param.StatusCodeColor(),
			param.StatusCode,
			param.ResetColor(),
			param.ClientIP,
			param.MethodColor(),
			param.Method,
			param.ResetColor(),
			param.Path,
			param.Latency,
			byteCountSI(param.BodySize),
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	// CORS
	corsConfig := cors.DefaultConfig()
	//corsConfig.AllowOrigins = []string{CTX.Services().Config().ServerFullUrl()}
	corsConfig.AllowOriginFunc = func(origin string) bool {
		if strings.Contains(origin, "localhost") {
			return true
		}

		if strings.Contains(origin, CTX.Services().Config().ServerUrl()) {
			return true
		}

		return false
	}
	corsConfig.AddAllowMethods("GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS")
	corsConfig.AllowCredentials = true
	router.Use(cors.New(corsConfig))

	router.Use(func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
		} else {
			c.Next()
		}
	})

	// Recovery для восстановления после НЕ 200 статус кодов.
	router.Use(gin.Recovery())

	// Sessions
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)
	store := cookie.NewStore(authKeyOne, encryptionKeyOne)

	secure := CTX.Services().Config().ServerProtocol() == "https"
	sameSite := http.SameSiteLaxMode

	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   14 * 24 * 60 * 60,
		Secure:   secure,
		HttpOnly: true,
		SameSite: sameSite,
	})

	router.Use(sessions.Sessions("CORE_SERVICE", store))

	return router
}

func InitDB(cfg services.Config) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		cfg.PostgresqlUser(),
		cfg.PostgresqlPassword(),
		cfg.PostgresqlHost(),
		cfg.PostgresqlPort(),
		cfg.PostgresqlDatabase(),
	)

	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("database connect failed")
		fmt.Println(err.Error())
		os.Exit(exitSqlStatus)
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("database open failed")
		fmt.Println(err.Error())
		os.Exit(exitSqlStatus)
	}

	err = sqlDB.Ping()
	if err != nil {
		fmt.Println("database ping failed")
		fmt.Println(err.Error())
		os.Exit(exitSqlStatus)
	}

	DB = db
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} json "{"message": "pong"}"
// @Router /ping [get]
func main() {
	// Maximize CPU usage for maximum performance
	runtime.GOMAXPROCS(runtime.NumCPU())

	encryptor := encryptorService.Make()

	cfg, err := config.Make(encryptor)
	if err != nil {
		fmt.Println(err)
		os.Exit(exitApp)
	}

	connection, err := postgresql_driver.Make(cfg.PostgresqlUser(), cfg.PostgresqlPassword(), cfg.PostgresqlHost(), cfg.PostgresqlPort(), cfg.PostgresqlDatabase())
	if err != nil {
		fmt.Println(err)
		os.Exit(exitApp)
	}

	domainCtx := &ctx{
		session: &session{},
		services: &svs{
			encryptor: encryptor,
			config:    cfg,
		},
		connection: connection,
	}

	CTX = domainCtx

	// Server Settings
	router := GinRouter()
	server := &http.Server{
		Addr:           ":" + CTX.Services().Config().ServerPort(),
		Handler:        router,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   5 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}

	// Инициализирует подключение к базе данных
	InitDB(cfg)

	router.Use(func(c *gin.Context) { // Добавляем везде контекст
		c.Set("context", domainCtx.MakeWithSession(&session{
			isAuth: sessions.Default(c).Get("sessionIsAuth") == true,
		}))
		c.Next()
	})

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")

	swagger := router.Group("/docs")
	swagger.Use()
	{
		swagger.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

	user := router.Group("/user")
	user.Use()
	{
		user.POST("", v1.Wrap(v1.CreateUserHandler))
		user.GET("/:userId", v1.Wrap(v1.GetUserHandler))
		user.GET("/users", v1.Wrap(v1.GetAllUsersHandler))
		user.PATCH("/:userId", v1.Wrap(v1.UpdateUserHandler))
		user.DELETE("/:userId", v1.Wrap(v1.DeleteUserHandler))
	}

	role := router.Group("/role")
	role.Use()
	{
		role.POST("", v1.Wrap(v1.CreateRoleHandler))
		role.GET("/:roleId", v1.Wrap(v1.GetRoleHandler))
		role.GET("/roles", v1.Wrap(v1.GetAllRoleHandler))
		role.PATCH("/:roleId", v1.Wrap(v1.UpdateRoleHandler))
		role.DELETE("/:roleId", v1.Wrap(v1.DeleteRoleHandler))
	}

	// Запуск сервера
	fmt.Printf("Starting HTTP Server on port :%s\n", CTX.Services().Config().ServerPort())
	err = server.ListenAndServe()
	if err != nil {
		fmt.Printf("Starting HTTP Server error: %s", err.Error())
		os.Exit(exitApp)
	}
}
