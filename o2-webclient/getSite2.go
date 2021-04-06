package main

import "github.com/java-akademie/myutils"

func doGetSite2() {
	myutils.H1("getSite2")

	getSite("http://google22.com", false)
	getSite("http://jmildner.ch", true)
	getSite("http://jmildner.ch/xxx", false)
	getSite("http://java-akademie.ch", true)
	getSite("http://java-akademie.ch/kurse.shtml", true)
	getSite("http://java-akademie.ch/xxx", false)
	getSite("http://google.com", false)
	getSite("http://facebook.com", false)
	getSite("http://swisscom.ch", false)
	getSite("http://udemy.com", false)
}
