package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/ricassiocosta/consolidation/internal/infra/db"
	handlers "github.com/ricassiocosta/consolidation/internal/infra/http"
	"github.com/ricassiocosta/consolidation/internal/infra/repository"
	"github.com/ricassiocosta/consolidation/pkg/uow"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dtb, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/cartola?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer dtb.Close()

	uow, err := uow.NewUow(ctx, dtb)
	if err != nil {
		panic(err)
	}

	registerRepositories(uow)

	http.HandleFunc("/players", handlers.ListPlayersHandler(ctx, *db.New(dtb)))

	http.ListenAndServe(":8080", nil)

}

func registerRepositories(uow *uow.Uow) {
	uow.Register("PlayerRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewPlayerRepository(uow.Db)
		repo.Queries = db.New(tx)

		return repo
	})

	uow.Register("MatchRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewMatchRepository(uow.Db)
		repo.Queries = db.New(tx)

		return repo
	})

	uow.Register("PlayerRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewPlayerRepository(uow.Db)
		repo.Queries = db.New(tx)

		return repo
	})

	uow.Register("MyTeamRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewMyTeamRepository(uow.Db)
		repo.Queries = db.New(tx)

		return repo
	})
}
