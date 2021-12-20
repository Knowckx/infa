package web

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
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
	log.Debug().Str("url", url).Interface("params", queryParams).Msg("send request get")
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
	log.Debug().Str("url", url).Interface("params", queryParams).Msg("send request get")
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
	log.Debug().Str("post url,", url).Interface("body", body).Send()
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
		log.Error().Int("Status Code", resp.StatusCode()).Msg("Http Request Failed")
		data := resp.Body()
		log.Error().Str("Body", string(data)).Send()
		return fmt.Errorf("Http Request Failed")
	}
	log.Debug().Int("Status Code", resp.StatusCode()).Msg("Http Request Success.")
	return nil
}
