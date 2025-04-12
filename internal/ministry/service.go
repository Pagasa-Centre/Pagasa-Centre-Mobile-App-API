package ministry

import (
	"context"
	"fmt"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/communication"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry/storage"
)

type MinistryService interface {
	All(ctx context.Context) ([]*domain.Ministry, error)
	AssignLeaderToMinistry(ctx context.Context, ministryID string, userID string) error
}

type service struct {
	logger               *zap.Logger
	ministryRepo         storage.MinistryRepository
	communicationService communication.CommunicationService
}

func NewMinistryService(
	logger *zap.Logger,
	ministryRepo storage.MinistryRepository,
	communicationService communication.CommunicationService,
) MinistryService {
	return &service{
		logger:               logger,
		ministryRepo:         ministryRepo,
		communicationService: communicationService,
	}
}

func (ms *service) AssignLeaderToMinistry(ctx context.Context, ministryID string, userID string) error {
	err := ms.ministryRepo.AssignLeaderToMinistry(ctx, ministryID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (ms *service) All(ctx context.Context) ([]*domain.Ministry, error) {
	ministryEntities, err := ms.ministryRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all ministries: %s", err)
	}

	var ministries []*domain.Ministry
	for _, entity := range ministryEntities {
		ministries = append(ministries, domain.ToDomain(entity))
	}

	return ministries, nil
}
