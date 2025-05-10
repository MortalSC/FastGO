package validation

import (
	"context"

	v1 "github.com/MortalSC/FastGO/pkg/api/apiserver/v1"
)

func (v *Validator) ValidateCreatePostRequest(ctx context.Context, req *v1.CreatePostRequest) error {
	return nil
}

func (v *Validator) ValidateUpdatePostRequest(ctx context.Context, req *v1.UpdatePostRequest) error {
	return nil
}

func (v *Validator) ValidateDeletePostRequest(ctx context.Context, req *v1.DeletePostRequest) error {
	return nil
}

func (v *Validator) ValidateGetPostRequest(ctx context.Context, req *v1.GetPostRequest) error {
	return nil
}

func (v *Validator) ValidateListPostRequest(ctx context.Context, req *v1.ListPostRequest) error {
	return nil
}
