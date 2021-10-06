package rabbit

type Command interface {
	GetId() string
}

type CommandModel struct {
	Id string
}

func (c CommandModel) GetId() string {
	return c.Id
}

type CreateSensor struct {
	Name       string
	Sensortype string
	CommandModel
}
