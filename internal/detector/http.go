package detector

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func JSONHelper(w http.ResponseWriter, statusCode int, output interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	bts, _ := json.Marshal(output)
	w.Write(bts)
}

func GetTaskHandler(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := chi.URLParam(r, "uuid")
		output := GetTaskOutput{
			UUID:  uuid,
			OK:    false,
			Error: "",
		}

		data, err := s.getScanInfo(uuid)
		if err != nil {
			output.Error = err.Error()
			JSONHelper(w, http.StatusInternalServerError, output)
			return
		}

		output.OK = true
		output.Result = data.Result
		JSONHelper(w, http.StatusOK, output)
	}
}

func PostTaskHandler(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		task := Task{UUID: uuid.NewV4().String()}
		output := PostTaskOutput{
			UUID:  task.UUID,
			OK:    false,
			Error: "",
		}

		bts, err := ioutil.ReadAll(r.Body)
		if err != nil {
			output.Error = err.Error()
			JSONHelper(w, http.StatusInternalServerError, output)
			return
		}

		input := Input{}
		err = json.Unmarshal(bts, &input)
		if err != nil {
			output.Error = err.Error()
			JSONHelper(w, http.StatusInternalServerError, output)
			return
		}

		task.HostList = input

		err = s.addToTaskQueue(task)
		if err != nil {
			output.Error = err.Error()
			JSONHelper(w, http.StatusInternalServerError, output)
			return
		}

		output.OK = true
		bts, _ = json.Marshal(output)

		w.WriteHeader(http.StatusOK)
		w.Write(bts)
	}
}

func GetRoutes(s *Service) http.Handler {
	router := chi.NewRouter()
	router.Post("/task", PostTaskHandler(s))
	router.Get("/task/{uuid}", GetTaskHandler(s))

	return router
}

func startHttpServer(s *Service, wg *sync.WaitGroup) error {
	defer wg.Done()
	route := GetRoutes(s)

	log.Println("Start http-server on :8080")
	server := &http.Server{
		Addr:           "0.0.0.0:8080",
		Handler:        route,
		ReadTimeout:    3 * time.Second,
		WriteTimeout:   3 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return server.ListenAndServe()
}
