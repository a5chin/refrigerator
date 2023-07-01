package entity

type Ingredient struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Weight uint   `json:"weight"`
}
