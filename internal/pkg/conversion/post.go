package conversion

import (
	"github.com/MortalSC/FastGO/internal/apiserver/model"
	apiv1 "github.com/MortalSC/FastGO/pkg/api/apiserver/v1"
	"github.com/jinzhu/copier"
)

func PostModelToPostV1(postModel *model.Post) *apiv1.Post {
	var protoPost apiv1.Post
	_ = copier.Copy(&protoPost, postModel)
	return &protoPost
}

func PostV1ToPostModel(protoPost *apiv1.Post) *model.Post {
	var postModel model.Post
	_ = copier.Copy(&postModel, protoPost)
	return &postModel
}
