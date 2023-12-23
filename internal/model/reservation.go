package model

import (
	"lms-backend/pkg/error/externalerrors"
	"lms-backend/util/sliceutil"
	"time"

	"gorm.io/gorm"
)

type ReservationStatus = string

type Reservation struct {
	gorm.Model

	UserID          uint              `gorm:"not null"`
	User            *User             `gorm:"->"`
	BookID          uint              `gorm:"not null"`
	Book            *Book             `gorm:"->"`
	Status          ReservationStatus `gorm:"not null"`
	ReservationDate time.Time         `gorm:"not null"` // Date before which the book is reserved
}

const (
	ReservationModelName = "reservation"
	ReservationTableName = "reservations"
)

const (
	ReservationStatusPending   ReservationStatus = "pending"
	ReservationStatusFulfilled ReservationStatus = "fulfilled"
)

const (
	MaximumReservations = 2
	ReservationDuration = 7 * 24 * time.Hour
)

func (r *Reservation) Create(db *gorm.DB) error {
	return db.Create(r).Error
}

func (r *Reservation) Update(db *gorm.DB) error {
	return db.Updates(r).Error
}

func (r *Reservation) Delete(db *gorm.DB) error {
	return db.Delete(r).Error
}

func (r *Reservation) ensureUserExistsAndPresent(db *gorm.DB) error {
	if r.UserID == 0 {
		return externalerrors.BadRequest("user id is required")
	}

	var exists int64
	result := db.Model(&User{}).Where("id = ?", r.UserID).Count(&exists)
	if err := result.Error; err != nil {
		return err
	}

	if exists == 0 {
		return externalerrors.BadRequest("user does not exist")
	}

	return nil
}

func (r *Reservation) ensureBookExistsAndPresent(db *gorm.DB) error {
	if r.BookID == 0 {
		return externalerrors.BadRequest("book id is required")
	}

	var exists int64
	result := db.Model(&Book{}).Where("id = ?", r.BookID).Count(&exists)
	if err := result.Error; err != nil {
		return err
	}

	if exists == 0 {
		return externalerrors.BadRequest("book does not exist")
	}

	return nil
}

func (r *Reservation) ValidateStatus() error {
	if r.Status == "" {
		return externalerrors.BadRequest("status is required")
	}

	if !sliceutil.Contains([]ReservationStatus{
		ReservationStatusPending,
		ReservationStatusFulfilled,
	}, r.Status) {
		return externalerrors.BadRequest("invalid status")
	}

	return nil
}

func (r *Reservation) Validate(db *gorm.DB) error {
	if err := r.ensureUserExistsAndPresent(db); err != nil {
		return err
	}

	if err := r.ensureBookExistsAndPresent(db); err != nil {
		return err
	}

	if err := r.ValidateStatus(); err != nil {
		return err
	}

	if r.ReservationDate.IsZero() {
		return externalerrors.BadRequest("reservation date is required")
	}

	return nil
}

func (r *Reservation) BeforeCreate(db *gorm.DB) error {
	return r.Validate(db)
}

func (r *Reservation) BeforeUpdate(db *gorm.DB) error {
	return r.Validate(db)
}
