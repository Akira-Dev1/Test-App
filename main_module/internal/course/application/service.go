package application

import (
	courseDomain "main_module/internal/course/domain"
)

type CourseService struct {
	Repo courseDomain.Repository
}

func NewCourseService(repo courseDomain.Repository) *CourseService {
	return &CourseService{
		Repo: repo,
	}
}