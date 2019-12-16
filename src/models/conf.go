package models

import (
	"time"
)
// 用户表
type User struct {
	Uid                   int64  `json:"uid"`
	Username              string  `json:"username"`
	Password              string  `json:"password"`
	UserRole             int64  `json:"user_role"`
	HeadSculpturePath   string  `json:"head_sculpture_path"`
	Label                 string   `json:"label"`
	Mailbox               string `json:"mailbox"`
	CreateTime           time.Time   `json:"create_time"`
	Status                int64  `json:status`
}
// 私信表
type PrivateLetter struct {
	PLid    int64  `json:"plid"`
	UserId   int64  `json:"user_id"`
	FriendId  int64  `json:"friend_id"`
	SenderId    int64 `json:"sender_id"`
	ReceiverId  int64 `json:"receiver_id"`
	MassageType   int64 `json:"massage_type"`
	MassageContent string `json:"massage_content"`
	SendTime   time.Time  `json:"send_time"`
	Status          int64 `json:"status"`
}
// 货品表
type Goods struct {
	Gid   int64 `json:"gid"`
	Uid   int64 `json:"uid"`
	Gname string `json:"gname"`
	Gprice  float64 `json:"gprice"`
	Gdetail string `json:"gdetail"`
	CategoryId  int64 `json:"category_id"`
	Gtid int64 `json:"gt_id"`
	ClickNumber int64 `json:"click_number"`
	Status  int64  `json:"status"`
	MobilePhoneNumber string `json:"mobilephone_number"`
	Gliaison  string `json:"gliaison"`
	Openid  string `json:"openid"`
	Qq       string `json:"qq"`
	ReleaseTime time.Time `json:"release_time"`

	Username string `json:"username"`
	TypeName string `json:"type_name"`
	Category string `json:"category"`
	FirstImgPath string `json:"first_img_path"`
}
type Image struct {
	Gid float64         `json:"gid"`       // 货品id
	Image_name string   `json:"image_name"`  // 图片名称
	Image_ext string    `json:"image_ext"`    // 图片扩展名
	Save_path string    `json:"save_path"`   // 图片路径
	Image_size float64  `json:"image_size"`     // 图片大小
}
// 货品标签表
type GoodsLabel struct {
	Gid  int64 `json:"gid"`
	LabelName  string `json:"label_name"`
	CreateTime  time.Time `json:"create_time"`
}

// 货品种类表
// 货品标签表
type GoodsType struct {
	Gtid  int64 `json:"gt_id"`
	TypeName  string `json:"type_name"`
	CreateTime  time.Time `json:"create_time"`
}
// 货品图片表
type GoodsImg struct {
	Id  int64  `json:"id"`
	Gid int64  `json:"gid"`
	ImageName string `json:"image_name"`
	ImageExt  string  `json:"image_ext"`
	SavePath  string  `json:"save_path"`
	ImageSize float64 `json:"image_size"`
	ReleaseTime time.Time `json:"release_time"`
}

// 收藏表
type Favour struct {
	Fid  int64 `json:"fid"`
	Gid  int64 `json:"gid"`
	//Uid  int64 `json:"uid"`
	//CreateTime time.Time `json:"create_time"`
}

//收藏夹
type FavourDir struct{
	Fdid                int64       `json:"fd_id"`
	Foldername          string    `json:"foldername"`
	CreateTime         time.Time `json:"create_time"`
	Uid                 int64       `json:"uid"`
	Sketch              string    `json:"sketch"`
	AuthorityLevel     int64       `json:"authority_level"`
}