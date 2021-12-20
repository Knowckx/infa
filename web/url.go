package web

import (
	"net/url"

	"github.com/pkg/errors"
)

// url like xxx.com?a=b
func GetUrlParasKey(urlStr string, key string) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", errors.WithStack(err)
	}
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return "", errors.WithStack(err)
	}
	val, ok := m[key]
	if !ok {
		return "", nil
	}
	if len(val) == 0 {
		return "", errors.Errorf("length of slice is 0")
	}
	return val[0], nil
}
