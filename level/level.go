package remsponderlevel

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/yayuyokitano/kitaipu"
)

func init() {
	functions.HTTP("level", level)
}

func level(writer http.ResponseWriter, request *http.Request) {

	var command kitaipu.Command
	if err := json.NewDecoder(request.Body).Decode(&command); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Println(`{"message": "Failed to decode request body", "severity": "error"}`)
		log.Println(err)
		return
	}

	contentType, res, err := kitaipu.InteractionResponse{
		Type: kitaipu.CallbackUpdateMessage,
		Data: kitaipu.InteractionResponseMessage{
			Content: "Hallo!",
		},
	}.Prepare()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Println(`{"message": "Failed to prepare response", "severity": "error"}`)
		log.Println(err)
		return
	}

	writer.Header().Set("Content-Type", contentType)
	writer.WriteHeader(http.StatusOK)
	fmt.Fprint(writer, res)

}
