package model

import (
	"math/rand"
	"time"
)

// BasicModel 基础模型
// 用于定义公共字段
type BasicModel struct {
	ID         int        `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time  `json:"created_at"`
	CreatedBy  int        `json:"created_by"`
	UpdatedAt  *time.Time `json:"updated_at"`
	UpdatedBy  *int       `json:"updated_by"`
	RemoveFlag bool       `json:"remove_flag"`
}

// SetCreatedInfo 设置模型创建信息
func (m BasicModel) SetCreatedInfo() {
	rand.Seed(time.Now().Unix())
	m.CreatedAt = time.Now()
	// TODO 完善获取创建人信息的方法
	m.CreatedBy = rand.Intn(20)
	m.RemoveFlag = false
}

// SetUpdateInfo 设置模型更新信息
func (m BasicModel) SetUpdateInfo() {
	rand.Seed(time.Now().Unix())
	n := time.Now()
	by := rand.Intn(20)
	m.UpdatedAt = &n
	// // TODO 完善获取创建人信息的方法
	m.UpdatedBy = &by
}
