package dao

type LogoutToken struct {
	TokenSeq int64  `gorm:"Column:token_seq;PRIMARY_KEY"`
	AdminSeq int64  `gorm:"Column:admin_seq"`
	Token    string `gorm:"Column:token"`
}

func (l LogoutToken) TableName() string {
	return "logout_token"
}
