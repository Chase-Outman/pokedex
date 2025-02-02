package main

type Location struct {
	Count   int                 `json:"count"`
	Next    string              `json:"next"`
	Prev    string              `json:"previous"`
	Results []map[string]string `json:"results"`
}
