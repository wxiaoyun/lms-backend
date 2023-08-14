package cron

import (
	"lms-backend/internal/cron/finejob"

	cronn "github.com/robfig/cron/v3"
)

func RunJobs() *cronn.Cron {
	cr := cronn.New()

	_, err := cr.AddFunc("@every 1h", finejob.DetectOverdueLoansAndCreateFine)
	if err != nil {
		panic(err)
	}

	cr.Start()
	return cr
}
