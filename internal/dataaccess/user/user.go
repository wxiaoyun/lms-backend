package user

import (
	"lms-backend/internal/dataaccess/loan"
	"lms-backend/internal/dataaccess/reservation"
	"lms-backend/internal/model"
	"lms-backend/internal/orm"
	"lms-backend/pkg/error/externalerrors"
	"lms-backend/util"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func preloadAssociations(db *gorm.DB) *gorm.DB {
	return db.Preload("Person")
}

func Read(db *gorm.DB, id int64) (*model.User, error) {
	var user model.User
	result := db.Model(&model.User{}).
		Scopes(preloadAssociations).
		Where("id = ?", id).
		First(&user)
	if err := result.Error; err != nil {
		if orm.IsRecordNotFound(err) {
			return nil, orm.ErrRecordNotFound(model.UserModelName)
		}
		return nil, err
	}

	return &user, nil
}

func ReadByUsername(db *gorm.DB, username string) (*model.User, error) {
	var user model.User
	result := db.Model(&model.User{}).
		Scopes(preloadAssociations).
		Where("username = ?", username).
		First(&user)
	if err := result.Error; err != nil {
		if orm.IsRecordNotFound(err) {
			return nil, orm.ErrRecordNotFound(model.UserModelName)
		}
		return nil, err
	}

	return &user, nil
}

func ReadByEmail(db *gorm.DB, email string) (*model.User, error) {
	var user model.User
	result := db.Model(&model.User{}).
		Scopes(preloadAssociations).
		Where("email = ?", email).
		First(&user)
	if err := result.Error; err != nil {
		if orm.IsRecordNotFound(err) {
			return nil, orm.ErrRecordNotFound(model.UserModelName)
		}
		return nil, err
	}

	return &user, nil
}

func Create(db *gorm.DB, user *model.User) (*model.User, error) {
	if err := user.Create(db); err != nil {
		return nil, err
	}

	return user, nil
}

func Update(db *gorm.DB, user *model.User) (*model.User, error) {
	if user.Person != nil && user.Person.ID != 0 {
		if err := user.Person.Update(db); err != nil {
			return nil, err
		}
	} else {
		if err := user.Person.Create(db); err != nil {
			return nil, err
		}
	}

	if err := user.Update(db); err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateParticulars(db *gorm.DB, user *model.User) (*model.User, error) {
	originalUser, err := Read(db, int64(user.ID))
	if err != nil {
		return nil, err
	}

	// Password will not be updated here
	user.EncryptedPassword = originalUser.EncryptedPassword
	return Update(db, user)
}

func Delete(db *gorm.DB, id int64) (*model.User, error) {
	usr, err := Read(db, id)
	if err != nil {
		return nil, err
	}

	if err := usr.Delete(db); err != nil {
		return nil, err
	}

	return usr, nil
}

func Login(db *gorm.DB, user *model.User) (*model.User, error) {
	var userInDB model.User

	if user.Username != "" {
		usr, err := ReadByUsername(db, user.Username)
		if err != nil {
			return nil, err
		}
		userInDB = *usr
	} else if user.Email != "" {
		usr, err := ReadByEmail(db, user.Email)
		if err != nil {
			return nil, err
		}
		userInDB = *usr
	} else {
		return nil, externalerrors.BadRequest("Email or Username is required")
	}

	util.Debug(user.EncryptedPassword)

	err := bcrypt.CompareHashAndPassword([]byte(userInDB.EncryptedPassword), []byte(user.EncryptedPassword))
	if err != nil {
		return nil, externalerrors.Unauthorized("user not found or invalid password")
	}

	userInDB.LastSignInAt = userInDB.CurrentSignInAt
	userInDB.CurrentSignInAt = time.Now()
	userInDB.SignInCount++

	return Update(db, &userInDB)
}

func GetUserName(db *gorm.DB, id int64) (string, error) {
	var name string
	result := db.Model(&model.User{}).
		Select("username").
		Where("id = ?", id).
		First(&name)
	if err := result.Error; err != nil {
		if orm.IsRecordNotFound(err) {
			return "", orm.ErrRecordNotFound(model.UserModelName)
		}
		return "", err
	}

	return name, nil
}

func UpdateRoles(db *gorm.DB, userID int64, roleIDs []int64) (*model.User, error) {
	usr, err := Read(db, userID)
	if err != nil {
		return nil, err
	}

	if err := usr.UpdateRoles(db, roleIDs); err != nil {
		return nil, err
	}

	return usr, nil
}

func GetAbilities(db *gorm.DB, userID int64) ([]model.Ability, error) {
	var abilities []model.Ability

	result := db.
		Model(&model.Ability{}).
		Joins("JOIN role_abilities ON role_abilities.ability_id = abilities.id").
		Joins("JOIN user_roles ON user_roles.role_id = role_abilities.role_id").
		Where("user_roles.user_id = ?", userID).
		Order("abilities.name ASC").
		Find(&abilities)

	if result.Error != nil {
		return nil, result.Error
	}

	return abilities, nil
}

func GetRoles(db *gorm.DB, userID int64) ([]model.Role, error) {
	var roles []model.Role

	result := db.
		Model(&model.Role{}).
		Joins("JOIN user_roles ON user_roles.role_id = roles.id").
		Where("user_roles.user_id = ?", userID).
		Order("roles.id ASC").
		Find(&roles)

	if result.Error != nil {
		return nil, result.Error
	}

	util.Debug(roles)

	return roles, nil
}

func HasExceededMaxLoan(db *gorm.DB, userID int64) (bool, error) {
	loans, err := loan.ReadOutstandingLoansByUserID(db, userID)
	if err != nil {
		return false, err
	}

	return len(loans) >= model.MaximumLoans, nil
}

func HasExceededMaxReservation(db *gorm.DB, userID int64) (bool, error) {
	reservations, err := reservation.ReadOutstandingReservationsByUserID(db, userID)
	if err != nil {
		return false, err
	}

	return len(reservations) > model.MaximumReservations, nil
}
