package databases

import "learning/app/model"

func migration() {
	DB.AutoMigrate(&model.User{})
}
