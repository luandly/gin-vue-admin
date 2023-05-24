// 自动生成模板Order
package clothing

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Order 结构体
type Order struct {
	global.GVA_MODEL
	OrderNo    string     `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:;"`
	PayNo      string     `json:"payNo" form:"payNo" gorm:"column:pay_no;comment:;"`
	CompanyID  uint       `json:"companyID" form:"companyID" gorm:"column:company_id;comment:;"`
	UserID     uint       `json:"userID" form:"userID" gorm:"column:user_id;comment:;"`
	Price      float64    `json:"price" form:"price" gorm:"column:price;comment:;type:decimal(10,2);"`
	Amount     float64    `json:"amount" form:"amount" gorm:"column:amount;comment:;type:decimal(10,2);"`
	Status     int        `json:"status" form:"status" gorm:"column:status;comment:;"`
	PayStatus  int        `json:"payStatus" form:"payStatus" gorm:"column:pay_status;comment:;"`
	Day        int        `json:"day" form:"day" gorm:"column:day;comment:;"`
	ClerkCount int        `json:"clerkCount" form:"clerkCount" gorm:"column:clerk_count;comment:;"`
	PayAt      *time.Time `json:"payAt" form:"payAt" gorm:"column:pay_at;comment:;"`
	CreatedBy  uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy  uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy  uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName Order 表名
func (Order) TableName() string {
	return "order"
}
