package shelper

import (
	"fmt"
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/model"
	"lms-backend/internal/policy/abilities"
	"lms-backend/internal/policy/roles"
	"lms-backend/util/sliceutil"

	"gorm.io/gorm"
)

func SeedRoleAndAbility(db *gorm.DB) error {
	var abts = abilities.GetAllAbilities()
	result := db.Create(&abts)
	if result.Error != nil {
		return result.Error
	}

	var rls = roles.GetAllRoles()
	result = db.Create(&rls)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func LinkRoleAndAbility(db *gorm.DB) error {
	var (
		rbMap = map[string][]string{
			roles.SystemAdmin.Name: {
				abilities.CanManageAll.Name,
			},

			roles.LibraryAdmin.Name: {
				abilities.CanReadAuditLog.Name,
				abilities.CanCreateAuditLog.Name,

				abilities.CanCreateUser.Name,
				abilities.CanReadUser.Name,
				abilities.CanUpdateUser.Name,
				abilities.CanDeleteUser.Name,
				abilities.CanUpdateUserRole.Name,

				abilities.CanCreatePerson.Name,
				abilities.CanUpdatePerson.Name,

				abilities.CanReadBook.Name,
				abilities.CanCreateBook.Name,
				abilities.CanUpdateBook.Name,
				abilities.CanDeleteBook.Name,

				abilities.CanLoanBook.Name,
				abilities.CanReturnBook.Name,
				abilities.CanRenewBook.Name,

				abilities.CanReadReservation.Name,
				abilities.CanCreateReservation.Name,
				abilities.CanCancelReservation.Name,

				abilities.CanReadFine.Name,
				abilities.CanSettleFine.Name,
				abilities.CanDeleteFine.Name,

				abilities.CanReadBookMark.Name,
				abilities.CanCreateBookMark.Name,
				abilities.CanDeleteBookMark.Name,

				abilities.CanManageBookRecords.Name,
			},

			roles.Staff.Name: {
				abilities.CanReadUser.Name,
				abilities.CanCreateUser.Name,

				abilities.CanReadBook.Name,

				abilities.CanLoanBook.Name,
				abilities.CanReturnBook.Name,
				abilities.CanRenewBook.Name,

				abilities.CanReadReservation.Name,
				abilities.CanCreateReservation.Name,
				abilities.CanCancelReservation.Name,

				abilities.CanReadFine.Name,
				abilities.CanSettleFine.Name,
				abilities.CanDeleteFine.Name,

				abilities.CanReadBookMark.Name,
				abilities.CanCreateBookMark.Name,
				abilities.CanDeleteBookMark.Name,

				abilities.CanManageBookRecords.Name,
			},

			roles.Basic.Name: {
				abilities.CanReadBook.Name,
				abilities.CanCreateBookMark.Name,
			},
		}
	)

	// Assign abilities to roles
	for _, role := range roles.GetAllRoles() {
		for _, ability := range abilities.GetAllAbilities() {
			if !sliceutil.Contains(rbMap[role.Name], ability.Name) {
				continue
			}

			if err := db.Exec(
				fmt.Sprintf("%s %s %s %s",
					"INSERT INTO role_abilities (role_id, ability_id)",
					"SELECT",
					"(SELECT id FROM roles WHERE name = ?),",
					"(SELECT id FROM abilities WHERE name = ?)",
				),
				role.Name, ability.Name,
			).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func LinkUserWithRoles(db *gorm.DB) error {
	var sysAdminRoleID int64

	// Get sysadmin role id
	if err := db.Model(&model.Role{}).
		Select("id").
		Where("name = ?", roles.SystemAdmin.Name).
		First(&sysAdminRoleID).
		Error; err != nil {
		return err
	}

	// Assign all abilities to sysadmin
	if _, err := user.UpdateRoles(db,
		int64(Users[0].ID),
		sysAdminRoleID,
	); err != nil {
		return err
	}

	var libAdminRoleID int64

	// Get libadmin role id
	if err := db.Model(&model.Role{}).
		Select("id").
		Where("name = ?", roles.LibraryAdmin.Name).
		First(&libAdminRoleID).
		Error; err != nil {
		return err
	}

	// Assign all abilities to libadmin
	if _, err := user.UpdateRoles(db,
		int64(Users[1].ID),
		libAdminRoleID,
	); err != nil {
		return err
	}

	var staffRoleID int64

	// Get staff role id
	if err := db.Model(&model.Role{}).
		Select("id").
		Where("name = ?", roles.Staff.Name).
		First(&staffRoleID).
		Error; err != nil {
		return err
	}

	// Assign all abilities to staff
	if _, err := user.UpdateRoles(db,
		int64(Users[2].ID),
		staffRoleID,
	); err != nil {
		return err
	}

	for _, usr := range Users[3:] {
		if _, err := user.UpdateRoles(db,
			int64(usr.ID),
			model.MemberRole,
		); err != nil {
			return err
		}
	}

	return nil
}
