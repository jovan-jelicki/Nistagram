package persistence

import (
	"github.com/david-drvar/xws2021-nistagram/content_service/model"
	"time"
)

type Post struct {
	Id          string `gorm:"primaryKey"`
	UserId      string
	IsAd        bool
	Type        model.PostType
	Description string
	Location    string
	CreatedAt   time.Time
}

type Story struct {
	Post
	IsCloseFriends bool
}

type Media struct {
	Id       string `gorm:"primaryKey"`
	Type     model.MediaType
	PostId   string
	Filename string
	OrderNum int
}

type Tag struct {
	UserId  string `gorm:"primaryKey"`
	MediaId string `gorm:"primaryKey"`
}

type Hashtag struct {
	Id   string `gorm:"primaryKey"`
	Text string `gorm:"unique"`
}

type HashtagObjava struct {
	HashtagId string `gorm:"primaryKey"`
	ObjavaId  string `gorm:"primaryKey"`
}

type Collection struct {
	Id     string `gorm:"primaryKey"`
	Name   string
	UserId string
}

type Favorites struct {
	PostId       string `gorm:"primaryKey"`
	UserId       string `gorm:"primaryKey"`
	CollectionId string
}

type Like struct {
	PostId string `gorm:"primaryKey"`
	UserId string `gorm:"primaryKey"`
	IsLike bool
}

type Comment struct {
	Id        string `gorm:"primaryKey"`
	PostId    string
	UserId    string
	Content   string
	CreatedAt time.Time
}

type Highlight struct {
	Id string `gorm:"primaryKey"`
	UserId string
	Name   string
}

type HighlightStory struct {
	HighlightId string `gorm:"primaryKey"`
	StoryId     string `gorm:"primaryKey"`
}

type RegistrationRequest struct {
	Id        string `gorm:"primaryKey"`
	UserId    string
	CreatedAt time.Time
	Status    model.RequestStatus
}

type Ad struct {
	Id           string `gorm:"primaryKey"`
	Link         string
	CampaignId   string
	PostId       string
	LinkClickNum int
}

type Campaign struct {
	Id           string `gorm:"primaryKey"`
	IsOneTime    bool
	StartDate    time.Time
	EndDate      time.Time
	PlacementNum int
	AgentId      string
	IdAdCategory string
	LastUpdated  time.Time
}

type CampaignInfluencerRequest struct {
	CampaignId   string `gorm:"primaryKey"`
	InfluencerId string `gorm:"primaryKey"`
	Status       model.RequestStatus
}

type ContentComplaint struct {
	Id       string `gorm:"primaryKey"`
	Category model.ComplaintCategory
	PostId   string
	Status   model.RequestStatus
}

type AdCategory struct {
	Id   string `gorm:"primaryKey"`
	Name string
}

type UserAdCategories struct {
	UserId       string `gorm:"primaryKey"`
	IdAdCategory string `gorm:"primaryKey"`
}

type CampaignChanges struct {
	CampaignId   string `gorm:"primaryKey"`
	AdCategoryId string
	StartDate    time.Time
	EndDate      time.Time
	PlacementNum int
}
