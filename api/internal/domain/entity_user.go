package domain

type User struct {
	uuid            UUID
	googleAccountID GoogleAccountID
}

func NewUser(
	uuid UUID,
	googleAccountID GoogleAccountID,
) User {
	return User{
		uuid:            uuid,
		googleAccountID: googleAccountID,
	}
}

func (u User) UUID() UUID                       { return u.uuid }
func (u User) GoogleAccountID() GoogleAccountID { return u.googleAccountID }
