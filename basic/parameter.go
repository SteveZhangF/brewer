package basic

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type Filter struct {
	Left     string      `json:"left"`
	Right    interface{} `json:"right"`
	Operator string      `json:"operator"`
}

func ValueToString(v interface{}) string {
	result := ""
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(v)
		result = "("
		for i := 0; i < s.Len(); i++ {
			if result == "(" {
				result = result + fmt.Sprintf("%v", s.Index(i))
			} else {
				result = result + "," + fmt.Sprintf("%v", s.Index(i))
			}
		}
		result = result + ")"
	default:
		result = fmt.Sprintf("%v", reflect.ValueOf(v))
	}
	return result
}

func (f *Filter) String() string {
	return fmt.Sprintf("%v_%v_%v", f.Left, f.Operator, f.Right)
}

func (f *Filter) Expression() string {
	o := strings.ToUpper(f.Operator)
	switch o {
	case "IN":
		v := ValueToString(f.Right)
		return "(" + f.Left + " IN " + v + ")"
	case "NOT IN":
		v := ValueToString(f.Right)
		return "(!(" + f.Left + " IN " + v + "))"
	case "RAW":
		return "(" + f.Left + ")"
	case "=":
		return fmt.Sprintf("(%v == %v)", f.Left, f.Right)
	default:
		return fmt.Sprintf("(%v %v %v)", f.Left, f.Operator, f.Right)
	}
}

type CommonParameter struct {
	Sorts   string    `json:"sorts"`
	Page    int       `json:"page"`
	PerPage int       `json:"per_page"`
	Filters []*Filter `json:"filters"`
	Fields  []string  `json:"fields"`
}

type Order string

const (
	DESC = "desc"
	ASC  = "asc"
)

func (cp CommonParameter) MSorts() (string, Order) {
	sp := strings.Split(cp.Sorts, "|")
	if len(sp) == 2 {
		return sp[0], getOrder(sp[1])
	}
	return "", ASC
}

func (cp *CommonParameter) Valid() error {
	if cp.Page <= 0 {
		cp.Page = 1
	}

	if cp.PerPage <= 0 {
		cp.PerPage = 40
	}

	return nil
}

func (cp *CommonParameter) AddFilter(left string, operator string, right interface{}) {
	f := &Filter{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
	cp.Filters = append(cp.Filters, f)
}

func (cp *CommonParameter) String() string {
	js, _ := json.Marshal(cp)
	return string(js)
}

func getOrder(s string) Order {
	s = strings.ToUpper(s)
	switch s {
	case "-1", "DESC":
		return DESC
	case "1", "ASC":
		return ASC
	}
	return ASC
}
