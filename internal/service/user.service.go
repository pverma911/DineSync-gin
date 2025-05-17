package service

type UserService struct{}

type IUserService interface{
	GetUsers() []map[string]any
}


func (s *UserService)GetUsers() []map[string]any {
	return []map[string]any{
		{
			"id":   1,
			"name": "Alice",
			"age":  25,
		},
		{
			"id":   2,
			"name": "Trent",
			"age":  22,
		},
	}
}
