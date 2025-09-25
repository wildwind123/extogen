package ogenapi

type OptString struct {
	Value string
	Set   bool
}

type OptInt struct {
	Value int
	Set   bool
}

type OptFloat64 struct {
	Value float64
	Set   bool
}

type OptNilInt struct {
	Value int
	Set   bool
	Null  bool
}
