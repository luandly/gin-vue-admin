// 自动生成模板UserWallet
package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// UserWallet 结构体
type UserWallet struct {
	global.GVA_MODEL
	UserID    *int     `json:"userID" form:"userID" gorm:"column:user_id;comment:;"`
	CompanyID *int     `json:"companyID" form:"companyID" gorm:"column:company_id;comment:;"`
	Wages     *float64 `json:"wages" form:"wages" gorm:"column:wages;type:decimal(10,2);comment:;"`
	CreatedBy uint     `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint     `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint     `gorm:"column:deleted_by;comment:删除者"`
}

// TableName UserWallet 表名
func (UserWallet) TableName() string {
	return "user_wallet"
}
