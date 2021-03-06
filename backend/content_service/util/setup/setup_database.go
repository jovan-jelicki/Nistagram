package setup

import (
	"github.com/david-drvar/xws2021-nistagram/content_service/model/persistence"
	"gorm.io/gorm"
)

func FillDatabase(db *gorm.DB) error {
	// dropTables(db)

	err := db.AutoMigrate(&persistence.Post{},
		&persistence.Story{},
		&persistence.Media{},
		&persistence.Tag{},
		&persistence.Collection{},
		&persistence.Favorites{},
		&persistence.Like{},
		&persistence.Comment{},
		&persistence.Highlight{},
		&persistence.HighlightStory{},
		&persistence.RegistrationRequest{},
		&persistence.Ad{},
		&persistence.Campaign{},
		&persistence.CampaignInfluencerRequest{},
		&persistence.ContentComplaint{},
		&persistence.AdCategory{},
		&persistence.UserAdCategories{},
		&persistence.CampaignChanges{},
		&persistence.Hashtag{},
		&persistence.HashtagObjava{},
	)

	return err
}

func dropTables(db *gorm.DB){
	if db.Migrator().HasTable(&persistence.Post{}) {
		db.Migrator().DropTable(&persistence.Post{},
			&persistence.Story{},
			&persistence.Media{},
			&persistence.Tag{},
			&persistence.Collection{},
			&persistence.Favorites{},
			&persistence.Like{},
			&persistence.Comment{},
			&persistence.Highlight{},
			&persistence.HighlightStory{},
			&persistence.RegistrationRequest{},
			&persistence.Ad{},
			&persistence.Campaign{},
			&persistence.CampaignInfluencerRequest{},
			&persistence.ContentComplaint{},
			&persistence.AdCategory{},
			&persistence.UserAdCategories{},
			&persistence.CampaignChanges{},
			&persistence.Hashtag{},
			&persistence.HashtagObjava{},
		)
	}
}
