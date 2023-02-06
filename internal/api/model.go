package api

import "fmt"

type MissingQueryError struct {
	Parameter string
}

func (error MissingQueryError) Error() string {
	return fmt.Sprintf("Query parameter %s is missing", error.Parameter)
}
