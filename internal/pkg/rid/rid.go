package rid

import "github.com/MortalSC/FastGO/internal/commonpkg/id"

const defaultABC = "abcdefghijklmnopqrstuvwxyz1234567890"

type ResourceID string

const (
	UserID ResourceID = "userID"
	PostID ResourceID = "postID"
)

func (rid ResourceID) String() string {
	return string(rid)
}

func (rid ResourceID) New(counter uint64) string {
	uniqueStr := id.NewCode(
		counter,
		id.WithCodeChars([]rune(defaultABC)),
		id.WithCodeL(6),
		id.WithCodeSalt(Salt()),
	)
	return rid.String() + "-" + uniqueStr
}
