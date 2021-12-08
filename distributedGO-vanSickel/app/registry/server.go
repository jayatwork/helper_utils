package registry

import ( 
	"sync"
	"net/http"
	"log"
	"encoding/json"
)

const ServerPort = ":3000"
const ServicesURL= "http://localhost" + ServerPort + "/services"

type registry struct {
	registration []Registration
	mutex 		*sync.Mutex
}

func (r *registry) add(reg Registration) error {
	r.mutex.Lock()
	r.registration = append(r.registration, reg)
	r.mutex.Unlock()
	return nil
}

var reg = registry{registration: make([]Registration, 0),
	mutex: new(sync.Mutex),
}

type RegistryService struct {}

func (s RegistryService) ServerHTTP(w http.ResponseWriter, req *http.Request) {
	log.Println("Request Received")
	switch req.Method {
	case http.MethodPost:
		dec := json.NewDecoder(req.Body)
		var r Registration
		err := dec.Decode(&r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	default:
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
	}
}
