package publishing

import (
	"context"

	"gov-cms/app/internal/svc"
	"gov-cms/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishingLogic {
	return &PublishingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishingLogic) Publishing(req *types.PublishingListReq) (resp *types.PublishingListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
