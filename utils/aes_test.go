package utils

import (
	"encoding/base64"
	"net/url"
	"reflect"
	"testing"
)

func TestAesDecryption(t *testing.T) {
	origData := "b1Mrk7C2GVw8DFvWXQ07LJ2Iw%2FQ0FMoZnPLI01xneUhwGTkpeJlf9byx%2F77CTn4VN6RoEwkB2w7kkRNBeOHi0yzgs%2BWWSM3u1hwJwBRiLtrQqd61hBu3xQJIcDGDKHa3NNPNr5VJ67eMD2l4V7oNh3wtiRRDwmWPk6vwPqr8HMXP5fMXYrZ7l5Jk1URpr5h6%2FSpMXYgzDYKaeqq2G%2F5w24anuNB7mRUin0VhUF6oNkWKL2m1P70KABQpM7ISFBH2QbSPsy5uKLs5XdSURrF%2B5KM3qfO6qLL00sgMTJsAsuc8RrTOiMFJXVBg2b%2FTkG1iJuAxul%2FEOUJqGPz85OTXiZ7AljrvO87L0XA5q5aiSMw1OiPMSNUK0DdZsSKv280BytquLHfnR5ZaJb9wj4jGNjQgM6eKXbwU%2FUTNeVcP7G5bebPBnAt7G%2FXfjfxcodQztzQ8MxydGbWjjEISqOoxGQ%3D%3D" // 待加密的数据

	clientSecret := "f0051f017c2cde1507dfc809f6772184"
	unescape, _ := url.QueryUnescape(origData)
	origByte, _ := base64.StdEncoding.DecodeString(unescape)
	//(origData)

	type args struct {
		key       []byte
		iv        []byte
		plainText []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{"AesEncryption", args{key: []byte(clientSecret), iv: []byte("0102030405060708"), plainText: origByte},
			[]byte(`{"appId": 80619, "buyerId": 12255578320, "buyerPhone": "15600300151", "env": "PROD", "kdtId": 97848743, "openId": "TtPFscDu869539917284249600", "orderNo": "23202109082225040106117", "payTime": 1631111104000, "price": 0, "shopRole": "SINGLE_SHOP", "shopType": "WSC", "skuIntervalText": "31", "skuVersionText": "测试版套餐", "status": 20, "type": "APP_SUBSCRIBE"}`),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AesDecryption(tt.args.key, tt.args.iv, tt.args.plainText)
			if (err != nil) != tt.wantErr {
				t.Errorf("DesEncryption() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DesEncryption() = \n%v, \n want \n%v \n%v \n%v ", got, tt.want, string(got), string(tt.want))
			}
		})
	}
}
