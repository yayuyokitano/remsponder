package remsponder

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/yayuyokitano/kitaipu"
)

type VerifiedInteraction struct {
	Interaction kitaipu.Command `json:"interaction"`
	Token       string          `json:"token"`
}
type PubSubMessage struct {
	Data []byte `json:"data"`
}

var pool *pgxpool.Pool

func createPool() (err error) {
	if pool == nil {
		ctx := context.Background()

		pool, err = pgxpool.Connect(ctx, os.Getenv("DATABASE_PRIVATE_URL"))
	}
	return
}

func Respond(ctx context.Context, m PubSubMessage) {

	var verifiedInteraction VerifiedInteraction
	if err := json.Unmarshal(m.Data, &verifiedInteraction); err != nil {
		fmt.Print("Failed to unmarshal interaction", err)
		return
	}

	if verifiedInteraction.Token != os.Getenv("DISCORD_SECRET") {
		fmt.Print("Invalid request token")
		return
	}
	interaction := verifiedInteraction.Interaction

	err := createPool()
	if err != nil {
		fmt.Print("Failed to create pool", err)
		return
	}

	res, err := callInteraction(interaction)
	if err != nil {
		fmt.Print(err)
		return
	}
	contentType, b, err := res.Prepare()
	fmt.Println(string(b))

	client := http.Client{}
	url := fmt.Sprintf("%s/webhooks/%s/%s/messages/@original", os.Getenv("DISCORD_BASE_URI"), interaction.ApplicationID, interaction.Token)
	req, err := http.NewRequest("PATCH", url, bytes.NewReader(b))
	if err != nil {
		fmt.Print(err)
		return
	}
	req.Header.Set("Content-Type", contentType)
	resp, err := client.Do(req)
	rawBody, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(rawBody))
	if err != nil {
		fmt.Print(err)
	}
	return

}
