package validation

import (
	"context"

	v1 "github.com/MortalSC/FastGO/pkg/api/apiserver/v1"
)

func (v *Validator) ValidateCreateUserRequest(ctx context.Context, req *v1.CreateUserRequest) error {
	return nil
}

func (v *Validator) ValidateUpdateUserRequest(ctx context.Context, req *v1.UpdateUserRequest) error {
	return nil
}

func (v *Validator) ValidateDeleteUserRequest(ctx context.Context, req *v1.DeleteUserRequest) error {
	return nil
}

func (v *Validator) ValidateGetUserRequest(ctx context.Context, req *v1.GetUserRequest) error {
	return nil
}

func (v *Validator) ValidateListUserRequest(ctx context.Context, req *v1.ListUserRequest) error {
	return nil
}

// ======= login with token ========
func (v *Validator) ValidateLoginRequest(ctx context.Context, req *v1.LoginRequest) error {
	return nil
}

func (v *Validator) ValidateRefreshTokenRequest(ctx context.Context, req *v1.RefreshTokenRequest) error {
	return nil
}

func (v *Validator) ValidateChangePasswordRequest(ctx context.Context, req *v1.ChangePasswordRequest) error {
	return nil
}
