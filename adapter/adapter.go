package adapter

type ICreateServer interface {
	CreateServer(cpu,mem float64)
}

type AWSClient struct {}

func (client *AWSClient) RunInstance(cpu,mem float64) {
	panic("implement me")
}

type AWSAdapter struct {
	client AWSClient
}

func (adapter AWSAdapter) CreateServer(cpu,mem float64) {
	adapter.client.RunInstance(cpu,mem)
}

type ALiYunClient struct {}

func (client *ALiYunClient) StartInstance(cpu,mem float64) {
	panic("implement me")
}

type ALiYunAdapter struct {
	client ALiYunClient
}

func (adapter ALiYunAdapter) CreateServer(cpu,mem float64) {
	adapter.client.StartInstance(cpu,mem)
}


