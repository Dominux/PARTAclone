package group

import "./permission"


type Group struct {
	name string
	permissions []*permission.Permission
}
