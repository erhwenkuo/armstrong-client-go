package armstrong

// SingIn - Get a new token for user
func (c *Client) SignIn() (*AuthResponse, error) {
	// TODO: call OAuth server to retrieve token
	ar := AuthResponse{
		Username: c.Auth.Username,
		Password: c.Auth.Password,
		Token: "xxxx-xxxx-xxxx",
	}

	return &ar, nil
}

// SingOut - Revoke the token for a user
func (c *Client) SignOut() error {
	// TODO: call OAuth server to revoke token
	return nil
}