package go_micro_srv_user

import (
	"github.com/jinzhu/gorm"

	"github.com/pborman/uuid"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {

	id := uuid.NewUUID()

	return scope.SetColumn("Id", id.String())
}