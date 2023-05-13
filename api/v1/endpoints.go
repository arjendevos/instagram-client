package v1

import (
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/arjendevos/instagram-client/models/dm"
	"github.com/arjendevos/instagram-client/models/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var limit = 11

type Endpoints struct {
	ctx context.Context
	db  *sql.DB
}

func InitEndpoints(ctx context.Context, db *sql.DB) *Endpoints {
	return &Endpoints{ctx: ctx, db: db}
}

func Profile(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("profile.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var payload responses.Profile
	err = json.Unmarshal(data, &payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}

func (e *Endpoints) GetProfiles(c *fiber.Ctx) error {
	cursor := c.Query("crsr")

	mods, err := CursorPagination(cursor, "instagram_id")
	if err != nil {
		return HandleErr(c, err, 400)
	}
	mods = append(mods, qm.Limit(limit))
	accounts, err := dm.InstagramAccounts(mods...).All(e.ctx, e.db)
	if err != nil {
		return HandleErr(c, err, 404)
	}

	p := ""
	n := ""

	if len(cursor) > 0 {
		p = "p_" + accounts[0].InstagramID
	}

	if len(accounts) == limit {
		n = "n_" + accounts[limit-1].InstagramID

	}

	return c.JSON(&fiber.Map{
		"success":    true,
		"accounts":   accounts[:limit-1],
		"pagination": PaginationResponse(p, n),
	})
}

// func postHandler(c *fiber.Ctx, db *sql.DB) error {
// 	return c.SendString("Hello")
// }

// func putHandler(c *fiber.Ctx, db *sql.DB) error {
// 	return c.SendString("Hello")
// }

// func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
// 	return c.SendString("Hello")
// }
