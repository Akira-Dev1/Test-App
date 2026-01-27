package domain

import (
    "time"
)

type UsersIDS struct {
    IDS         []string    `json:"users_ids,omitempty"`
}


type UserData struct {
    Name		*string     `json:"name,omitempty"`
    Courses		[]Course   `json:"courses,omitempty"`
    Grades	    []Grade    `json:"assessments,omitempty"`
    Tests       []Test     `json:"tests,omitempty"`
    Roles       []string    `json:"roles,omitempty"`
    IsBlocked   *bool       `json:"is_blocked,omitempty"`
}



type Grade struct {
    AttemptID   *int        `json:"attempt_id,omitempty"`
    CourseTitle *string     `json:"course_title,omitempty"`
    TestTitle   *string     `json:"test_title,omitempty"` 
    Date        *time.Time  `json:"date,omitempty"` 
    Status      *string     `json:"status,omitempty"` 
    Score       *int        `json:"score,omitempty"` 
    MaxScore    *int        `json:"max_score,omitempty"` 
}
type Test struct {
    Title       *string     `json:"title,omitempty"` 
    TestID      *int        `json:"test_id,omitempty"` 
    CourseID    *int        `json:"course_id,omitempty"` 

}
type Course struct {
    Title       *string     `json:"title,omitempty"` 
    Description *string     `json:"description,omitempty"` 
    CourseID    *int        `json:"course_id,omitempty"` 
}



