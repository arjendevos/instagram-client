package helpers

import "github.com/volatiletech/null/v8"

func ns(v string) null.String {
	return null.NewString(v, true)
}
