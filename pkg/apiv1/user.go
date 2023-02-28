package apiv1

import (
	"fmt"
	"regexp"

	"github.com/Yevhen-N/EPAM_Final_Work/pkg/db/model"
	"github.com/Yevhen-N/EPAM_Final_Work/pkg/utils/generator"
)

type UserResponse struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Role     string `json:"role"`

	Accounts []AccountResponse `json:"accounts"`
	Logs     []LogResponse     `json:"logs"`
}

type UserStatusRequest struct {
	ID     int64  `json:"id"`
	Status string `json:"status" enum:"active, blocked"`
}

func (u *UserStatusRequest) Validate() error {
	if u.ID == 0 {
		return fmt.Errorf("id must not be empty")
	}

	switch u.Status {
	case model.UserStatusActive, model.UserStatusBlocked:
		// nothing to do
	default:
		return fmt.Errorf("unsupported status: %s", u.Status)
	}
	return nil
}

type UserRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var emailRegexValidation = regexp.MustCompile(generator.PassReg)

func (u *UserRequest) Validate() error {
	if u.FullName == "" {
		return fmt.Errorf("user name must not be empty")

	}
	if u.Email == "" {
		return fmt.Errorf("user e-mail must not be empty")

	}
	if !(emailRegexValidation.MatchString(u.Email)) {
		return fmt.Errorf("e-mail not format")

	}
	if u.Password == "" {
		return fmt.Errorf("user password must not be empty")
	}
	return nil
}
