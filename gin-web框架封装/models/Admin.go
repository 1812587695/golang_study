package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Admin struct {
	ID         int       `gorm:"primary_key" json:"id"`
	Account    string    `json:"account"`
	Phone      string    `json:"phone"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Qq         string    `json:"qq"`
	Password   string    `json:"-"`
	Department string    `json:"department"`
	Position   string    `json:"position"`
	LoggedIp   string    `json:"logged_ip"`
	IsEnabled  int       `json:"is_enabled"`
	Remarks    string    `json:"remarks"`
	LoggedAt   time.Time `json:"logged_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
type AdminDetailsResult struct {
	ID         int       `gorm:"primary_key" json:"id"`
	Role_id    int       `json:"role_id"`
	Account    string    `json:"account"`
	Phone      string    `json:"phone"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Qq         string    `json:"qq"`
	Password   string    `json:"-"`
	Department string    `json:"department"`
	Position   string    `json:"position"`
	LoggedIp   string    `json:"logged_ip"`
	IsEnabled  int       `json:"is_enabled"`
	Remarks    string    `json:"remarks"`
	LoggedAt   time.Time `json:"logged_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type AdminListResult struct {
	Id         int       `json:"id"`
	Account    string    `json:"account"`
	RoleName   string    `json:"role_name"`
	Department string    `json:"department"`
	Position   string    `json:"position"`
	UpdatedAt  time.Time `json:"updated_at"`
	IsEnabled  int       `json:"is_enabled"`
}

func (a *Admin) TableName() string {
	return "admins"
}

// 查找用户
func FindAdminAccount(account string) *Admin {
	var admin Admin
	err := DB.Model(&Admin{}).Where("account=?", account).First(&admin).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return &admin
}

func GetAdmins(account string, role_id int, begin_time string, end_time string, page int, pageSize int) ([]*AdminListResult, *Paginator) {
	var admin []*AdminListResult

	item := DB.Table("admins as t1 ").Select("t1.id, t1.account, t3.name as role_name, t1.department, t1.position, t1.updated_at,t1.is_enabled").Joins("left join admin_role as t2 on t1.id = t2.admin_id ").Joins("left join roles as t3 on t2.role_id= t3.id").Find(&admin)
	if account != "" {
		item = item.Where("t1.account like ?", "%"+account+"%")
	}
	if role_id != 0 {
		item = item.Where("t3.id=?", role_id)
	}
	if begin_time != "" {
		item = item.Where("t1.created_at > ?", begin_time)
	}
	if end_time != "" {
		item = item.Where("t1.created_at < ?", end_time)
	}

	var count int
	item.Count(&count)
	paginate := NewPage(page, pageSize, count)
	err := item.Offset(paginate.Offset).Limit(paginate.Limit).Find(&admin).Error
	if err != nil {
		panic(err)
	}
	return admin, &paginate
}

// 通过id查找单个用户
func GetAdminById(id int) (*AdminDetailsResult, error) {
	var admin AdminDetailsResult
	if err := DB.Table("admins as t1").Select("t1.*, t3.name as role_name, t3.id as role_id").Joins("left join admin_role as t2 on t1.id = t2.admin_id ").Joins("left join roles as t3 on t2.role_id = t3.id").Where("t1.id=?", id).Find(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

// 创建用户
func (admin *Admin) AddAdmin(role_id int) error {

	if err := DB.Model(&Admin{}).Create(&admin).Error; err != nil {
		return err
	}
	var admin_role = AdminRole{AdminId: admin.ID, RoleId: role_id}
	if err := DB.Table("admin_role").Create(admin_role).Error; err != nil {
		return err
	}
	return nil
}

// 修改用户

func (a *Admin) UpdateAdmin(role_id int) error {
	if err := DB.Table("admins").Where("id=?", a.ID).Updates(&a).Error; err != nil {
		return err
	}
	if role_id != 0 {
		DB.Delete(AdminRole{}, "admin_id=?", a.ID)
		var admin_role = AdminRole{AdminId: a.ID, RoleId: role_id}
		if err := DB.Table("admin_role").Create(admin_role).Error; err != nil {
			return err
		}
	}
	return nil
}

// 删除用户

func DeletedAdmin(id int) error {
	if err := DB.Where("id =?", id).Delete(&Admin{}).Error; err != nil {
		return err
	}
	return nil
}

func GetAdminRoleId(admin_id int) (int) {
	var role_id int
	admin_role := DB.Table("admin_role").Select("role_id").Where("admin_id=?", admin_id).Row()
	_ = admin_role.Scan(&role_id)

	return role_id
}

func (a *Admin) AdminRoleId() (int) {
	var role_id int
	admin_role := DB.Table("admin_role").Select("role_id").Where("admin_id=?", a.ID).Row()
	_ = admin_role.Scan(&role_id)

	return role_id
}
