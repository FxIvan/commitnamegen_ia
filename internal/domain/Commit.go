package domain

import "time"

type Commit struct {
	Type        string    `json:"type" binding:"required"`
	IDTicket    int       `json:"idticket" binding:"required"`
	Description string    `json:"description" binding:"required"`
	MetaInfo    time.Time `json:"creation_time"`
}
