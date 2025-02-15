package controller

type Controller struct {
	Authentication      AuthenticationController
	User                UserController
	Location            LocationController
	Business            BusinessController
	BusinessCategory    BusinessCategoryController
	BusinessTransaction BusinessTransactionController
}
