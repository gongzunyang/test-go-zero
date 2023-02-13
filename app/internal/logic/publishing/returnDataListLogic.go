package publishing

import (
	"context"

	"gov-cms/app/internal/svc"
	"gov-cms/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnDataListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnDataListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnDataListLogic {
	return &ReturnDataListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnDataListLogic) ReturnDataList(req *types.SearchReq) (*types.SearchReply, error) {
	// todo: add your logic here and delete this line
	return &types.SearchReply{
		Name:  "gzy",
		Count: 123,
	}, nil
}
