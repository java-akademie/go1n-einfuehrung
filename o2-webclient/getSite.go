package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/java-akademie/myutils"
)

type isite struct {
	link        string
	proto       string
	contentType string
	status      string
	statusCode  int
	bodyLen     int
	body        string
}

var body string

func getSite(link string, showBody bool) {

	myutils.H2(link)
	myutils.Wait()

	is, err := getISite(link)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	showSite(is, showBody)
}

func showSite(is isite, showBody bool) {

	format := `
Proto.........: %v
StatusCode....: %v
Status........: %v
Link..........: %v
ContentType...: %v
BodyLen.......: %v

`
	fmt.Printf(
		format,
		is.proto,
		is.statusCode,
		is.status,
		is.link,
		is.contentType,
		is.bodyLen,
	)

	if showBody {
		fmt.Printf("Body: \n%v \n", is.body)
	}
}

func getISite(link string) (isite, error) {

	var is isite

	is.link = link

	resp, err := http.Get(link)
	if err != nil {
		return is, err
	}

	is.proto = resp.Proto
	is.status = resp.Status
	is.statusCode = resp.StatusCode
	is.contentType = fmt.Sprintf("%v", resp.Header["Content-Type"])

	w := writer{}
	bLen, _ := io.Copy(w, resp.Body)

	is.bodyLen = int(bLen)
	is.body = body

	return is, nil
}

type writer struct{}

// implements the interface Writer in io
func (writer) Write(bs []byte) (int, error) {
	body = fmt.Sprint(string(bs))
	return len(bs), nil
}
