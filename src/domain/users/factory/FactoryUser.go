package domainFactoryUser

type FactoryUserInterface interface {
	GetUserEntity() DomainFactoryUser
}

type DomainFactoryUser struct{}

func (f DomainFactoryUser) GetUserEntity() DomainFactoryUser {
	return DomainFactoryUser{}
}
