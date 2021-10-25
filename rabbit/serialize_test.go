package rabbit

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerializer_UnmarshalEvent(t *testing.T) {

	ser := NewSerializer(CreateSensor{}, UpdateSensor{})
	createcomm := CreateSensor{
		Name: "living area",
		CommandModel: CommandModel{
			Id: "se2e-eia2",
		},
	}
	data, _ := json.Marshal(createcomm)

	wrp := WrapperEvent{
		Type: "CreateSensor",
		Data: data,
	}

	wrpdata, _ := json.Marshal(wrp)

	com := ser.UnmarshalEvent(wrpdata)
	actual, _ := com.(*CreateSensor)

	assert.Equal(t, createcomm.Name, actual.Name)
	assert.Equal(t, createcomm.Id, actual.Id)

}
