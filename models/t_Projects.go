package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type TProjects struct {
	Id                      int       `orm:"column(project_ID);auto"`
	ProjectName             string    `orm:"column(project_Name);size(255)"`
	ProjectDescription      string    `orm:"column(project_Description);size(1024)"`
	ProjectSource           int       `orm:"column(project_Source);null"`
	ProjectFile             int       `orm:"column(project_File);null"`
	ProjectExpectedAmount   string    `orm:"column(project_ExpectedAmount);null"`
	ProjectExpectedDuration string    `orm:"column(project_ExpectedDuration);null"`
	ProjectDealAmount       string    `orm:"column(project_DealAmount);null"`
	ProjectDealDuration     string    `orm:"column(project_DealDuration);null"`
	ProjectStatus           int       `orm:"column(project_Status);null"`
	ProjectType             int       `orm:"column(project_Type);null"`
	ProjectWonerID          int       `orm:"column(project_OwnerID);null"`
	ProjectCreatedTime      time.Time `orm:"column(project_CreatedTime);type(datetime)"`
	ProjectStartTime        time.Time `orm:"column(project_StartTime);type(datetime)"`
	ProjectCompleteTime     time.Time `orm:"column(project_CompleteTime);type(datetime)"`
	ProjectEndTime          time.Time `orm:"column(project_EndTime);type(datetime)"`
	ProjectInfoUpdateTime   time.Time `orm:"column(project_InfoUpdateTime);type(timestamp)"`
}

func (t *TProjects) TableName() string {
	return "t_Projects"
}

func init() {
	orm.RegisterModel(new(TProjects))
}

// AddTProjects insert a new TProjects into database and returns
// last inserted Id on success.
func AddTProjects(m *TProjects) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTProjectsById retrieves TProjects by Id. Returns error if
// Id doesn't exist
func GetTProjectsById(id int) (v *TProjects, err error) {
	o := orm.NewOrm()
	v = &TProjects{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTProjects retrieves all TProjects matches certain condition. Returns empty list if
// no records exist
func GetAllTProjects(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TProjects))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []TProjects
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateTProjects updates TProjects by Id and returns error if
// the record to be updated doesn't exist
func UpdateTProjectsById(m *TProjects) (err error) {
	o := orm.NewOrm()
	v := TProjects{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTProjects deletes TProjects by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTProjects(id int) (err error) {
	o := orm.NewOrm()
	v := TProjects{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TProjects{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
