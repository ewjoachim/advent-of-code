package boilerplate

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/zellyn/kooky"
	_ "github.com/zellyn/kooky/firefox"
)

func Fetch(year, day uint) *bytes.Buffer {

	cookies := kooky.ReadCookies(
		kooky.Valid,
		kooky.DomainHasSuffix(`adventofcode.com`),
		kooky.Name(`session`),
	)
	if len(cookies) == 0 {
		panic("no cookies found")
	}
	cookie := cookies[0]
	client := &http.Client{}
	aocUrl, err := url.Parse("https://adventofcode.com")
	if err != nil {
		panic(err)
	}
	jarCookie := &http.Cookie{
		Name:  cookie.Name,
		Value: cookie.Value,
	}

	client.Jar, err = cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	client.Jar.SetCookies(aocUrl, []*http.Cookie{jarCookie})
	resp, err := client.Get(fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(resp.Body)
	if err != nil {
		panic(err)
	}
	return buffer
}
