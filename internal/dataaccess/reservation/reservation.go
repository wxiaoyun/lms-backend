package reservation

import (
	"lms-backend/internal/model"
	"lms-backend/internal/orm"
	"lms-backend/pkg/error/externalerrors"
	"time"

	"gorm.io/gorm"
)

func preloadBookUserAssociations(db *gorm.DB) *gorm.DB {
	return db.
		Preload("BookCopy").
		Preload("BookCopy.Book").
		Preload("User").
		Preload("User.Person")
}

func Read(db *gorm.DB, reservationID int64) (*model.Reservation, error) {
	var reservation model.Reservation

	result := db.Model(&model.Reservation{}).
		Where("id = ?", reservationID).
		First(&reservation)

	if err := result.Error; err != nil {
		if orm.IsRecordNotFound(err) {
			return nil, orm.ErrRecordNotFound(model.ReservationModelName)
		}
		return nil, err
	}

	return &reservation, nil
}

func ReadDetailed(db *gorm.DB, reservationID int64) (*model.Reservation, error) {
	var reservation model.Reservation

	result := db.Model(&model.Reservation{}).
		Scopes(preloadBookUserAssociations).
		Where("id = ?", reservationID).
		First(&reservation)

	if err := result.Error; err != nil {
		if orm.IsRecordNotFound(err) {
			return nil, orm.ErrRecordNotFound(model.ReservationModelName)
		}
		return nil, err
	}

	return &reservation, nil
}

func Delete(db *gorm.DB, reservationID int64) (*model.Reservation, error) {
	reservation, err := Read(db, reservationID)
	if err != nil {
		return nil, err
	}

	if err := reservation.Delete(db); err != nil {
		return nil, err
	}

	return ReadDetailed(db, reservationID)
}

func Count(db *gorm.DB) (int64, error) {
	var count int64

	result := orm.CloneSession(db).
		Model(&model.Reservation{}).
		Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func List(db *gorm.DB) ([]model.Reservation, error) {
	var rvs []model.Reservation

	result := db.Model(&model.Reservation{}).
		Find(&rvs)
	if result.Error != nil {
		return nil, result.Error
	}

	return rvs, nil
}

// Returns slice of reservations that is pending and reservation date is after now
func ReadByBookID(db *gorm.DB, bookID int64) ([]model.Reservation, error) {
	var reservations []model.Reservation

	result := db.Model(&model.Reservation{}).
		Where("book_id = ?", bookID).
		Order("created_at DESC").
		Find(&reservations)

	if result.Error != nil {
		return nil, result.Error
	}

	return reservations, nil
}
func ListWithBookUser(db *gorm.DB) ([]model.Reservation, error) {
	var res []model.Reservation

	result := db.Model(&model.Reservation{}).
		Scopes(preloadBookUserAssociations).
		Find(&res)
	if result.Error != nil {
		return nil, result.Error
	}

	return res, nil
}

// Returns slice of reservations that is pending and reservation date is after now
func ReadOutstandingReservationsByBookID(db *gorm.DB, bookID int64) ([]model.Reservation, error) {
	var reservations []model.Reservation

	result := db.Model(&model.Reservation{}).
		Where("book_id = ?", bookID).
		Where("status = ?", model.ReservationStatusPending).
		Where("reservation_date >= NOW()").
		Order("created_at DESC").
		Find(&reservations)

	if result.Error != nil {
		return nil, result.Error
	}

	return reservations, nil
}

// Returns slice of reservations that is pending and reservation date is after now
func ReadOutstandingReservationsByUserID(db *gorm.DB, userID int64) ([]model.Reservation, error) {
	var reservations []model.Reservation

	result := db.Model(&model.Reservation{}).
		Where("user_id = ?", userID).
		Where("status = ?", model.ReservationStatusPending).
		Where("reservation_date >= NOW()").
		Order("created_at DESC").
		Find(&reservations)

	if result.Error != nil {
		return nil, result.Error
	}

	return reservations, nil
}

func CountOutstandingReservationsByUserID(db *gorm.DB, userID int64) (int64, error) {
	var count int64

	result := db.Model(&model.Reservation{}).
		Where("user_id = ?", userID).
		Where("status = ?", model.ReservationStatusPending).
		Where("reservation_date >= NOW()").
		Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

// Assumes that the book is available.
//
// Relevant checks should be done before calling this function.
//
// User should not have more than maximum reservations and loans.
//
// Book should be neither on loan nor on reserve.
func ReserveBook(db *gorm.DB, userID, copyID int64) (*model.Reservation, error) {
	reservation := &model.Reservation{
		UserID:          uint(userID),
		BookCopyID:      uint(copyID),
		Status:          model.ReservationStatusPending,
		ReservationDate: time.Now().Add(model.ReservationDuration),
	}

	if err := reservation.Create(db); err != nil {
		return nil, err
	}

	return ReadDetailed(db, int64(reservation.ID))
}

// Sets the status of the reservation to fulfilled.
//
// This can either be the user retrieving the book or canceling the reservation.
func FullfilReservation(db *gorm.DB, reservationID int64) (*model.Reservation, error) {
	reservation, err := ReadDetailed(db, reservationID)
	if err != nil {
		return nil, err
	}

	if reservation.Status != model.ReservationStatusPending {
		return nil, externalerrors.BadRequest("reservation is not pending")
	}

	reservation.Status = model.ReservationStatusFulfilled
	if err := reservation.Update(db); err != nil {
		return nil, err
	}

	return reservation, nil
}
