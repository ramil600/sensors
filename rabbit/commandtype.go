package rabbit

type Command interface {
	GetId() string
}

type CommandModel struct {
	Id string
	Type string
}

func (c CommandModel) GetId() string {
	return c.Id
}

type CreateSensor struct {
	Name       string
	Sensortype string
	CommandModel
}

type UpdateSensor struct {
	CommandModel
	Name string
	Sensortype string
}
