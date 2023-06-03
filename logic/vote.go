package logic

import (
	"strconv"
	"zouyi/bluebell/dao/redis"
	"zouyi/bluebell/model"

	"go.uber.org/zap"
)

// redis score
// vote + 432

// Voting restriction: Posts exceeding one week are not allowed to vote.
// After expiration,save the data in redis to mysql.
// Delete data from mysql when it expires.

func PostVote(userId uint64, vf *model.VoteForm) error {
	zap.L().Debug("logic.post vote",
		zap.Uint64("userId", userId),
		zap.String("postId", vf.PostID),
		zap.Int8("attitude", vf.Attitude),
	)
	return redis.VoteForPost(strconv.Itoa(int(userId)), vf.PostID, float64(vf.Attitude))

}
