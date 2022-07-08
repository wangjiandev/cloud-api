package host

import "context"

type Service interface {
	CreateHost(context.Context, *Host) (*Host, error)
	QueryHost(context.Context, *QueryHostRequest) (*HostSet, error)
	DescribeHost(context.Context, *DescribeHostRequest) (*Host, error)
	UpdateHost(context.Context, *UpdateHostRequest) (*Host, error)
	DeleteHost(context.Context, *DeleteHostRequest) (*Host, error)
}

type Vendor int

const (
	PRIVATE_IDC Vendor = iota
	ALIYUN
	TXYUN
)

func NewHost() *Host {
	return &Host{
		Resource: &Resource{},
		Describe: &Describe{},
	}
}

type Host struct {
	*Resource
	*Describe
}

type Resource struct {
	Id          string `json:"id"`
	Vendor      Vendor `json:"vendor"`
	Region      string `json:"region"`
	CreateAt    int64  `json:"create_at"`
	ExpireAt    int64  `json:"expire_at"`
	Type        string `json:"type"` //规格
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Tags        string `json:"tags"`
	UpdateAt    int64  `json:"update_at"`
	SyncAt      int64  `json:"sync_at"`
	Account     string `json:"account"`
	PublicIP    string `json:"public_ip"`
	PrivateIP   string `json:"private_ip"`
	PayType     string `json:"pay_type"`
}

type Describe struct {
	CPU          string `json:"cpu"`
	Memory       string `json:"memory"`
	GPUAmount    string `json:"gpu_amount"`
	GPUType      string `json:"gpu_type"`
	OSType       string `json:"os_type"`
	OSName       string `json:"os_name"`
	SerialNumber string `json:"serial_number"`
}

type HostSet struct {
	Items []*Host
	Total int64
}

type QueryHostRequest struct {
}

type DescribeHostRequest struct {
}

type UpdateHostRequest struct {
	*Describe
}

type DeleteHostRequest struct {
	Id string
}
