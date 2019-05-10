package user

type userMongoRepositrory interface {
	FindByID()
	FindAll()
}
