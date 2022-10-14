package utils

import (
	"12-startups-one-month/graph/model"
	db "12-startups-one-month/internal"
	"strconv"
	"strings"
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

// convert []db.Role to []model.UserType
func ConvertRoleToUserType(roles []db.Role) []model.UserType {
	var res []model.UserType
	for _, role := range roles {
		res = append(res, model.UserType(strings.ToLower(string(role))))
	}
	return res
}

// convert string to int using strconv
func StrToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return i
}
