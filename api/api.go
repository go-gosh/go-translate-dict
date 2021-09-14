package api

import (
	"log"
	"os"
)

const (
	EnvAppKey    = "GO_YOUDAO_DICT_KEY"
	EnvAppSecret = "GO_YOUDAO_DICT_SECRET"

	SignTypeV3 = "v3"

	headerContentType        = "Content-Type"
	contentTypeFormUrlEncode = "application/x-www-form-urlencoded"
)

var (
	Api = "https://openapi.youdao.com"

	AppKey    string
	AppSecret string
)

func init() {
	AppKey, _ = os.LookupEnv(EnvAppKey)
	if AppKey == "" {
		log.Fatalf("please set envoriment variable [%s]", EnvAppKey)
	}
	AppSecret, _ = os.LookupEnv(EnvAppSecret)
	if AppSecret == "" {
		log.Fatalf("please set envoriment variable [%s]", EnvAppSecret)
	}
}
