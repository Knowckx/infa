package web

import (
	"fmt"
	"time"

	"log/slog"

	"github.com/cockroachdb/errors"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

type RestClient struct {
	*resty.Client
}

func NewRestClient() *RestClient {
	out := &RestClient{}
	out.Client = resty.New().SetTimeout(30*time.Second).SetHeader("Accept", "application/json")
	return out
}

func (cli *RestClient) DoGet(url string, queryParams map[string]string) (*gjson.Result, error) {
	slog.Debug("send request get", "url", url, "params", queryParams)
	if queryParams != nil {
		cli.Client = cli.SetQueryParams(queryParams)
	}
	resp, err := cli.R().Get(url)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	err = CheckRestyResponse(resp)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	cli.SetCookies(resp.Cookies())
	res := gjson.ParseBytes(resp.Body())
	return &res, nil
}

func (cli *RestClient) DoGetResp(url string, queryParams map[string]string) (*resty.Response, error) {
	slog.Debug("send request get", "url", url, "params", queryParams)
	if queryParams != nil {
		cli.Client = cli.SetQueryParams(queryParams)
	}
	resp, err := cli.R().Get(url)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	err = CheckRestyResponse(resp)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	cli.SetCookies(resp.Cookies())
	return resp, nil
}

func (cli *RestClient) DoPost(url string, body interface{}) (*gjson.Result, error) {
	slog.Debug("post", "url", url, "body", body)
	resp, err := cli.R().SetBody(body).Post(url)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	err = CheckRestyResponse(resp)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	cli.SetCookies(resp.Cookies())
	res := gjson.ParseBytes(resp.Body())
	return &res, nil
}

func CheckRestyResponse(resp *resty.Response) error {
	if resp.IsError() {
		slog.Error("Http Request Failed", "Status Code", resp.StatusCode())
		data := resp.Body()
		slog.Error("body", "body", string(data))
		return fmt.Errorf("Http Request Failed")
	}
	slog.Debug("Http Request Success.", "Status Code", resp.StatusCode())
	return nil
}
