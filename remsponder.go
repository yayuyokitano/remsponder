package remsponder

import (
	"reflect"

	"github.com/yayuyokitano/kitaipu"
)

type Remsponder struct{}

func CallInteraction(interaction kitaipu.Command) (resp kitaipu.InteractionResponse, err error) {

	r := Remsponder{}
	method := reflect.ValueOf(r).MethodByName(interaction.Data.Name)
	param := []reflect.Value{reflect.ValueOf(interaction)}
	res := method.Call(param)

	err = res[1].Interface().(error)
	if err != nil {
		return
	}
	resp = res[0].Interface().(kitaipu.InteractionResponse)

	return
}
