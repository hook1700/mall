package resp

type User struct {
	Id string `json:"id"`
	Key string `json:"key"`
	UserId string `json:"userId"`
	NickName string `json:"nickName"`
	Mobile string `json:"mobile"`
	Address string `json:"address"`
	IsDeleted bool `json:"isDeleted"`
	IsLocked bool `json:"isLocked"`
}
