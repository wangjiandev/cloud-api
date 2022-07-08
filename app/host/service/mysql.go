package service

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/wangjiandev/cloud-api/app/host"
)

// 接口实现的静态检查
var _ host.Service = (*HostService)(nil)

func NewHostService() *HostService {
	return &HostService{
		// 初始化logger服务,标识以下是host服务的子logger
		l: zap.L().Named("HOST"),
	}
}

type HostService struct {
	l logger.Logger
}
