package redis

import "errors"

var (
	ErrorVoteTimeExpire = errors.New("已过投票时间")
	ErrVoteRepested     = errors.New("已经投过票了")
)
