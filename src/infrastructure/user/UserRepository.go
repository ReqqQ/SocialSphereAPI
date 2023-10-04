package infrastructureuser

import "fmt"

type UserRepository struct {
}

func (u UserRepository) Invoke() interface{} {
	fmt.Println("echo")

	return UserRepository{}
}
