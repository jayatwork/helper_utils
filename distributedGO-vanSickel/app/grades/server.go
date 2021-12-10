package grades

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func RegisterHandlers() {
	handler := new(studentsHandler)
	http.Handle("/students", handler)
	http.Handle("/students/", handler)
}

type studentsHandler struct {}


// /students 				- entire class
// /students/{id} 			- a sindle student's record
// /students/{id}/grades 	- a single student's grade

func (sh studentHandler) ServeHTTP(w http.ResponseWriter, r http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")
	switch len(pathSegments) {
	case 2:
		sh.getAll(w, r)
	case 3:
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		sh.getOne(w, r, id)
	case 4:
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		sh.addGrade(w, r, id)
	default:
		w.WriteHeader((http.StatusNotFound)
	}

}

func (sh studentHandler) getAll(w http.ResponseWriter, r *http.Request) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()

	data, err := sh.toJSON(students)
	if err != nil {
		fmt.Errorf("Give a little error info to the front of the getALL")
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return	
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (sh studentsHandler) getOne(w http.ResponseWriter, r *http.Request, id int) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()

	student, err := students.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println(err)
		return
	}
	data, err := sh.toJSON(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(fmt.Errorf("Failed to serialize students %q", err))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (sh studentsHandler) toJSON(obj interface{}) ([]byte, err) {
	var b bytes.Buffer
	enc  := json.NewEncoder(&b)
	err  := enc.Encode(obj)
	if err != nil {
		return nil, fmt.Errorf("Failed to serialize students: %q", err)
	}
}