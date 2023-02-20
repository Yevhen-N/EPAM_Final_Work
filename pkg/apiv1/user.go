package apiv1

import "fmt"

type UserResponse struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Lock     bool   `json:"lock"`
	Admin    bool   `json:"admin"`

	Accounts []AccountResponse `json:"accounts"`
	Logs     []LogResponse     `json:"logs"`
}

type UserLockRequest struct {
	ID   int64 `json:"id"`
	Lock bool  `json:"lock"`
}

func (u *UserLockRequest) Validate() error {
	if u.ID == 0 {
		return fmt.Errorf("id must not be empty")
	}
	return nil
}

type UserRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserRequest) Validate() error {
	if u.FullName == "" {
		return fmt.Errorf("user name must not be empty")
	}
	if u.Email == "" {
		return fmt.Errorf("user e-mail must not be empty")
	}
	if u.Password == "" {
		return fmt.Errorf("user password must not be empty")
	}
	return nil
}
