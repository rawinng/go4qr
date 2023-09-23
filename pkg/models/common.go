package models

// declare interface
type CommonJson interface {
	ToJson() string
	fromJson(string)
}
