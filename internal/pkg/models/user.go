package models

type User struct {
	ID           uint   `json:"-"`
	Login        string `json:"login"`
	Email        string `json:"email"`
	Password     string `json:"password,omitempty"`
	HashPassword string `json:"-"`
	LastVisit    string `json:"lastVisit"`
	Score        uint   `json:"score"`
	Picture      string `json:"avatarUrl"`
	Games        uint   `json:"count"`
}

type Session struct {
	Email     string `json:"email"`
	SessionID string `json:"sessionid"`
}

var Sessions map[string]string

func init() {
	Sessions = make(map[string]string)
}

func ReturnCountOfSessions() int {
	return len(Sessions)
}

type LeadersInfo struct {
	Count       uint `json:"count"`
	UsersOnPage uint `json:"usersOnPage"`
}
