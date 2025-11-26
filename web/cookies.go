package web

import (
	"context"
	"net/http"
	"os"

	infa "github.com/Knowckx/infa"
	"github.com/Knowckx/infa/util"
	"github.com/browserutils/kooky"
	"github.com/browserutils/kooky/browser/chrome"
	"github.com/rs/zerolog/log"
)

func GetHostCookies(host string) []*http.Cookie {
	koo := GetChromeCookies()
	outs := []*http.Cookie{}
	for _, ko := range koo {
		if ko.Domain == host {
			ht := ko.Cookie
			outs = append(outs, &ht)
		}
	}
	return outs
}

func GetChromeCookies() []*kooky.Cookie {
	dir, _ := os.UserConfigDir() // "/<USER>/Library/Application Support/"
	cookiesFile := dir + "/Google/Chrome/Default/Cookies"
	cookies, err := chrome.ReadCookies(context.TODO(), cookiesFile)
	if err != nil {
		log.Error().Stack().Err(err).Send()
	}
	return cookies
}

func PrintCookies(cks []*http.Cookie) {
	infa.Printf("Domain Name value Path -- len %d", len(cks))
	for _, ck := range cks {
		PrintCookie(ck)
	}
}

func PrintCookie(ck *http.Cookie) {
	val := util.ShortStr(ck.Value)
	infa.Printf("%s %s %s %s", ck.Domain, ck.Name, val, ck.Path)
}
