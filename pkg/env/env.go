package env

import (
	"WarpGPT/pkg/logger"
	"flag"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type ENV struct {
	Proxy           string
	Port            int
	Host            string
	Verify          bool
	AuthKey         string
	ArkoseMust      bool
	OpenaiHost      string
	OpenaiApiHost   string
	ProxyPoolUrl    string
	UserAgent       string
	RedisAddress    string
	RedisPasswd     string
	RedisDB         int
	PostgreSQLDBURI string
	CapSolver       string
	CapClientID     string
}

var Env ENV
var EnvFile string

func init() {
	flag.StringVar(&EnvFile, "e", ".env", "The env file path")
	flag.Parse()
	err := godotenv.Load(EnvFile)
	if err != nil {
		logger.Log.Error("初始化加载文件报错", err)
		return
	}
	port, err := strconv.Atoi(os.Getenv("port"))
	if err != nil {
		port = 5000
	}
	verify, err := strconv.ParseBool(os.Getenv("verify"))
	if err != nil {
		verify = false
	}
	arkoseMust, err := strconv.ParseBool(os.Getenv("verify"))
	if err != nil {
		arkoseMust = false
	}
	OpenaiHost := os.Getenv("openai_host")
	if OpenaiHost == "" {
		OpenaiHost = "chat.openai.com"
	}
	openaiApiHost := os.Getenv("openai_api_host")
	if openaiApiHost == "" {
		openaiApiHost = "api.openai.com"
	}
	loglevel := os.Getenv("log_level")
	if loglevel == "" {
		loglevel = "info"
	}
	proxyPoolUrl := os.Getenv("proxy_pool_url")
	redisAddress := os.Getenv("redis_address")
	if proxyPoolUrl != "" && redisAddress == "" {
		panic("配置proxyPoolUrl后未配置redis_address")
	}
	redisDb, err := strconv.Atoi(os.Getenv("redis_db"))
	if err != nil && proxyPoolUrl != "" {
		panic("DB填写出现问题")
	}
	Env = ENV{
		Proxy:           os.Getenv("proxy"),
		Port:            port,
		Host:            os.Getenv("host"),
		Verify:          verify,
		AuthKey:         os.Getenv("auth_key"),
		ArkoseMust:      arkoseMust,
		OpenaiHost:      OpenaiHost,
		OpenaiApiHost:   openaiApiHost,
		ProxyPoolUrl:    proxyPoolUrl,
		UserAgent:       "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Safari/605.1.15",
		RedisAddress:    redisAddress,
		RedisPasswd:     os.Getenv("redis_passwd"),
		RedisDB:         redisDb,
		PostgreSQLDBURI: os.Getenv("postgreSQL_db_URI"),
		CapSolver:       os.Getenv("cap_solver"),
		CapClientID:     os.Getenv("cap_client_id"),
	}
	logger.Log.Info("环境信息为：", Env)
}
