package shelper

import (
	"fmt"
	"lms-backend/internal/model"
	"lms-backend/util/random"

	"gorm.io/gorm"
)

const (
	SystemAdminEmail  = "admin@gmail.com"
	LibraryAdminEmail = "libadmin@gmail.com"
	StaffEmail        = "staff@gmail.com"
)

var (
	Users []model.User = []model.User{}
)

func SeedUsersAndPeople(db *gorm.DB, num int64) error {
	Users = make([]model.User, num)

	var count int64

	result := db.Model(&model.User{}).Count(&count)
	if result.Error != nil {
		return result.Error
	}

	if count >= num {
		return nil
	}

	// Generate 100 users
	Users[0] = model.User{
		Username:          "admin",
		Email:             SystemAdminEmail,
		EncryptedPassword: "P4ssw0rd!",
		Person: &model.Person{
			FullName:           "Admin",
			PreferredName:      "Admin",
			LanguagePreference: "English",
		},
	}
	Users[1] = model.User{
		Username:          "libadmin",
		Email:             LibraryAdminEmail,
		EncryptedPassword: "P4ssw0rd!",
		Person: &model.Person{
			FullName:           "Library Admin",
			PreferredName:      "Library Admin",
			LanguagePreference: "English",
		},
	}
	Users[2] = model.User{
		Username:          "staff",
		Email:             StaffEmail,
		EncryptedPassword: "P4ssw0rd!",
		Person: &model.Person{
			FullName:           "Staff",
			PreferredName:      "Staff",
			LanguagePreference: "English",
		},
	}

	for i := 4; i <= int(num); i++ {
		Users[i-1] = model.User{
			Username:          fmt.Sprintf("user%d", i),
			Email:             fmt.Sprintf("user%d@gmail.com", i),
			EncryptedPassword: "P4ssw0rd!",
			Person: &model.Person{
				FullName:           random.RandWords(random.RandInt(2, 10)),
				PreferredName:      random.RandWords(random.RandInt(2, 10)),
				LanguagePreference: random.RandWords(random.RandInt(1, 3)),
			},
		}
	}

	return db.Create(&Users).Error
}
