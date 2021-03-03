package detector

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

var workingDir = ""

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "License test")
}

var _ = BeforeSuite(func() {
})

var _ = Describe("Main check test", func() {
	Context("Default configuration", func() {
		It("Check if probe works", func() {
			data := []byte(`["test.com"]`)
			input := bytes.NewReader(data)
			req, _ := http.NewRequest("POST", "/task", input)

			dep := NewService(1)
			dep.probes = []Probe{
				&ProbeMock{
					RunFunc: func(host string) bool {
						return true
					},
				},
			}

			setCalledForUUID := ""
			dep.output = &StorageMock{
				GetFunc: nil,
				SetFunc: func(uuid string, result interface{}) error {
					setCalledForUUID = uuid
					return nil
				},
			}

			dep.startWorkerPool(1, &sync.WaitGroup{})

			handler := GetRoutes(dep)

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			Expect(rr.Result().StatusCode).To(Equal(http.StatusOK))

			output := PostTaskOutput{}
			err := json.Unmarshal(rr.Body.Bytes(), &output)
			Expect(err).NotTo(HaveOccurred())
			Expect(output.UUID).NotTo(Equal(""))
			Expect(setCalledForUUID).To(Equal(output.UUID))

		})
	})
})

var _ = AfterSuite(func() {
})
