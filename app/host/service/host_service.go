package service

import (
	"context"

	"github.com/infraboard/mcube/logger"
	"github.com/wangjiandev/cloud-api/app/host"
)

func (s *HostService) CreateHost(ctx context.Context, host *host.Host) (*host.Host, error) {
	s.l.Debug("create host")
	s.l.Debugf("create host %s", host.Name)
	// 携带额外meta数据,常用于TRACE系统,快速搜索request-01打印的日志
	s.l.With(logger.NewAny("request-id", "request-01")).Debug("create host")
	return nil, nil
}

func (s *HostService) QueryHost(ctx context.Context, host *host.QueryHostRequest) (*host.HostSet, error) {
	return nil, nil
}

func (s *HostService) DescribeHost(ctx context.Context, host *host.DescribeHostRequest) (*host.Host, error) {
	return nil, nil
}

func (s *HostService) UpdateHost(ctx context.Context, host *host.UpdateHostRequest) (*host.Host, error) {
	return nil, nil
}

func (s *HostService) DeleteHost(ctx context.Context, host *host.DeleteHostRequest) (*host.Host, error) {
	return nil, nil
}
