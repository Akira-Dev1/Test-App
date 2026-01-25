package domain

type Course struct {
    ID			*int    `json:"id,omitempty"`
    Title		*string `json:"title,omitempty"`
    Description	*string `json:"description,omitempty"`
    AuthorID    *string `json:"author_id,omitempty"`
    IsDeleted	*bool   `json:"is_deleted,omitempty"`
}

type Test struct {
    ID          *int    `json:"id,omitempty"`
    CourseID    *int    `json:"course_id,omitempty"`
    Title       *string `json:"title,omitempty"`
    IsActive    *bool   `json:"is_active,omitempty"`
    IsDeleted   *bool   `json:"is_deleted,omitempty"`
};