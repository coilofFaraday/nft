package model

type NFT struct {
	Name   string `json:"name"`
	Unique string `json:"unique"`
	Owner  string `json:"owner"`
}
