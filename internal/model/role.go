package model

type RoleAddInput struct {
	Name string
	Desc string
}

type RoleAddOutput struct {
	RoleId uint
}

type RoleDeleteInput struct {
	Id int
}

type RoleDeleteOutput struct {
	Id int
}
