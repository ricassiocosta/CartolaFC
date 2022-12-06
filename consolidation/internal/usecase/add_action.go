package usecase

import (
	"context"

	"github.com/ricassiocosta/consolidation/internal/domain/entity"
	"github.com/ricassiocosta/consolidation/internal/domain/repository"
	"github.com/ricassiocosta/consolidation/pkg/uow"
)

type AddActionInput struct {
	MatchID  string
	TeamID   string
	PlayerID string
	Action   string
	Minute   int
}

type AddActionUseCase struct {
	Uow         uow.UowInterface
	ActionTable entity.ActionTable
}

func (a *AddActionUseCase) Execute(ctx context.Context, input AddActionInput) error {
	return a.Uow.Do(ctx, func(uow *uow.Uow) error {
		matchRepo := a.getMatchRepository(ctx)
		myTeamRepo := a.getMyTeamRepository(ctx)
		playerRepo := a.getPlayerRepository(ctx)

		match, err := matchRepo.FindByID(ctx, input.MatchID)
		if err != nil {
			return nil
		}

		score, err := a.ActionTable.GetScore(input.Action)
		if err != nil {
			return nil
		}

		theAction := entity.NewGameAction(input.PlayerID, input.Minute, input.Action, score)
		match.Actions = append(match.Actions, *theAction)

		err = matchRepo.SaveActions(ctx, match, float64(score))
		if err != nil {
			return nil
		}

		player, err := playerRepo.FindByID(ctx, input.PlayerID)
		if err != nil {
			return nil
		}

		player.Price += float64(score)
		err = playerRepo.Update(ctx, player)
		if err != nil {
			return nil
		}

		myTeam, err := myTeamRepo.FindByID(ctx, input.TeamID)
		if err != nil {
			return nil
		}

		err = myTeamRepo.AddScore(ctx, myTeam, float64(score))
		if err != nil {
			return nil
		}

		return nil
	})
}

func (a *AddActionUseCase) getMatchRepository(ctx context.Context) repository.MatchRepositoryInterface {
	matchRepository, err := a.Uow.GetRepository(ctx, "MatchRepository")
	if err != nil {
		panic(err)
	}

	return matchRepository.(repository.MatchRepositoryInterface)
}

func (a *AddActionUseCase) getMyTeamRepository(ctx context.Context) repository.MyTeamRepositoryInterface {
	myTeamRepository, err := a.Uow.GetRepository(ctx, "MyTeamRepository")
	if err != nil {
		panic(err)
	}

	return myTeamRepository.(repository.MyTeamRepositoryInterface)
}

func (a *AddActionUseCase) getPlayerRepository(ctx context.Context) repository.PlayerRepositoryInterface {
	playerRepository, err := a.Uow.GetRepository(ctx, "PlayerRepository")
	if err != nil {
		panic(err)
	}

	return playerRepository.(repository.PlayerRepositoryInterface)
}
