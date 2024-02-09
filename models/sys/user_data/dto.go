package user_data

import "go-deploy/models/dto/v1/body"

func (ud *UserData) ToDTO() body.UserDataRead {
	return body.UserDataRead{
		ID:   ud.ID,
		Data: ud.Data,
	}
}