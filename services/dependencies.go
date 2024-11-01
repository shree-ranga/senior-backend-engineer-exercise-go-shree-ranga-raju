package services

import (
	appcontext "syndio/appContext"
	"syndio/repo"
)

type ServerDependencies struct {
	SyndioService SyndioService
}

func InstantiateServerDependencies() *ServerDependencies {
	db := appcontext.GetDB()
	syndioRepo := repo.NewSyndioRepo(db)
	syndioService := NewSyndioService(syndioRepo)

	return &ServerDependencies{
		SyndioService: syndioService,
	}
}
