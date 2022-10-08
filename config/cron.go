package config

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

type Days struct {
	Day      string
	Friday   time.Time
	Saturday time.Time
	Sunday   time.Time
}

func (server *Server) runCron(c **cron.Cron, config *API) {
	// ctx := context.Background()
	(*c) = cron.New()

	(*c).Start()
	log.Printf("%+v\n", (*c).Entries())
}
