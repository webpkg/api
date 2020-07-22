package model

// CreateAuth return *Auth
func CreateAuth() *Auth {

	auth := &Auth{}

	return auth
}

// Auth model
type Auth struct {
	UserID uint64 `json:"userID"`
	Right  int64  `json:"right"`
}

// CreateAuthCollection return *AuthCollection
func CreateAuthCollection() *AuthCollection {

	authCollection := &AuthCollection{}

	return authCollection
}

// AuthCollection Auth list
type AuthCollection []Auth

// Len return len
func (o *AuthCollection) Len() int { return len(*o) }

// Swap swap i, j
func (o *AuthCollection) Swap(i, j int) { (*o)[i], (*o)[j] = (*o)[j], (*o)[i] }

// Less compare i, j
func (o *AuthCollection) Less(i, j int) bool { return (*o)[i].UserID < (*o)[j].UserID }
