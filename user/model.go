package user

type User struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
	PaidUser bool `json:"padid_user"`
	FavoriteTopic1 string `json:"favorite_topic_1"`
	FavoriteTopic2 string `json:"favorite_topic_2"`
	FavoriteTopic3 string `json:"favorite_topic_3"`
	FavoriteTopic4 string `json:"favorite_topic_4"`
	FavoriteTopic5 string `json:"favorite_topic_5"`
}
