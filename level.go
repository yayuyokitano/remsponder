package remsponder

import (
	"context"
	"fmt"
	"math"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/yayuyokitano/kitaipu"
)

var pool *pgxpool.Pool

func (r Remsponder) Level(interaction kitaipu.Command) (resp kitaipu.InteractionResponse, err error) {

	pool, err = pgxpool.Connect(context.Background(), os.Getenv("DATABASE_PRIVATE_URL"))
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	switch interaction.Data.Options[0].Name {
	case "display":
		resp, err = levelDisplay(interaction, interaction.Data.Options[0].Options)
		break
	}

	return

}

func levelDisplay(interaction kitaipu.Command, opts kitaipu.Options) (resp kitaipu.InteractionResponse, err error) {

	var userID string
	user, ok := opts.Get("user")
	if ok {
		userID = user.Value
	} else {
		userID = interaction.Member.User.ID
	}

	//select the xp number from the database
	var xp int64
	err = pool.QueryRow(context.Background(), "SELECT xp FROM guildXP WHERE guildID = $1 AND userID = $2", interaction.GuildID, userID).Scan(&xp)
	if err != nil {
		xp = 0
		err = nil
	}

	level := getLevel(xp)
	xpPrev := getRequiredXP(level)
	xpCur := xp - xpPrev
	xpNext := getRequiredXP(level+1) - xpPrev
	xpLeft := xpNext - xpCur
	xpPct := int(100 * float64(xpCur) / float64(xpNext))

	resp = kitaipu.InteractionResponse{
		Content: fmt.Sprintf("%s has %d xp, level %d, %d%% until level %d (%dxp remaining)", interaction.Data.Resolved.Users[userID].Username, xp, level, xpPct, level+1, xpLeft),
	}

	return
}

func getRequiredXP(level int64) int64 {
	return int64(((5.0/3)*math.Pow(float64(level-1), 3))+(27.5*math.Pow(float64(level-1), 2))+((755.0/6)*float64(level-1))) + 100
}

// Im sorry
func getLevel(xp int64) int64 {
	a := 5.0 / 3
	d := -float64(xp) + 99 //float imprecision stuff
	g := (d / a) - 82.5
	h := math.Pow(g, 2)/4 - 131.35464537
	sqrth := math.Sqrt(h)
	R := sqrth - g/2
	S := math.Pow(R, 1.0/3)
	T := -(g / 2) - sqrth
	U := math.Pow(T, 1.0/3)
	L := U + S - 5.5
	return int64(L + 1)
}
