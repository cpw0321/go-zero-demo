package infos

import (
	"context"

	"cpw/go-zero-demo/internal/svc"
	"cpw/go-zero-demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListInfosLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListInfosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListInfosLogic {
	return &ListInfosLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListInfosLogic) ListInfos(req *types.ListInfosReq) (resp *types.ListBasicRes, err error) {
	// todo: add your logic here and delete this line

	return
}
