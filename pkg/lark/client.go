package lark

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

const larkAPI = "https://open.larksuite.com/open-apis/bot/v2/hook/"

// Client lark client
type Client struct {
	AccessToken string
	Secret      string
}

// NewClient new lark client
func NewClient(accessToken, secret string) *Client {
	return &Client{
		AccessToken: accessToken,
		Secret:      secret,
	}
}

// Response response struct
type Response struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

// Send will send message
func (d *Client) Send(ctx context.Context, message Message) (*Response, error) {
	res := &Response{}

	if len(d.AccessToken) < 1 {
		return res, fmt.Errorf("accesstoken is empty")
	}

	timestamp := time.Now().Unix()
	sign, err := GenSign(d.Secret, timestamp)
	if err != nil {
		return res, err
	}

	body := message.Body()
	body["timestamp"] = strconv.FormatInt(timestamp, 10)
	body["sign"] = sign

	client := resty.New()
	URL := fmt.Sprintf("%v%v", larkAPI, d.AccessToken)
	resp, err := client.SetTransport(&http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}).SetRetryCount(3).R().SetContext(ctx).
		SetBody(body).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetResult(&Response{}).
		ForceContentType("application/json").
		Post(URL)

	if err != nil {
		return nil, err
	}

	result := resp.Result().(*Response)
	if result.Code != 0 {
		return res, fmt.Errorf("send message to lark error = %s", result.Msg)
	}
	return result, nil
}
