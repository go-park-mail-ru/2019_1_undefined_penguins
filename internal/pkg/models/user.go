package models

type User struct {
	ID           uint   `json:"-"`
	Login        string `json:"login"`
	Email        string `json:"email"`
	Password     string `json:"password,omitempty"`
	HashPassword string `json:"-"`
	LastVisit    string `json:"lastVisit"`
	Score        uint   `json:"score"`
	avatarName   string `json:"avatarName"`
	avatarBlob   string `json:"avatarBlob"`
}

// `type SignUpStruct struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// type MeStruct struct {
// 	Login string `json:"login"`
// 	Email string `json:"email"`
// 	Name  string `json:"name"`
// 	Score uint   `json:"score"`
// }`

type Session struct {
	Email     string `json:"email"`
	SessionID string `json:"sessionid"`
}

//в данный момент map[sessionid]email, но надо приводить к виду map[sessionid]id
var Sessions map[string]string

func init() {
	Sessions = make(map[string]string)
}

// var Users = map[string]User{
// 	"a.penguin1@corp.mail.ru": User{
// 		ID:           1,
// 		Login:        "Penguin1",
// 		Email:        "a.penguin1@corp.mail.ru",
// 		Name:         "Пингвин Северного Полюса",
// 		HashPassword: "$2a$14$9s00w8l7VKS2gRr2mtmg..1hvANedLWgmux3yOjkS80dTZlXLnKs2",
// 		LastVisit:    "25.02.2019",
// 		Score:        0,
// 		avatarName:   "default1.png",
// 		avatarBlob:   "./images/user.svg",
// 	},
// 	"b.penguin2@corp.mail.ru": User{
// 		ID:           2,
// 		Login:        "Penguin2",
// 		Email:        "b.penguin2@corp.mail.ru",
// 		Name:         "Пингвин Южного Полюса",
// 		HashPassword: "$2a$14$9s00w8l7VKS2gRr2mtmg..1hvANedLWgmux3yOjkS80dTZlXLnKs2",
// 		LastVisit:    "25.02.2019",
// 		Score:        100500,
// 		avatarName:   "default2.png",
// 		avatarBlob:   "./images/user.svg",
// 	},
// 	"c.penguin3@corp.mail.ru": User{
// 		ID:           3,
// 		Login:        "Penguin3",
// 		Email:        "c.penguin3@corp.mail.ru",
// 		Name:         "Залетный Пингвин",
// 		HashPassword: "$2a$14$9s00w8l7VKS2gRr2mtmg..1hvANedLWgmux3yOjkS80dTZlXLnKs2",
// 		LastVisit:    "25.02.2019",
// 		Score:        173,
// 		avatarName:   "default3.png",
// 		avatarBlob:   "./images/user.svg",
// 	},
// 	"d.penguin4@corp.mail.ru": User{
// 		ID:           4,
// 		Login:        "Penguin4",
// 		Email:        "d.penguin4@corp.mail.ru",
// 		Name:         "Рядовой Пингвин",
// 		HashPassword: "$2a$04$U2BYDHAfGa2cqJwlhSA2D.XyWD8kq1sAvh2s8nRlV5huDEJLF8pDu",
// 		LastVisit:    "25.02.2019",
// 		Score:        72,
// 		avatarName:   "default4.png",
// 		avatarBlob:   "./images/user.svg",
// 	},
// }

// end later remove hardcode
