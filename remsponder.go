package remsponder

import (
	"reflect"
	"unicode"

	"github.com/yayuyokitano/kitaipu"
)

type Remsponder struct{}

func CallInteraction(interaction kitaipu.Command) (resp kitaipu.InteractionResponse, err error) {

	r := Remsponder{}
	lower := []rune(interaction.Data.Name)
	lower[0] = unicode.ToUpper(lower[0])
	name := string(lower)
	method := reflect.ValueOf(r).MethodByName(name)
	param := []reflect.Value{reflect.ValueOf(interaction)}
	res := method.Call(param)

	err, ok := res[1].Interface().(error)
	if ok { // ok should be false only if err is not actually an error
		return
	}
	resp = res[0].Interface().(kitaipu.InteractionResponse)

	return
}
