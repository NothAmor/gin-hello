package models

import "time"

type Flights struct {
	FlightID                         int       `json:"id" gorm:"primary_key;autoIncrement;uniqueIndex"` // 航班唯一ID
	Company                          string    `gorm:"not null"`                                        // 航班所属公司
	SellerCompany                    string    `gorm:"not null"`                                        // 卖家所属公司
	SellerName                       string    `gorm:"not null"`                                        // 卖家姓名
	FlightNumber                     string    `gorm:"not null"`                                        // 航班号
	FlightType                       string    `gorm:"not null"`                                        // 客机类型 大，中，小，特殊
	FlightFood                       bool      `gorm:"not null"`                                        // 飞机途中是否提供餐食
	StartCity                        string    `gorm:"not null"`                                        // 飞机出发城市
	StartAirport                     string    `gorm:"not null"`                                        // 飞机出发机场
	StartTime                        time.Time `gorm:"not null"`                                        // 飞机出发时间
	ArriveCity                       string    `gorm:"not null"`                                        // 飞机到达城市
	ArriveAirport                    string    `gorm:"not null"`                                        // 飞机到达机场
	ArriveTime                       time.Time `gorm:"not null"`                                        // 飞机到达时间
	FirstClassSeatCapacity           int       // 头等舱总共座位数
	FirstClassSeatLeft               int       // 头等舱剩余座位数
	FirstClassTicketAdultPrice       int       // 头等舱成人票价
	FirstClassTicketChildrenPrice    int       // 头等舱儿童票价
	BusinessClassSeatCapacity        int       // 商务舱总共座位数
	BusinessClassSeatLeft            int       // 商务舱剩余座位数
	BusinessClassTicketAdultPrice    int       // 商务舱成人票价
	BusinessClassTicketChildrenPrice int       // 商务舱儿童票价
	EconomyClassSeatCapacity         int       // 经济舱总共座位数
	EconomyClassSeatLeft             int       // 经济舱剩余座位数
	EconomyClassTicketAdultPrice     int       // 经济舱成人票价
	EconomyClassTicketChildrenPrice  int       // 经济舱儿童票价
	AirportConstructionCost          int       // 机建费
	FuelCost                         int       // 燃油费
}
