package conversion

import (
	"github.com/MortalSC/FastGO/internal/apiserver/model"
	"github.com/MortalSC/FastGO/internal/commonpkg/core"
	apiv1 "github.com/MortalSC/FastGO/pkg/api/apiserver/v1"
)

func UserModelToUserV1(userModel *model.User) *apiv1.User {
	var protoUser apiv1.User
	_ = core.CopyWithConverters(&protoUser, userModel)
	return &protoUser
}

func UserV1ToUserModel(protoUser *apiv1.User) *model.User {
	var userModel model.User
	_ = core.CopyWithConverters(&userModel, protoUser)
	return &userModel
}
