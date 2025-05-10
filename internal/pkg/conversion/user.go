package conversion

import (
	"github.com/MortalSC/FastGO/internal/apiserver/model"
	apiv1 "github.com/MortalSC/FastGO/pkg/api/apiserver/v1"
	"github.com/jinzhu/copier"
)

func UserModelToUserV1(userModel *model.User) *apiv1.User {
	var protoUser apiv1.User
	_ = copier.Copy(&protoUser, userModel)
	return &protoUser
}

func UserV1ToUserModel(protoUser *apiv1.User) *model.User {
	var userModel model.User
	_ = copier.Copy(&userModel, protoUser)
	return &userModel
}
