package model

type RequiredParams struct {
	Body map[string]any `form:"body,required" json:"body,required"`
}
