package user

import (
	"./group"
	"./permission"
)


type User struct {
	id uint
	firstname string
	surname string
	lastname string
	groups []*group.Group
}

func (u User) hasPerm(perm permission.Permission) bool {
	for _, g := range u.groups {
		for _, p := range g.permissions {
			if p == perm {
				return true
			}
		}
	}
	return false
}
