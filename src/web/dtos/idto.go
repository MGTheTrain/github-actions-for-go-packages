package dtos

type IDto interface {
	Validate() []string
}
