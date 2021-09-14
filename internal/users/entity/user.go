package entity


type User struct {
	Id uint
	Firstname string
	Surname string
	Lastname string
	Groups []*Group
}

func (u User) hasPerm(perm Permission) bool {
	for _, g := range u.Groups {
		for _, p := range g.Permissions {
			if p == perm {
				return true
			}
		}
	}
	return false
}
