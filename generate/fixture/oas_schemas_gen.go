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

type OptBool struct {
	Value bool
	Set   bool
}

type OptNilString struct {
	Value string
	Set   bool
	Null  bool
}

type OptNilInt struct {
	Value int
	Set   bool
	Null  bool
}

type OptNilInt64 struct {
	Value int64
	Set   bool
	Null  bool
}

type OptNilFloat64 struct {
	Value float64
	Set   bool
	Null  bool
}

type OptNilBool struct {
	Value bool
	Set   bool
	Null  bool
}
