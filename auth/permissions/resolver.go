package permissions

import "auth/domain"

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
	case "student":
		return []string{
			ProfileRead,
			CoursesRead,
		}
	case string(domain.RoleAdmin):
		return []string{
			ProfileRead,
			ProfileUpdate,
			CoursesRead,
			AdminPanel,
		}
	default:
		return nil
	}
}
