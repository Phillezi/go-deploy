package user

type CreateParams struct {
	Username      string
	FirstName     string
	LastName      string
	Email         string
	IsAdmin       bool
	EffectiveRole *EffectiveRole
}

type UpdateParams struct {
	PublicKeys *[]PublicKey
	Onboarded  *bool
}
