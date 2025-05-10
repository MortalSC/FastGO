package conversion

import (
	"github.com/MortalSC/FastGO/internal/apiserver/model"
	"github.com/MortalSC/FastGO/internal/commonpkg/core"
	apiv1 "github.com/MortalSC/FastGO/pkg/api/apiserver/v1"
)

func PostModelToPostV1(postModel *model.Post) *apiv1.Post {
	var protoPost apiv1.Post
	_ = core.CopyWithConverters(&protoPost, postModel)
	return &protoPost
}

func PostV1ToPostModel(protoPost *apiv1.Post) *model.Post {
	var postModel model.Post
	_ = core.CopyWithConverters(&postModel, protoPost)
	return &postModel
}
