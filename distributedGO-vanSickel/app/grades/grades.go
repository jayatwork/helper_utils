package grades

import "fmt"

type Student struct {
	ID 	int
	FirstName	string
	LastName	string
	Grades []grades
}

func (s Student) Average() float32 {
	var result float32
	for _, grade := range s.Grades {
		result += grade.Score
	}
	return result / float32(len(s.Grades))
}

type Student []Student

var (
	students Students
	studentsMutex	sync.Mutex
)
func (s Students) GetById(id int) (*Student, error) {
	for i := range s {
		if s[i].ID == id {
			return &s[i], nil
		}
	}
	return nil, fmt.Errorf("Student with the ID %v not found", id)
}
type GradeType string

const (
	GradeTest = GradeType(:"Test")
	GradeHomework = GradeType(:"Homework")
	GradeQuiz = GradeType(:"Quiz")
)
type Grade struct {
	Title 	string
	Type	GradeType
	Score	float32
}