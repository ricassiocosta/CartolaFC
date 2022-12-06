package usecase

import (
	"context"
	"strconv"
	"strings"

	"github.com/ricassiocosta/consolidation/internal/domain/entity"
	"github.com/ricassiocosta/consolidation/internal/domain/repository"
	"github.com/ricassiocosta/consolidation/pkg/uow"
)

type UpdateMatchResultInput struct {
	ID     string
	Result string
}

type UpdateMatchResultUseCase struct {
	Uow uow.UowInterface
}

func NewMatchUpdateResultUseCase(uow uow.UowInterface) *UpdateMatchResultUseCase {
	return &UpdateMatchResultUseCase{
		Uow: uow,
	}
}

func (u *UpdateMatchResultUseCase) Execute(ctx context.Context, input UpdateMatchResultInput) error {
	err := u.Uow.Do(ctx, func(_ *uow.Uow) error {
		matchRepo := u.getMatchRepository(ctx)
		match, err := matchRepo.FindByID(ctx, input.ID)
		if err != nil {
			return err
		}
		matchResult := strings.Split(input.Result, "-")
		// convert results to int
		teamAResult, _ := strconv.Atoi(matchResult[0])
		teamBResult, _ := strconv.Atoi(matchResult[1])
		match.Result = *entity.NewMatchResult(teamAResult, teamBResult)

		err = matchRepo.Update(ctx, match)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (u *UpdateMatchResultUseCase) getMatchRepository(ctx context.Context) repository.MatchRepositoryInterface {
	matchRepository, err := u.Uow.GetRepository(ctx, "MatchRepository")
	if err != nil {
		panic(err)
	}
	return matchRepository.(repository.MatchRepositoryInterface)
}
