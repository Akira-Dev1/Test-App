package domain

type Test struct {
    ID          *int       	`json:"id,omitempty"`
    CourseID    *int      	`json:"course_id,omitempty"`
    Title       *string    	`json:"title,omitempty"`
    QuestionIDs []int     	`json:"question_ids,omitempty"`
    IsActive    *bool      	`json:"is_active,omitempty"`
    IsDeleted   *bool      	`json:"is_deleted,omitempty"`
}
type TestQuestionIDS struct {
    IDS         []int   	`json:"question_ids,omitempty"`
}
type StudentsIDS struct {
    IDS         []string   	`json:"students_ids,omitempty"`
}


// Оценки
type Score struct {
	UserID		*string		`json:"user_id,omitempty"`
	Score 		*float64	`json:"score,omitempty"`
}
type Scores struct {
	Scores 		[]Score		`json:"scores,omitempty"`
}


// Ответы
type Answer struct {
	QuestionID	*int		`json:"question_id,omitempty"`
	Answer		*int		`json:"answer,omitempty"`
}
type Attempt struct {
	UserID 		*string		`json:"user_id,omitempty"`
	AttemptID	*int		`json:"attempt_id,omitempty"`
	Answers 	[]Answer	`json:"answers,omitempty"`
}
type Attempts struct {
	Attempts 	[]Attempt	`json:"attempts,omitempty"`
}
