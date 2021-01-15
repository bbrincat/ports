package domain

// Domain interfaces and Structs that repositories  and service classes depend on

type PortSenderService interface {
	SendPort(Port) error
}

type PortLoaderService interface {
	LoadPorts(PortSenderService) error
}

type AllPort struct {
  code map[string]Port `json:"code"`
}

type Port struct {
	code string`json:"name"`
	name string`json:"city"`
	//TODO add rest of fields
}
