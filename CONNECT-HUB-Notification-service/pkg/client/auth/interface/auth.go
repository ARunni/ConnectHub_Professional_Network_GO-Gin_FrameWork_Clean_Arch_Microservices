package interfaces

import "github.com/ARunni/ConnetHub_Notification/pkg/utils/models"


type Newauthclient interface{
	UserData(userid int)(models.UserData,error)
}