package adapter

type Adapter interface {
	Connect() DatabaseAdapter
}
