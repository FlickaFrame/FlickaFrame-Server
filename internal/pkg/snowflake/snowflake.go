package snowflake

import "github.com/bwmarrin/snowflake"

const (
	userIDBits = iota
	videoIDBits
	commentIDBits
	tagIDBits
	favoriteIDBits
)

type IdGenerator interface {
	Generate() int64
}

var (
	UserIDNode, _     = snowflake.NewNode(userIDBits)
	VideoIDNode, _    = snowflake.NewNode(videoIDBits)
	CommentIDNode, _  = snowflake.NewNode(commentIDBits)
	TagIDNode, _      = snowflake.NewNode(tagIDBits)
	FavoriteIDNode, _ = snowflake.NewNode(favoriteIDBits)
)
