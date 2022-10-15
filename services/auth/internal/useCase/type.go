package useCase

type SignInResult struct {
	Token   string
	Refresh string
}

type RefreshResult struct {
	Token   string
	Refresh string
}
