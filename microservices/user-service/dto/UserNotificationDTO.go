package dto

type UserNotificationDTO struct {
	Id string `json:"id"`
	Username string `json:"username"`
	ReceivePostNotifications bool `json:"receivePostNotifications"`
	ReceiveCommentNotifications bool `json:"receiveCommentNotifications"`
	ReceiveMessagesNotifications bool `json:"receive-messages-notifications"`
}

