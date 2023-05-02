// 自动生成模板AppUser
package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AppUser 结构体
type AppUser struct {
	global.GVA_MODEL
	UserName  string  `json:"userName" form:"userName" gorm:"column:user_name;comment:;"`
	Password  string  `json:"password" form:"password" gorm:"column:password;comment:;size:64;"`
	Nickname  string  `json:"nickname" form:"nickname" gorm:"column:nickname;comment:;size:20;"`
	Wages     float64 `json:"wages" form:"wages" gorm:"column:wages;type:decimal(10,2);comment:;"`
	Status    *int    `json:"status" form:"status" gorm:"column:status;comment:;size:1;"`
	PhoneNum  string  `json:"phoneNum" form:"phoneNum" gorm:"column:phone_num;comment:;size:11;"`
	OpenID    string  `json:"openID" form:"openID" gorm:"column:open_id;comment:;size:64;"`
	UnionID   string  `json:"unionID" form:"unionID" gorm:"column:union_id;comment:;size:64;"`
	CreatedBy uint    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint    `gorm:"column:deleted_by;comment:删除者"`
}

// TableName AppUser 表名
func (AppUser) TableName() string {
	return "app_user"
}
