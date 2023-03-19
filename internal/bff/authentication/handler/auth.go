package handler

type AuthHandler interface {
	Login(c *fib)
}

type authHandler struct {
}
