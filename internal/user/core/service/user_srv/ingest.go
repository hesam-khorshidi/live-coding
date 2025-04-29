package user_srv

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"live-coding/internal/user/core/domain"
	"sync"
)

func (s Service) Ingest(ctx context.Context) error {
	users, err := s.userFileSrv.ReadUsers(ctx)
	if err != nil {
		return fmt.Errorf("failed to read users: %w", err)
	}

	jobs := make(chan domain.User, len(users))
	errs := make(chan error, len(users))
	var wg sync.WaitGroup

	for i := 0; i < s.workerCount; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for user := range jobs {
				for i := range user.Addresses {
					user.Addresses[i].ID = uuid.New()
				}
				if err := s.userRepo.Save(ctx, user); err != nil {
					errs <- fmt.Errorf("worker %d: %w", workerID, err)
				}
			}
		}(i + 1)
	}

	for _, user := range users {
		jobs <- user
	}
	close(jobs)

	wg.Wait()
	close(errs)

	var allErrors []error
	for err := range errs {
		allErrors = append(allErrors, err)
	}

	if len(allErrors) > 0 {
		return fmt.Errorf("processing finished with %d error(s): %w", len(allErrors), errors.Join(allErrors...))
	}

	return nil
}
