package utils

import (
	"12-startups-one-month/graph/model"
	db "12-startups-one-month/internal"
	"strings"
)

// check if the string is in the array
func InArray(str string, arr []string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

// check if current user has one of the role needed
func HasRole(roles []db.Role, needRoles []model.UserType) bool {
	var roleArrNeeded []string
	for _, v := range needRoles {
		roleArrNeeded = append(roleArrNeeded, strings.ToLower(string(v)))
	}
	for _, role := range roles {
		if InArray(string(role), roleArrNeeded) {
			return true
		}
	}
	return false
}
