package models

import "time"
// 用户表
type User struct {
	Uid                   int  `json:"uid"`
	Username              string  `json:"username"`
	Password              string  `json:"password"`
	UserRole             int  `json:"user_role"`
	HeadSculpturePath   string  `json:"head_sculpture_path"`
	Label                 string   `json:"label"`
	CreateTime           time.Time   `json:"create_time"`
	Status                int  `json:status`
}
// 私信表
type PrivateLetter struct {
	PLid    int  `json:"plid"`
	UserId   int  `json:"user_id"`
	FriendId  int  `json:"friend_id"`
	SenderId    int `json:"sender_id"`
	ReceiverId  int `json:"receiver_id"`
	MassageType   int `json:"massage_type"`
	MassageContent string `json:"massage_content"`
	SendTime   time.Time  `json:"send_time"`
	Status          int `json:"status"`
}
// 货品表
type Goods struct {
	Gid   int `json:"gid"`
	Uid   int `json:"uid"`
	Gname string `json:"gname"`
	Gprice  float64 `json:"gprice"`
	Gdetail string `json:"gdetail"`
	CategoryId  int `json:"category_id"`
	ClickNumber int `json:"click_number"`
	Status  int  `json:"status"`
	MobilePhoneNumber string `json:"mobilephone_number"`
	Gliaison  string `json:"gliaison"`
	Openid  string `json:"openid"`
	Qq       string `json:"qq"`
	ReleaseTime time.Time `json:"release_time"`
}

// 货品标签表
type GoodsLabel struct {
	Gid  int `json:"gid"`
	LabelName  string `json:"label_name"`
	CreateTime  time.Time `json:"create_time"`
}

// 货品图片表
type GoodsImg struct {
	Id  int  `json:"id"`
	Gid int  `json:"gid"`
	ImageName string `json:"image_name"`
	ImageExt  string  `json:"image_ext"`
	SavePath  string  `json:"save_path"`
	ImageSize float64 `json:"image_size"`
	ReleaseTime time.Time `json:"release_time"`
}

// 收藏表
type Favour struct {
	Fid  int `json:"fid"`
	Gid  int `json:"gid"`
	//Uid  int `json:"uid"`
	//CreateTime time.Time `json:"create_time"`
}

//收藏夹
type FavourDir struct{
	Fdid                int       `json:"fdid"`
	Foldername          string    `json:"foldername"`
	CreateTime         time.Time `json:"create_time"`
	Uid                 int       `json:"uid"`
	Sketch              string    `json:"sketch"`
	AuthorityLevel     int       `json:"authority_level"`
}