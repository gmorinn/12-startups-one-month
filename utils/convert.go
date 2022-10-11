package utils

import (
	db "12-startups-one-month/internal"
)

// convert []db.Role to []string
func ConvertRoleToString(roles []db.Role) []string {
	var res []string
	for _, role := range roles {
		res = append(res, string(role))
	}
	return res
}

// convert []string to []db.Role
func ConvertStringToRole(roles []string) []db.Role {
	var res []db.Role
	for _, role := range roles {
		res = append(res, db.Role(role))
	}
	return res
}
