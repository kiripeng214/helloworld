package data

import (
	"github.com/google/wire"
	"helloworld/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewStudentRepo)

// Data .
type Data struct {
	// TODO warpped database client
}

// NewData .
func NewData(c *conf.Data) (*Data, error) {
	return &Data{}, nil
}
