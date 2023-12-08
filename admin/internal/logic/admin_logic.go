package logic

import (
	"context"
	"fmt"
	"iot-platform/helper"
	"iot-platform/models"

	"iot-platform/admin/internal/svc"
	"iot-platform/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLogic {
	return &AdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminLogic) Admin(req *types.DeviceListRequest) (resp *types.DeviceListReply, err error) {
	// todo: add your logic here and delete this line
	req.Size = helper.If(req.Size == 0, 20, req.Size).(int)
	req.Page = helper.If(req.Size == 0, 0, (req.Page-1)*req.Size).(int)
	var count int64

	resp = new(types.DeviceListReply)
	data := make([]*types.DeviceListBasic, 0)
	err = models.GetDeviceList(req.Name).Count(&count).Limit(req.Size).Offset(req.Page).Find(&data).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)

		return
	}
	resp.Count = count
	resp.List = data
	fmt.Println(resp)
	return

}
