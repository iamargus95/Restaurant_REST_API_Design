package resources

type SignupPayload struct {
	Name  string `binding:"required"`
	Email string `binding:"required,email"`
	Phone string `binding:"required,min=10,max=10"`
}
