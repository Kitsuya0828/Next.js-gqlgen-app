package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	// "github.com/Kitsuya0828/Next.js-gqlgen-app/backend/graph/model"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	// todos []*model.Todo
	DB    gorm.DB
}
