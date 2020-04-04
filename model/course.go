package model

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"math"
)

type Courses struct {
	Id       int    `gorm:"primary_key"`
	Question string `gorm:"type:varchar(500);not null;unique;"`
	Answer   string `gorm:"type:varchar(500);not null;unique;"`
}

//func (c *Courses) TableName() string {
//	return "chati"
//}

func (c Courses) FindQuestion(nameLike string) ([]Courses, int, error) {
	if nameLike == "" {
		return nil, 0, errors.New("未查询到改题目的记录")
	}
	rows := make([]Courses, 0)
	var total int
	err := DB.Where("question like ?", fmt.Sprintf("%%%s%%", nameLike)).Find(&rows).Count(&total).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return rows, 0, errors.New("未查询到改题目的记录")
		}
		return rows, 0, errors.New("服务器内部错误")
	}
	return rows, total, nil
}

// Page 分页返回
type Page struct {
	TotalRecord int         `json:"total_record"`
	TotalPage   int         `json:"total_page"`
	Records     interface{} `json:"records"`
	Offset      int         `json:"offset"`
	Limit       int         `json:"limit"`
	NextCursor  int         `json:"next_cursor"`
}

var ErrPageParamError = fmt.Errorf("param 'result' err：%s", "unsupported destination, should be slice or struct")

// Pagination 分页
// @param db 数据库连接
// @param limit 每页条数
// @param offset 偏移量
// @param result 需要查询的结果集
func Pagination(db *gorm.DB, limit, offset int, result interface{}) (*Page, error) {
	// 如果每页条数<=0,则初始化为10条
	if limit <= 0 {
		limit = 10
	}
	// 如果偏移量小于0，则从0开始
	if offset < 0 {
		offset = 0
	}
	if result == nil {
		return nil, ErrPageParamError
	}

	var (
		page  = Page{}
		count int
	)

	err := db.Model(result).Count(&count).Error
	if err != nil {
		return nil, err
	}

	err = db.Limit(limit).Offset(offset).Find(result).Error

	if err != nil {
		return nil, err
	}

	page.TotalRecord = count
	page.Records = result

	page.Offset = offset
	page.Limit = limit
	page.TotalPage = int(math.Ceil(float64(count) / float64(limit)))

	if count > limit+offset {
		nextCursor := offset + limit
		page.NextCursor = nextCursor
	} else {
		page.NextCursor = 0
	}

	return &page, nil
}
