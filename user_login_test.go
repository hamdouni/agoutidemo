package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

var _ = Describe("UserLogin", func() {
	var page *agouti.Page

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	It("does some tests", func() {
		By("showing index", func() {
			Expect(page.Navigate(appURL)).To(Succeed())
			Expect(page).Should(HaveTitle("AGOUTI DEMO"))
		})
		By("navigating to end", func() {
			Expect(page.FindByLink("go to end").Click()).To(Succeed())
			Expect(page).To(HaveTitle("THE END"))
		})
	})
	It("does some dynamic stuff", func() {
		By("loging in", func() {
			Expect(page.Navigate(appURL)).To(Succeed())
			Expect(page.Find("#userid").Fill("username")).To(Succeed())
			Expect(page.Find("#pass").Fill("a password")).To(Succeed())
			Expect(page.Find("#myform").Submit()).To(Succeed())
			Expect(page).To(HaveTitle("THE END"))
		})
	})
})
