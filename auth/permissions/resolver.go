package permissions

import (
	"auth/domain"
	
)

// func ResolvePermissions(roles []domain.Role) []string {
func ResolvePermissions(roles []string) []string {
	permSet := make(map[string]struct{})

	for _, role := range roles {
		for _, p := range permissionsByRole(role) {
			permSet[p] = struct{}{}
		}
	}

	result := make([]string, 0, len(permSet))
	for p := range permSet {
		result = append(result, p)
	}

	return result
}

// func permissionsByRole(role domain.Role) []string {
func permissionsByRole(role string) []string {
	switch role {
		case string(domain.RoleStudent):
			// студент почти всегда работает "о себе"
			// permissions нужны только там, где доступ НЕ по умолчанию
			return []string{
				AnswerRead,
				AnswerUpdate,
				AnswerDel,
			}

		case string(domain.RoleTeacher):
			return []string{
				CourseInfoWrite,
				CourseTestList,
				CourseTestRead,
				CourseTestWrite,
				CourseTestAdd,
				CourseTestDel,
				CourseUserList,
				TestQuestAdd,
				TestQuestDel,
				TestQuestUpdate,
				TestAnswerRead,
				QuestListRead,
				QuestRead,
				QuestUpdate,
				QuestCreate,
				QuestDel,
			}

		case string(domain.RoleAdmin):
			return []string{
				UserListRead,
				UserFullNameWrite,
				UserDataRead,
				UserRolesRead,
				UserRolesWrite,
				UserBlockRead,
				UserBlockWrite,
				CourseAdd,
				CourseDel,
			}

		default:
			return nil
	}
}
