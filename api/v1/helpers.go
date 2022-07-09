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

func PaginationResponse(hasPrev bool, hasNext bool, p string, n string) *fiber.Map {
	pc := ""
	nc := ""

	if hasPrev {
		pc =
			base64.StdEncoding.EncodeToString([]byte("p_" + p))
	}

	if hasNext {
		nc = base64.StdEncoding.EncodeToString([]byte("n_" + n))

	}

	return &fiber.Map{
		"previous_cursor": pc,
		"next_cursor":     nc,
	}
}
