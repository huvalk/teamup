package feed

import "diplomaProject/application/models"

type Repository interface {
	Get(feedID int) (*models.Feed, error)
	GetByEvent(eventID int) (*models.Feed, error)
	Create(eventID int) (*models.Feed, error)
	AddUser(uid, eventID int) error
	RemoveUser(uid, eventID int) error
}
