package controller

type Controller struct {
	User                UserController
	Location            LocationController
	Business            BusinessController
	BusinessCategory    BusinessCategoryController
	BusinessTransaction BusinessTransactionController
}
