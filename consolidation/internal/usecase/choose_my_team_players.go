package usecase

import (
	"context"

	"github.com/ricassiocosta/consolidation/internal/domain/repository"
	"github.com/ricassiocosta/consolidation/internal/domain/service"
	"github.com/ricassiocosta/consolidation/pkg/uow"
)

type ChooseMyTeamPlayersInput struct {
	ID        string
	PlayersID []string
}

type ChooseMyTeamPlayersUseCase struct {
	Uow uow.UowInterface
}

func NewMyTeamChoosePlayersUseCase(uow uow.UowInterface) *ChooseMyTeamPlayersUseCase {
	return &ChooseMyTeamPlayersUseCase{
		Uow: uow,
	}
}

func (u *ChooseMyTeamPlayersUseCase) Execute(ctx context.Context, input ChooseMyTeamPlayersInput) error {
	err := u.Uow.Do(ctx, func(_ *uow.Uow) error {
		myTeamRepo := u.getMyTeamRepository(ctx)
		myTeam, err := myTeamRepo.FindByID(ctx, input.ID)
		if err != nil {
			return err
		}
		playerRepo := u.getPlayerRepository(ctx)
		players, err := playerRepo.FindAllByIDs(ctx, input.PlayersID)
		if err != nil {
			return err
		}
		service.ChoosePlayers(myTeam, players)
		err = myTeamRepo.SavePlayers(ctx, myTeam)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (u *ChooseMyTeamPlayersUseCase) getMyTeamRepository(ctx context.Context) repository.MyTeamRepositoryInterface {
	myTeamRepository, err := u.Uow.GetRepository(ctx, "MyTeamRepository")
	if err != nil {
		panic(err)
	}
	return myTeamRepository.(repository.MyTeamRepositoryInterface)
}

func (u *ChooseMyTeamPlayersUseCase) getPlayerRepository(ctx context.Context) repository.PlayerRepositoryInterface {
	playerRepository, err := u.Uow.GetRepository(ctx, "PlayerRepository")
	if err != nil {
		panic(err)
	}
	return playerRepository.(repository.PlayerRepositoryInterface)
}
