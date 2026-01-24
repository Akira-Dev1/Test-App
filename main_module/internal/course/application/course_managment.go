package application

import (
	"errors"

	authDomain "main_module/internal/auth/domain"
)

//GetDisciplineByID (Посмотреть информацию — Default Allow)
//CreateDiscipline (Создать — Permission: course:add)
//UpdateDiscipline (Изменить информацию — Owner or course:info:write)
//DeleteDiscipline (Удалить — Owner or course:del)


// Получить список курсов
func (s *CourseService) GetCourses(user *authDomain.UserContext) ([]map[string]any, error) {
	rule := AccessRule{
		DefaultAllow: true,
	}

	if !CheckAccess(user, rule, "") {
		return nil, errors.New("forbidden")
	}

	return s.Repo.GetAll()
}

func (s *CourseService) GetCourseByID(user *authDomain.UserContext, id string) (map[string]any, error) {
	rule := AccessRule{
		DefaultAllow: true,
	}

	if !CheckAccess(user, rule, "") {
		return nil, errors.New("forbidden")
	}

	return s.Repo.GetByID(id)
}
