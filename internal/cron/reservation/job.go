package reservationjob

import (
	audit "lms-backend/internal/auditlog"
	"lms-backend/internal/dataaccess/bookcopy"
	"lms-backend/internal/model"
)

func DetectOverdueResAndCancelRes() {
	var err error

	tx, rollBackOrCommit := audit.Begin(nil, "CRON Job: Detecting overdue reservations and canceling them")
	defer func() { rollBackOrCommit(err) }()

	var overdueRes []model.Reservation
	result := tx.Model(&model.Reservation{}).
		Where("status = ?", model.ReservationStatusPending).
		Where("reservation_date < NOW()").
		Find(&overdueRes)
	if result.Error != nil {
		return
	}

	for _, res := range overdueRes {
		_, err = bookcopy.CancelReservationCopy(tx, int64(res.ID))
		if err != nil {
			return
		}
	}
}
