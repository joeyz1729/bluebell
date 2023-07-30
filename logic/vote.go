package logic

import (
	"strconv"

	"github.com/YiZou89/bluebell/dao/redis"
	"github.com/YiZou89/bluebell/model"

	"go.uber.org/zap"
)

func PostVote(userId uint64, vf *model.VoteForm) error {
	zap.L().Debug("logic.post vote",
		zap.Uint64("userId", userId),
		zap.String("postId", vf.PostID),
		zap.Int8("attitude", vf.Attitude),
	)
	return redis.VoteForPost(strconv.Itoa(int(userId)), vf.PostID, float64(vf.Attitude))

}
