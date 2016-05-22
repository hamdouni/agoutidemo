package main_test

import (
	"strconv"

	. "agoutidemo"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"

	"testing"
)

func TestAgoutidemo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Agoutidemo Suite")
}

var agoutiDriver *agouti.WebDriver
var appPort int
var appURL string

var _ = BeforeSuite(func() {
	// Choose a WebDriver:

	agoutiDriver = agouti.PhantomJS()
	// agoutiDriver = agouti.Selenium()
	// agoutiDriver = agouti.ChromeDriver()

	Expect(agoutiDriver.Start()).To(Succeed())

	appPort = 9999
	appURL = "http://localhost:" + strconv.Itoa(appPort)
	go StartApp(appPort)

})

var _ = AfterSuite(func() {
	Expect(agoutiDriver.Stop()).To(Succeed())
})
