package model

import (
	"encoding/json"
	"fmt"
	"kenneth/backend/basic"
	"kenneth/backend/basic/errors"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (c *Audit) Count(param *basic.CommonParameter) (int64, error) {
	var count int64

	t := db.Where("1=1").Model(&Audit{})
	for _, f := range param.Filters {
		left := f.Left
		if strings.ToUpper(f.Left) == "SEARCH" {
			t = t.Where("name LIKE ? OR code LIKE ?", fmt.Sprintf("%%%v%%", f.Right), fmt.Sprintf("%%%v%%", f.Right))
			continue
		}
		if strings.ToUpper(f.Operator) == "IN" {
			t = t.Where(fmt.Sprintf(" %v %v (?)", left, f.Operator), f.Right)
		} else if strings.ToUpper(f.Operator) == "LIKE" {
			t = t.Where(fmt.Sprintf(" %v %v ?", left, f.Operator), fmt.Sprintf("%%%v%%", f.Right))
		} else {
			t = t.Where(fmt.Sprintf(" %v %v ?", left, f.Operator), f.Right)
		}
	}
	err := t.Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, errors.InternalErrorHappened.Error(err)
	}
	return count, nil
}

func (c *Audit) List(param *basic.CommonParameter) ([]*Audit, error) {
	var audits []*Audit
	if param.Sorts == "" {
		param.Sorts = "name|ASC"
	}
	sort, order := param.MSorts()

	t := db.Where("1=1").Preload(clause.Associations)
	for _, f := range param.Filters {
		left := f.Left
		if strings.ToUpper(f.Left) == "SEARCH" {
			t = t.Where("name LIKE ? OR code LIKE ?", fmt.Sprintf("%%%v%%", f.Right), fmt.Sprintf("%%%v%%", f.Right))
			continue
		}
		if strings.ToUpper(f.Operator) == "IN" {
			t = t.Where(fmt.Sprintf(" %v %v (?)", left, f.Operator), f.Right)
		} else if strings.ToUpper(f.Operator) == "LIKE" {
			t = t.Where(fmt.Sprintf(" %v %v ?", left, f.Operator), fmt.Sprintf("%%%v%%", f.Right))
		} else {
			t = t.Where(fmt.Sprintf(" %v %v ?", left, f.Operator), f.Right)
		}
	}
	t = t.Order(sort + " " + string(order))

	if param.Page != 0 {
		t = t.Offset((param.Page - 1) * param.PerPage).Limit(param.PerPage)
	}

	err := t.Find(&audits).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.InternalErrorHappened.Error(err)
	}
	return audits, nil
}

func (c *Audit) Search(param *basic.CommonParameter) (basic.CommonListResult, error) {
	param.Valid()
	count, err := c.Count(param)
	if err != nil {
		return basic.CommonListResult{}, err
	}
	if count == 0 {
		return basic.CommonListResult{}, nil
	}
	audits, err := c.List(param)
	if err != nil {
		return basic.CommonListResult{}, err
	}
	return basic.CommonListResult{
		Total:   int(count),
		Data:    audits,
		Page:    param.Page,
		PerPage: param.PerPage,
	}, nil
}

func (c *Audit) Find(id uint) (*Audit, error) {
	audit := &Audit{}
	err := db.Preload(clause.Associations).Where("id = ?", id).First(audit).Error
	if gorm.ErrRecordNotFound == err {
		return nil, errors.NotFound.Error(err)
	}
	if err != nil {
		return nil, err
	}
	return audit, nil
}

func (c *Audit) Create(data *Audit) (*Audit, error) {
	audit := &Audit{}
	js, _ := json.Marshal(data)
	json.Unmarshal(js, audit)
	err := db.Omit(clause.Associations).Create(audit).Error
	if err != nil {
		return nil, err
	}
	return audit, nil
}

func (c *Audit) Update(id uint, data map[string]interface{}) (*Audit, error) {
	audit := &Audit{}
	err := db.Where("id = ?", id).First(audit).Error
	if gorm.ErrRecordNotFound == err {
		return nil, errors.NotFound.Error(err)
	}
	if err != nil {
		return nil, err
	}
	js, _ := json.Marshal(data)
	json.Unmarshal(js, audit)
	err = db.Omit(clause.Associations).Save(audit).Error
	if err != nil {
		return nil, err
	}
	return audit, nil
}

func (c *Audit) Delete(id uint) error {
	return db.Delete(&Audit{}, id).Error
}
