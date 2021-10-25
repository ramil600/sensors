package rabbit

import (
	"encoding/json"
	"log"
	"reflect"
)

type Serializer struct {
	eventtypes map[string]reflect.Type
}

func (s Serializer) UnmarshalEvent(data []byte) Command {

	wrapper := WrapperEvent{}

	err := json.Unmarshal(data, &wrapper)

	if err != nil {
		log.Println(err)
	}
	log.Println(wrapper.Type)

	evttype, ok := s.eventtypes[wrapper.Type]
	if !ok {
		log.Fatal("EventType Not Found ")
	}

	log.Println("Event type retrived from Serializer is: ", evttype.Name())

	cmdreceived := reflect.New(evttype).Interface()

	err = json.Unmarshal(wrapper.Data, cmdreceived)

	if err != nil {
		log.Println(err)
	}

	log.Println("Received command name", cmdreceived)

	return cmdreceived.(Command)

}

func NewSerializer(evttypes ...Command) *Serializer {

	eventtypes := map[string]reflect.Type{}

	for _, evt := range evttypes {
		t := reflect.TypeOf(evt)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		eventtypes[t.Name()] = t
	}
	return &Serializer{
		eventtypes: eventtypes,
	}

}
