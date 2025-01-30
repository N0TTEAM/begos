package model

import "github.com/N0TTEAM/begos/internal/http/model/model1"

func GetAllModels() []interface{} {
	return []interface{}{
		&model1.User{},
		&Tes{},
	}
}
