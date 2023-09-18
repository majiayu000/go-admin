package models

import (
	"go-admin/common/models"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SysUser struct {
	Uid                 int       `gorm:"primaryKey;autoIncrement;comment:编码"  json:"f_uid"`
	Username            string    `json:"username" gorm:"column:f_userName; size:64;comment:用户名"`
	Password            string    `json:"-" gorm:"column:f_password; size:128;comment:密码"`
	NickName            string    `json:"nickName" gorm:"size:128;comment:昵称"`
	Phone               string    `json:"phone" gorm:"size:11;comment:手机号"`
	RoleId              int       `json:"roleId" gorm:"size:20;comment:角色ID"`
	Salt                string    `json:"-" gorm:"size:255;comment:加盐"`
	Avatar              string    `json:"avatar" gorm:"size:255;comment:头像"`
	Sex                 string    `json:"sex" gorm:"size:255;comment:性别"`
	Email               string    `json:"email" gorm:"size:128;comment:邮箱"`
	DeptId              int       `json:"deptId" gorm:"size:20;comment:部门"`
	PostId              int       `json:"postId" gorm:"size:20;comment:岗位"`
	Remark              string    `json:"remark" gorm:"size:255;comment:备注"`
	Status              string    `json:"status" gorm:"size:4;comment:状态"`
	DeptIds             []int     `json:"deptIds" gorm:"-"`
	PostIds             []int     `json:"postIds" gorm:"-"`
	RoleIds             []int     `json:"roleIds" gorm:"-"`
	Dept                *SysDept  `json:"dept"`
	InvitationCode      string    `gorm:"column:f_invitationCode"`
	InvitatedCode       string    `gorm:"column:f_invitatedCode"`
	MailBox             string    `gorm:"column:f_mailBox"`
	CreateTime          time.Time `gorm:"column:f_createTime"`
	IsBindGoogle        bool      `gorm:"column:f_isBindGoogle"`
	Secret              string    `gorm:"column:f_secret"`
	IsIDVerify          bool      `gorm:"column:f_isIDVerify"`
	Mobile              string    `gorm:"column:f_mobile"`
	InviteNumber        int       `gorm:"column:f_inviteNumber"`
	ClaimRewardNumber   int       `gorm:"column:f_claimRewardNumber"`
	ConcernCoinList     string    `gorm:"column:f_concernCoinList"`
	CollectStrategyList string    `gorm:"column:f_collectStrategyList"`
	UpdateTime          time.Time `gorm:"column:f_updateTime"`
	models.ControlBy
	models.ModelTime
}

// f_userName
// string
// 用户名
// f_role
// string
// 角色
// f_password
// string
// 密码
// f_createTime
// TIMESTAMP
// 插入时间
// f_updateTime
// TIMESTAMP
// 更新时间

// type SysUser struct {
// 	Id       int    `gorm:"type:int(11);primary_key;autoIncrement" json:"id"`
// 	Username string `gorm:"type:varchar(128);not null" json:"username"`
// 	Password string `gorm:"type:varchar(128);not null" json:"password"`
// 	Role     string `gorm:"type:varchar(128);not null" json:"role"`
// 	Phone    string `gorm:"type:varchar(128);not null" json:"phone"`

// 	BaseModel
// }

func (*SysUser) TableName() string {
	return "sys_user"
}

func (e *SysUser) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysUser) GetId() interface{} {
	return e.Uid
}

// Encrypt 加密
func (e *SysUser) Encrypt() (err error) {
	if e.Password == "" {
		return
	}

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost); err != nil {
		return
	} else {
		e.Password = string(hash)
		return
	}
}

func (e *SysUser) BeforeCreate(_ *gorm.DB) error {
	return e.Encrypt()
}

func (e *SysUser) BeforeUpdate(_ *gorm.DB) error {
	var err error
	if e.Password != "" {
		err = e.Encrypt()
	}
	return err
}

func (e *SysUser) AfterFind(_ *gorm.DB) error {
	e.DeptIds = []int{e.DeptId}
	e.PostIds = []int{e.PostId}
	e.RoleIds = []int{e.RoleId}
	return nil
}
