package remsponder

import (
	"github.com/yayuyokitano/kitaipu"
)

func (r Remsponder) level(interaction kitaipu.Command) (contentType string, b []byte, err error) {

	contentType, b, err = kitaipu.InteractionResponse{
		Type: kitaipu.CallbackUpdateMessage,
		Data: kitaipu.InteractionResponseMessage{
			Content: "Hallo!",
		},
	}.Prepare()

	return

}
