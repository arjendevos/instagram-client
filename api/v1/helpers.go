package v1

import (
	"fmt"
	"strings"

	"encoding/base64"

	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func CursorPagination(cursor string, field string) ([]qm.QueryMod, error) {
	var orderBy = "ASC"
	var whereOperartor = ">="
	if len(cursor) <= 0 {
		return []qm.QueryMod{qm.OrderBy(fmt.Sprintf("%v %v", field, orderBy))}, nil
	}

	fieldValue, prefix, err := decodeCursor(cursor)
	if err != nil {
		return nil, err
	}

	if prefix == "p" {
		orderBy = "DESC"
		whereOperartor = "<"
	}

	return []qm.QueryMod{qm.OrderBy(fmt.Sprintf("%v %v", field, orderBy)), qm.Where(fmt.Sprintf("%v %v ?", field, whereOperartor), fieldValue)}, nil
}

func decodeCursor(cursor string) (fv string, p string, err error) {
	d, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return "", "", err
	}

	s := strings.Split(string(d), "_")

	return s[1], s[0], nil
}

func PaginationResponse(p string, n string) *fiber.Map {
	return &fiber.Map{
		"previous_cursor": base64.StdEncoding.EncodeToString([]byte(p)),
		"next_cursor":     base64.StdEncoding.EncodeToString([]byte(n)),
	}
}

func HandleErr(c *fiber.Ctx, err error, sc int) error {
	return c.Status(400).JSON(&fiber.Map{
		"success": false,
		"error":   err.Error(),
	})
}
