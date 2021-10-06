package rabbit

import "context"


type Command interface {
	GetId() string
}

type CommandModel struct {
	Id string
}
func (c CommandModel) GetId() string {
	return c.Id
}
type Dispatcher interface {
	Apply( context.Context, Command) error
}

type CreateSensor struct {
	Name string
	Sensortype string
	CommandModel
}