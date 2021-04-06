package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/java-akademie/myutils"
)

type iSite1 struct {
	link        string
	proto       string
	contentType string
	status      string
	statusCode  int
	body        string
	bodyLen     int
}

var is1 = iSite1{}

func doGetSite1() {
	myutils.H1("getSite1")
	myutils.Wait()

	is1.link = "http://jmildner.ch"

	resp, err := http.Get(is1.link)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	is1.proto = resp.Proto
	is1.status = resp.Status
	is1.statusCode = resp.StatusCode
	is1.contentType = fmt.Sprintf("%v", resp.Header["Content-Type"])

	wr := writer1{}
	// io.Copy fills via wr is1.body
	l, _ := io.Copy(wr, resp.Body)

	is1.bodyLen = int(l)

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
		is1.proto,
		is1.statusCode,
		is1.status,
		is1.link,
		is1.contentType,
		is1.bodyLen,
	)

	fmt.Printf("Body: \n%v \n", is1.body)

	myutils.Wait()
}

type writer1 struct{}

// implements the interface Writer in io
func (writer1) Write(bs []byte) (int, error) {
	is1.body = fmt.Sprint(string(bs))
	return len(bs), nil
}
