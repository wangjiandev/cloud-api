package service_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/wangjiandev/cloud-api/app/host"
	"github.com/wangjiandev/cloud-api/app/host/service"
)

var (
	hs host.Service
)

func TestCreate(t *testing.T) {
	ins := host.NewHost()
	ins.Name = "test"
	hs.CreateHost(context.Background(), ins)
}

func init() {
	//需要单独初始化logger
	zap.DevelopmentSetup()
	hs = service.NewHostService()
}
