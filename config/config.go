package config

import "github.com/sgthdlll/bitget_golang_sdk/constants"

var (
	BaseUrl = "https://api.bitget.com"
	WsUrl   = "wss://ws.bitget.com/v2/ws/public"

	ApiKey        = ""
	SecretKey     = ""
	PASSPHRASE    = ""
	TimeoutSecond = 30
	SignType      = constants.SHA256
)

func SetBitgetApikey(apiKey, apiSecret string) {
	ApiKey = apiKey
	SecretKey = apiSecret
}
