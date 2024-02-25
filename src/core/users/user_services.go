package users

type UserServices struct {
	Repository  UserRepository
	AuthService AuthService
}

// NewUserService returns contract implementation of UserUseCases
func NewUserService(repository UserRepository, authService AuthService) UserUseCases {
	return &UserServices{
		Repository:  repository,
		AuthService: authService,
	}
}

// Login implements core.UserUseCases.
func (u *UserServices) Login(auth *AuthInput) (*AuthOutput, error) {
	hashedPassword, err := u.AuthService.HashPassword(auth.Password)
	if err != nil {
		return nil, err
	}
	auth.Password = hashedPassword

	user, err := u.Repository.Login(auth)
	if err != nil {
		return nil, err
	}

	token, err := u.AuthService.GenerateToken(user.ID, user.Name)
	if err != nil {
		return nil, err
	}

	authOutput := &AuthOutput{
		Token:    token,
		UserId:   user.ID,
		UserName: user.Name,
	}

	return authOutput, nil
}

// Create implements core.UserUseCases.
func (u *UserServices) Create(user *UserInput) (*User, error) {
	hashedPassword, err := u.AuthService.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword

	newUser, err := u.Repository.Create(user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

// Delete implements core.UserUseCases.
func (u *UserServices) Delete(id int64) (*User, error) {
	userDeleted, err := u.Repository.Delete(id)
	if err != nil {
		return nil, err
	}

	return userDeleted, nil
}

// GetOne implements core.UserUseCases.
func (u *UserServices) GetOne(id ...int64) (*User, error) {
	user, err := u.Repository.GetOne(id...)
	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, nil
	}

	return user, nil
}

// Update implements core.UserUseCases.
func (u *UserServices) Update(id int64, user *UserInput) (*User, error) {
	newPassword, err := u.AuthService.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = newPassword

	userUpdated, err := u.Repository.Update(id, user)
	if err != nil {
		return nil, err
	}

	return userUpdated, nil
}
