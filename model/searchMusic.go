package model

type SearchMusic struct {
	Music []Music `json:"Music"`
	Key   string  `json:"Key"`
}
