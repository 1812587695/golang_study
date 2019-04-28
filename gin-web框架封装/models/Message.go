package models

import "time"

type Message struct {
	ID int
	From int
	To int
	Type int
	Target int
	IsRead int
	CreatedAt time.Time
}

type MessageContent struct {
	ID int
	MsgId int
	Title string
	Content string
}

func NewMessage() *Message {
	return &Message{}
}


func (m *Message) TableName() string {
	return "media_message"
}

func (m *Message) Send() bool {
	err := DB.Model(m).Create(m).Error
	if err != nil {
		return false
	}
	return true
}

func (m *Message) SetContent(title string, content string) bool {
	err := DB.Table("media_message_content").Create(&MessageContent{MsgId:m.ID,Title:title, Content:content}).Error
	if err != nil {
		return false
	}
	return true
}