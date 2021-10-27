package rabbit

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerializer_UnmarshalEvent(t *testing.T) {

	createcomm := CreateSensor{
		Name: "living area",
		CommandModel: CommandModel{
			Id: "se2e-eia2",
		},
	}
	ser := NewSerializer(createcomm)

	// Generate wrapped CreateCommand
	data, _ := json.Marshal(createcomm)
	wrp := WrapperEvent{
		Type: "CreateSensor",
		Data: data,
	}
	wrpdata, err := json.Marshal(wrp)
	assert.NoError(t, err)

	com := ser.UnmarshalEvent(wrpdata)
	actual, ok := com.(*CreateSensor)

	assert.True(t, ok)
	assert.Equal(t, createcomm.Name, actual.Name)
	assert.Equal(t, createcomm.Id, actual.Id)

}
