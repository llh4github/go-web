package model

// Role 角色信息模型
type Role struct {
	BasicModel
	RoleName string `json:"role_name" gorm:"size:50"`
	Remark   string `json:"username" gorm:"size:255"`
}
