package user

func userResponseFromDBModel(u UserDB) UserResponse {
	return UserResponse{
		Username: u.Username,
		Coins:    u.Coins,
	}
}

func userDBModelFromCreateRequest(r CreateUserRequest) UserDB {
	return UserDB{
		Username: r.Username,
		Password: r.Password,
		Coins:    1000,
	}
}
