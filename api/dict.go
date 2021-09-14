package api

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"time"
	"unicode/utf8"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

var (
	DictApi = fmt.Sprintf("%s/api", Api)
)

// Translate query string to different natural language
//
// document link: https://ai.youdao.com/DOCSIRMA/html/%E8%87%AA%E7%84%B6%E8%AF%AD%E8%A8%80%E7%BF%BB%E8%AF%91/API%E6%96%87%E6%A1%A3/%E6%96%87%E6%9C%AC%E7%BF%BB%E8%AF%91%E6%9C%8D%E5%8A%A1/%E6%96%87%E6%9C%AC%E7%BF%BB%E8%AF%91%E6%9C%8D%E5%8A%A1-API%E6%96%87%E6%A1%A3.html
func Translate(query string, opts ...Option) ([]byte, error) {
	timestamp := fmt.Sprintf("%v", time.Now().Unix())
	body := map[string]string{
		QueryField:       query,
		FromLangField:    LangTypeAuto,
		ToLangField:      LangTypeCN,
		AppKeyField:      AppKey,
		SaltField:        uuid.New().String(),
		SignTypeField:    SignTypeV3,
		CurrentTimeField: timestamp,
	}
	for _, opt := range opts {
		if err := opt(body); err != nil {
			return nil, err
		}
	}
	signSha256(body)

	resp, err := resty.New().R().SetHeader(headerContentType, contentTypeFormUrlEncode).SetFormData(body).Post(DictApi)
	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}

func signSha256(body map[string]string) {
	input := body[QueryField]
	inputLen := utf8.RuneCountInString(input)
	if inputLen > 20 {
		runes := []rune(input)
		input = fmt.Sprintf("%s%v%s", string(runes[:10]), inputLen, string(runes[inputLen-10:]))
	}
	h := sha256.New()
	io.WriteString(h, AppKey)
	io.WriteString(h, input)
	io.WriteString(h, body[SaltField])
	io.WriteString(h, body[CurrentTimeField])
	io.WriteString(h, AppSecret)
	b := h.Sum(nil)
	body[SignField] = hex.EncodeToString(b)
}
