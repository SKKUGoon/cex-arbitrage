package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gopkg.in/yaml.v3"
	"kimchi/common"
	"kimchi/dao"
	"kimchi/ent"
	"log"
	"net/http"
	"os"
	"time"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type InternalServer struct {
	Conn          *gin.Engine
	LtDataBase    *ent.Client // Not yet operational
	CacheDataBase *redis.Client
}

type Webserver struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

// New
// Does all these things. Returns InternalServer structure
//  1. Create Gin engine with middleware
//  2. Add swagger doc
//  3. Connect to Redis database
//  4. TODO: connect to MySQL Database
//  5. TODO: connect to Telegram Service
func New() InternalServer {
	conn := gin.Default()
	conn.Use(common.CORSMiddleware())
	common.PrintGreenOk("Create gin engine. Attach middleware")
	conn.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	caching := dao.CacheNewConn("./Redis.yaml")
	common.PrintGreenOk("Create Redis connection for webserver")

	srv := InternalServer{
		Conn:          conn,
		CacheDataBase: caching,
	}
	return srv
}

func (client InternalServer) AddUtil() InternalServer {
	client.Conn.Use(common.CORSMiddleware())
	return client
}

func (client InternalServer) Serve(configFile string) *http.Server {
	c := client.strategyIEXA()

	webserverInfo := map[string]Webserver{}
	dat, err := os.ReadFile(configFile)
	if err != nil {
		log.Panicln("Webserver conn config file error:", err)
	}
	err = yaml.Unmarshal(dat, &webserverInfo)
	if err != nil {
		log.Panicln("Webserver conn config file parse error:", err)
	}
	// Internal webserver start information
	src := &http.Server{
		Addr:           webserverInfo["server"].Address + ":" + webserverInfo["server"].Port,
		Handler:        c.Conn,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return src
}

func (client InternalServer) strategyIEXA() InternalServer {
	client.Conn.POST("/band", func(context *gin.Context) {
		handleBandP(context, client.CacheDataBase)
	})
	client.Conn.POST("/premium", func(context *gin.Context) {
		handlePremiumP(context, client.CacheDataBase)
	})
	return client
}
