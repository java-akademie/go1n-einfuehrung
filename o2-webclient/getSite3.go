package main

import "github.com/java-akademie/myutils"

func doGetSite3() {
	myutils.H1("getSite3")

	getSite("http://localhost:8081", true)
	getSite("http://localhost:8082", true)
	getSite("http://localhost:8083", true)
	getSite("http://localhost:8083/xxx/yyy?name=hugo&farbe=rot&farbe=blau", true)
	getSite("http://localhost:8083?name=hugo&farbe=rot&farbe=blau", true)
	getSite("http://localhost:8084", true)
	getSite("http://localhost:8084/static/test1.html", true)
	getSite("http://localhost:8084/static/test2.html", true)
}
