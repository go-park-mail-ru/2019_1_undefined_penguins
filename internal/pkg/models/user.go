package models

//easyjson:json
type EasyJSONUser struct {
	ID                   uint64   `protobuf:"varint,1,opt,name=ID,json=iD,proto3" json:"ID"`
	Login                string   `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	HashPassword         string   `protobuf:"bytes,5,opt,name=hashPassword,proto3" json:"hashPassword,omitempty"`
	Score                uint64   `protobuf:"varint,6,opt,name=score,proto3" json:"score"`
	Picture              string   `protobuf:"bytes,7,opt,name=picture,proto3" json:"picture,omitempty"`
	Count                uint64   `protobuf:"varint,8,opt,name=count,proto3" json:"count"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (eju *EasyJSONUser) ToModelUser(user *User) *EasyJSONUser {
	result := new(EasyJSONUser)
	result.ID = user.ID
	result.Login = user.Login
	result.Email = user.Email
	result.Password = user.Password
	result.HashPassword = user.HashPassword
	result.Score = user.Score
	result.Picture = user.Picture
	result.Count = user.Count
	//result.XXX_NoUnkeyedLiteral = user.XXX_NoUnkeyedLiteral
	//result.XXX_unrecognized = user.XXX_unrecognized
	//result.XXX_sizecache = user.XXX_sizecache
	return result
}
var AuthManager AuthCheckerClient

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

type LeadersInfo1 struct {
	Count       uint `json:"count"`
	UsersOnPage uint `json:"usersOnPage"`
}
