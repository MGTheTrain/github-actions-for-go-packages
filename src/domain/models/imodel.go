package models

type IModel interface {
	Validate() []string
}
