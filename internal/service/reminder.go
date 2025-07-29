package service

import "github.com/piftai/mindfulBotV2/internal/port"

type Reminder interface {
	Add(reminder ReminderAdapter) (*port.Reminder, error)
}

type Handler struct {
}

func (s *Handler) Add(reminder ReminderAdapter) (*port.Reminder, error) {

}
