package remsponder

import (
	"github.com/yayuyokitano/kitaipu"
)

func (r Remsponder) Level(interaction kitaipu.Command) (resp kitaipu.InteractionResponse, err error) {

	resp = kitaipu.InteractionResponse{
		Type: kitaipu.CallbackChannelMessageWithSource,
		Data: kitaipu.InteractionResponseMessage{
			Content: "Hallo!",
		},
	}

	return

}
