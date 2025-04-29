package file

import (
	"context"
	"encoding/json"
	"live-coding/internal/user/adapter/outbound/file/dto"
	"live-coding/internal/user/core/domain"
	"live-coding/pkg/slice"
	"os"
)

func (r Service) ReadUsers(ctx context.Context) ([]domain.User, error) {
	data, err := os.ReadFile(r.FilePath)
	if err != nil {
		return nil, err
	}

	var users []dto.User
	if err = json.Unmarshal(data, &users); err != nil {
		return nil, err
	}

	return slice.Convert(users, dto.UserToDomain), nil
}
