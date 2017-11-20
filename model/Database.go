package model

import "github.com/aerogo/nano"

// DB is the database.
var DB = nano.New(5000).Namespace("kcc").RegisterTypes(
	(*Page)(nil),
	(*Story)(nil),
)
