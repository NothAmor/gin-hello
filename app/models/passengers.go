package models

import "time"

type Passengers struct {
	ID             int       `json:"id" gorm:"primary_key;autoIncrement;uniqueIndex"` //乘机人ID
	Email          string    `gorm:"not null"`                                        //此条数据应属于这个邮箱的用户
	Name           string    `gorm:"not null"`                                        //乘机人姓名
	IdentifyMethod string    `gorm:"not null"`                                        //证件类型
	IDCardNumber   string    `gorm:"not null"`                                        //乘机人身份证号
	PassportNumber string    `gorm:"not null"`                                        //护照号
	Birthday       time.Time `gorm:"not null"`                                        //出生日期
	PhoneNumber    string    `gorm:"not null"`                                        //手机号
}
