package main_module


// Для курсов
type Course struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	AuthorID    string `json:"author_id,omitempty"`
}

type CoursesResponse struct {
	Courses []Course `json:"courses"`
}

type CoursePayload struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type StudentsListResponse struct {
	StudentIDs []string `json:"student_ids"`
}

type UserPayload struct {
	UserID string `json:"user_id"`
}


// Для тестов
type Test struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type TestsListResponse struct {
	Tests []Test `json:"tests"`
}

type TestStatusResponse struct {
	IsActive bool `json:"is_active"`
}



// Для юзеров
type UserProfileData struct {
	UserID  string `json:"user_id"`
	Courses any    `json:"courses,omitempty"`
	Tests   any    `json:"tests,omitempty"`
	Grades  any    `json:"grades,omitempty"`
}

type UserInfo struct {
	UserID   string `json:"user_id"`
	FullName string `json:"full_name"`
}

type UserRolesResponse struct {
	Roles []string `json:"roles"`
}



// Для вопросов
type Question struct {
	ID            int      `json:"id"`
	Version       int      `json:"version,omitempty"`
	AuthorID      string   `json:"author_id,omitempty"`
	Title         string   `json:"title"`
	Content       string   `json:"content,omitempty"`
	Options       []string `json:"options,omitempty"`
	CorrectOption int      `json:"correct_option,omitempty"`
}

type QuestionListResponse struct {
	Questions []Question `json:"questions"`
}




// Для попыток
type Score struct {
	UserID string  `json:"user_id"`
	Score  float64 `json:"score"`
}
type AnswerDetail struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type AttemptDetails struct {
	UserID  string         `json:"user_id"`
	Answers []AnswerDetail `json:"answers"`
}

type PassedUsersResponse struct {
	UserIDs []string `json:"user_ids"`
}

type ScoresResponse struct {
	Scores []Score `json:"scores"`
}

type AttemptsResponse struct {
	Attempts []AttemptDetails `json:"attempts"`
}
