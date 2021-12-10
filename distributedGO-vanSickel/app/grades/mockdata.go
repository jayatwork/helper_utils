package grades

func init() {
	students = []Student {
		Student{
			ID:			1,
			LastName:	"Averill",
			Grades:	[]Grade{
				Grade{
					Title:	"Quiz 1"
					Type:	GradeQuiz,
					Score:	100,
				},
				Grade{
					Title:	"Week 1 Homework"
					Type:	GradeHomework,
					Score:	85,
				},
				Grade{
					Title:	"Quiz 2"
					Type:	GradeQuiz,
					Score:	70,
				},
			},
		},
		Student{
			ID:			1,
			LastName:	"Garrar",
			Grades:	[]Grade{
				Grade{
					Title:	"Quiz 1"
					Type:	GradeQuiz,
					Score:	23,
				},
				Grade{
					Title:	"Week 1 Homework"
					Type:	GradeHomework,
					Score:	58,
				},
				Grade{
					Title:	"Quiz 2"
					Type:	GradeQuiz,
					Score:	105,
				},
			},
		},
		Student{
			ID:			1,
			LastName:	"Poopa",
			Grades:	[]Grade{
				Grade{
					Title:	"Quiz 1"
					Type:	GradeQuiz,
					Score:	85,
				},
				Grade{
					Title:	"Week 1 Homework"
					Type:	GradeHomework,
					Score:	79,
				},
				Grade{
					Title:	"Quiz 2"
					Type:	GradeQuiz,
					Score:	99,
				},
			},
		},
	}



}