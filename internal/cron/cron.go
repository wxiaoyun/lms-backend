package cron

import (
	finejob "lms-backend/internal/cron/fine"
	reservationjob "lms-backend/internal/cron/reservation"

	cronn "github.com/robfig/cron/v3"
)

func RunJobs() *cronn.Cron {
	cr := cronn.New()

	_, err := cr.AddFunc("@every 1h", finejob.DetectOverdueLoansAndCreateFine)
	if err != nil {
		panic(err)
	}

	_, err = cr.AddFunc("@every 12h", reservationjob.DetectOverdueResAndCancelRes)
	if err != nil {
		panic(err)
	}

	cr.Start()
	return cr
}
