package response

import "gochat/model"

type SearchResponse struct {
	User  model.User  `json:"user"`
	Group model.Group `json:"group"`
}
