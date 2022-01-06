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
	mutex 		*sync.RWMutex
}

func (r *registry) add(reg Registration) error {
	r.mutex.Lock()
	r.registration = append(r.registration, reg)
	r.mutex.Unlock()
	err := r.sendRequiredServices(reg)
	return nil
}

func (r registry) sendRequiredServices(reg Registration) error {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var p patch
	for _, serviceReg := range r.registrations {
		for _, reqService := range reg.RequiredServices {
			if serviceReg.ServiceName == reqService {
				p.Added = append(p.Added, patchEntry {
					Name: serviceReg.ServiceName
					URL: serviceReg.ServiceURL,
				})
			}
		}
	}
	err := r.sendPatch(p, reg.SendUpdateURL)
	if err != nil {
		return err
	}
	return nil
}

func (r registry) sendPatch(p patch, url string) error {
	d, err := json.Marshal(p)
	if err != nil {
		return err
	}
	_, err = http.Post(url, "application/json", bytes.NewBuffer(d))
	if err != nil {
		return err
	}
	return nil
}

var reg = registry{registration: make([]Registration, 0),
	mutex: new(sync.RWMutex),
}

type RegistryService struct {}

func (s RegistryService) ServeHTTP(w http.ResponseWriter, req *http.Request) {
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
