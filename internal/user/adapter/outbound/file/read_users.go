package file

import (
	"context"
	"encoding/json"
	"live-coding/internal/user/core/domain"
	"os"
)

func (r Service) ReadUsers(ctx context.Context) ([]domain.User, error) {
	data, err := os.ReadFile(r.FilePath)
	if err != nil {
		return nil, err
	}

	var users []domain.User
	if err = json.Unmarshal(data, &users); err != nil {
		return nil, err
	}

	return users, nil
}
