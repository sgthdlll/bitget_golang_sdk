package test

import (
	"fmt"
	"github.com/sgthdlll/bitget_golang_sdk/internal_pkg/model"
	"github.com/sgthdlll/bitget_golang_sdk/pkg/client/ws"
	"testing"
)

func TestBitgetWsClient_New(t *testing.T) {

	client := new(ws.BitgetWsClient).Init(false, func(message string) {
		fmt.Println("default error:" + message)
	}, func(message string) {
		fmt.Println("default error:" + message)
	})

	var channels []model.SubscribeReq
	subReq1 := model.SubscribeReq{
		InstType: "USDT-FUTURES",
		Channel:  "books5",
		InstId:   "BTCUSDT",
	}
	channels = append(channels, subReq1)
	client.Subscribe(channels, func(message string) {
		fmt.Println("appoint:" + message)
	})

	select {}
}
