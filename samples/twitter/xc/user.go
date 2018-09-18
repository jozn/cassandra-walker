package xc

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gocql/gocql"
)

////////////////////////////////////////// Query seletor updater and deleter /////////////////////////

func (a *User) Exists() bool {
	return a._exists
}

func (a *User) Deleted() bool {
	return a._deleted
}

type User_Selector struct {
	wheres      []whereClause
	selectCol   []string
	orderBy     []string //" order by id desc //for ints
	limit       int
	allowFilter bool
}

type User_Updater struct {
	wheres  []whereClause
	updates map[string]interface{}
}

type User_Deleter struct {
	wheres    []whereClause
	deleteCol []string
}

//////////////////// just Selector
func (u *User_Selector) Limit(limit int) *User_Selector {
	u.limit = limit
	return u
}

func (u *User_Selector) AllowFiltering() *User_Selector {
	u.allowFilter = true
	return u
}

func NewUser_Selector() *User_Selector {
	u := User_Selector{}
	return &u
}

func NewUser_Updater() *User_Updater {
	u := User_Updater{}
	u.updates = make(map[string]interface{})
	return &u
}

func NewUser_Deleter() *User_Deleter {
	u := User_Deleter{}
	return &u
}

//each select columns

func (u *User_Selector) Select_CreatedTime() *User_Selector {
	u.selectCol = append(u.selectCol, "created_time")
	return u
}

//each column orders //just ints
func (u *User_Selector) OrderBy_CreatedTime_Desc() *User_Selector {
	u.orderBy = append(u.orderBy, " created_time DESC")
	return u
}

func (u *User_Selector) OrderBy_CreatedTime_Asc() *User_Selector {
	u.orderBy = append(u.orderBy, " created_time ASC")
	return u
}

func (u *User_Selector) Select_FullName() *User_Selector {
	u.selectCol = append(u.selectCol, "full_name")
	return u
}

//each column orders //just ints
func (u *User_Selector) OrderBy_FullName_Desc() *User_Selector {
	u.orderBy = append(u.orderBy, " full_name DESC")
	return u
}

func (u *User_Selector) OrderBy_FullName_Asc() *User_Selector {
	u.orderBy = append(u.orderBy, " full_name ASC")
	return u
}

func (u *User_Selector) Select_UserId() *User_Selector {
	u.selectCol = append(u.selectCol, "user_id")
	return u
}

//each column orders //just ints
func (u *User_Selector) OrderBy_UserId_Desc() *User_Selector {
	u.orderBy = append(u.orderBy, " user_id DESC")
	return u
}

func (u *User_Selector) OrderBy_UserId_Asc() *User_Selector {
	u.orderBy = append(u.orderBy, " user_id ASC")
	return u
}

func (u *User_Selector) Select_UserName() *User_Selector {
	u.selectCol = append(u.selectCol, "user_name")
	return u
}

//each column orders //just ints
func (u *User_Selector) OrderBy_UserName_Desc() *User_Selector {
	u.orderBy = append(u.orderBy, " user_name DESC")
	return u
}

func (u *User_Selector) OrderBy_UserName_Asc() *User_Selector {
	u.orderBy = append(u.orderBy, " user_name ASC")
	return u
}

//////////////////// just Deleter
//each column delete

func (u *User_Deleter) Delete_CreatedTime() *User_Deleter {
	u.deleteCol = append(u.deleteCol, "created_time")
	return u
}

func (u *User_Deleter) Delete_FullName() *User_Deleter {
	u.deleteCol = append(u.deleteCol, "full_name")
	return u
}

func (u *User_Deleter) Delete_UserId() *User_Deleter {
	u.deleteCol = append(u.deleteCol, "user_id")
	return u
}

func (u *User_Deleter) Delete_UserName() *User_Deleter {
	u.deleteCol = append(u.deleteCol, "user_name")
	return u
}

//////////////////// End of just Deleter

//////////////////// just Updater
//each column delete

func (u *User_Updater) CreatedTime(newVal int) *User_Updater {
	u.updates["created_time = ? "] = newVal
	return u
}

func (u *User_Updater) FullName(newVal string) *User_Updater {
	u.updates["full_name = ? "] = newVal
	return u
}

func (u *User_Updater) UserId(newVal int) *User_Updater {
	u.updates["user_id = ? "] = newVal
	return u
}

func (u *User_Updater) UserName(newVal string) *User_Updater {
	u.updates["user_name = ? "] = newVal
	return u
}

//////////////////// End just Updater

//{_Eq_Filtering  =  CreatedTime_Eq_Filtering}

func (d *User_Deleter) CreatedTime_Eq_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " created_time = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  CreatedTime_LT_Filtering}

func (d *User_Deleter) CreatedTime_LT_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " created_time < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  CreatedTime_LE_Filtering}

func (d *User_Deleter) CreatedTime_LE_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " created_time <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  CreatedTime_GT_Filtering}

func (d *User_Deleter) CreatedTime_GT_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " created_time > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  CreatedTime_GE_Filtering}

func (d *User_Deleter) CreatedTime_GE_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " created_time >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering And = And And_CreatedTime_Eq_Filtering}

func (d *User_Deleter) And_CreatedTime_Eq_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And created_time = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_CreatedTime_LT_Filtering}

func (d *User_Deleter) And_CreatedTime_LT_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And created_time < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_CreatedTime_LE_Filtering}

func (d *User_Deleter) And_CreatedTime_LE_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And created_time <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_CreatedTime_GT_Filtering}

func (d *User_Deleter) And_CreatedTime_GT_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And created_time > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_CreatedTime_GE_Filtering}

func (d *User_Deleter) And_CreatedTime_GE_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And created_time >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering Or = Or Or_CreatedTime_Eq_Filtering}

func (d *User_Deleter) Or_CreatedTime_Eq_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or created_time = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_CreatedTime_LT_Filtering}

func (d *User_Deleter) Or_CreatedTime_LT_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or created_time < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_CreatedTime_LE_Filtering}

func (d *User_Deleter) Or_CreatedTime_LE_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or created_time <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_CreatedTime_GT_Filtering}

func (d *User_Deleter) Or_CreatedTime_GT_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or created_time > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_CreatedTime_GE_Filtering}

func (d *User_Deleter) Or_CreatedTime_GE_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or created_time >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  FullName_Eq_FILTERING}

func (d *User_Deleter) FullName_Eq_FILTERING(val string) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " full_name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_FullName_Eq_FILTERING}

func (d *User_Deleter) And_FullName_Eq_FILTERING(val string) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And full_name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_FullName_Eq_FILTERING}

func (d *User_Deleter) Or_FullName_Eq_FILTERING(val string) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or full_name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq  =  UserId_Eq}

func (d *User_Deleter) UserId_Eq(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  UserId_LT_Filtering}

func (d *User_Deleter) UserId_LT_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  UserId_LE_Filtering}

func (d *User_Deleter) UserId_LE_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  UserId_GT_Filtering}

func (d *User_Deleter) UserId_GT_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  UserId_GE_Filtering}

func (d *User_Deleter) UserId_GE_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq And = And And_UserId_Eq}

func (d *User_Deleter) And_UserId_Eq(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_UserId_LT_Filtering}

func (d *User_Deleter) And_UserId_LT_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_UserId_LE_Filtering}

func (d *User_Deleter) And_UserId_LE_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_UserId_GT_Filtering}

func (d *User_Deleter) And_UserId_GT_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_UserId_GE_Filtering}

func (d *User_Deleter) And_UserId_GE_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq Or = Or Or_UserId_Eq}

func (d *User_Deleter) Or_UserId_Eq(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_UserId_LT_Filtering}

func (d *User_Deleter) Or_UserId_LT_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_UserId_LE_Filtering}

func (d *User_Deleter) Or_UserId_LE_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_UserId_GT_Filtering}

func (d *User_Deleter) Or_UserId_GT_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_UserId_GE_Filtering}

func (d *User_Deleter) Or_UserId_GE_Filtering(val int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  UserName_Eq_FILTERING}

func (d *User_Deleter) UserName_Eq_FILTERING(val string) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_UserName_Eq_FILTERING}

func (d *User_Deleter) And_UserName_Eq_FILTERING(val string) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_UserName_Eq_FILTERING}

func (d *User_Deleter) Or_UserName_Eq_FILTERING(val string) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering  =  CreatedTime_Eq_Filtering}

func (d *User_Updater) CreatedTime_Eq_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " created_time = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  CreatedTime_LT_Filtering}

func (d *User_Updater) CreatedTime_LT_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " created_time < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  CreatedTime_LE_Filtering}

func (d *User_Updater) CreatedTime_LE_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " created_time <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  CreatedTime_GT_Filtering}

func (d *User_Updater) CreatedTime_GT_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " created_time > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  CreatedTime_GE_Filtering}

func (d *User_Updater) CreatedTime_GE_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " created_time >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering And = And And_CreatedTime_Eq_Filtering}

func (d *User_Updater) And_CreatedTime_Eq_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And created_time = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_CreatedTime_LT_Filtering}

func (d *User_Updater) And_CreatedTime_LT_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And created_time < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_CreatedTime_LE_Filtering}

func (d *User_Updater) And_CreatedTime_LE_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And created_time <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_CreatedTime_GT_Filtering}

func (d *User_Updater) And_CreatedTime_GT_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And created_time > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_CreatedTime_GE_Filtering}

func (d *User_Updater) And_CreatedTime_GE_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And created_time >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering Or = Or Or_CreatedTime_Eq_Filtering}

func (d *User_Updater) Or_CreatedTime_Eq_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or created_time = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_CreatedTime_LT_Filtering}

func (d *User_Updater) Or_CreatedTime_LT_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or created_time < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_CreatedTime_LE_Filtering}

func (d *User_Updater) Or_CreatedTime_LE_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or created_time <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_CreatedTime_GT_Filtering}

func (d *User_Updater) Or_CreatedTime_GT_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or created_time > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_CreatedTime_GE_Filtering}

func (d *User_Updater) Or_CreatedTime_GE_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or created_time >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  FullName_Eq_FILTERING}

func (d *User_Updater) FullName_Eq_FILTERING(val string) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " full_name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_FullName_Eq_FILTERING}

func (d *User_Updater) And_FullName_Eq_FILTERING(val string) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And full_name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_FullName_Eq_FILTERING}

func (d *User_Updater) Or_FullName_Eq_FILTERING(val string) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or full_name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq  =  UserId_Eq}

func (d *User_Updater) UserId_Eq(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  UserId_LT_Filtering}

func (d *User_Updater) UserId_LT_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  UserId_LE_Filtering}

func (d *User_Updater) UserId_LE_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  UserId_GT_Filtering}

func (d *User_Updater) UserId_GT_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  UserId_GE_Filtering}

func (d *User_Updater) UserId_GE_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq And = And And_UserId_Eq}

func (d *User_Updater) And_UserId_Eq(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_UserId_LT_Filtering}

func (d *User_Updater) And_UserId_LT_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_UserId_LE_Filtering}

func (d *User_Updater) And_UserId_LE_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_UserId_GT_Filtering}

func (d *User_Updater) And_UserId_GT_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_UserId_GE_Filtering}

func (d *User_Updater) And_UserId_GE_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq Or = Or Or_UserId_Eq}

func (d *User_Updater) Or_UserId_Eq(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_UserId_LT_Filtering}

func (d *User_Updater) Or_UserId_LT_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_UserId_LE_Filtering}

func (d *User_Updater) Or_UserId_LE_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_UserId_GT_Filtering}

func (d *User_Updater) Or_UserId_GT_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_UserId_GE_Filtering}

func (d *User_Updater) Or_UserId_GE_Filtering(val int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  UserName_Eq_FILTERING}

func (d *User_Updater) UserName_Eq_FILTERING(val string) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_UserName_Eq_FILTERING}

func (d *User_Updater) And_UserName_Eq_FILTERING(val string) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_UserName_Eq_FILTERING}

func (d *User_Updater) Or_UserName_Eq_FILTERING(val string) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering  =  CreatedTime_Eq_Filtering}

func (d *User_Selector) CreatedTime_Eq_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " created_time = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  CreatedTime_LT_Filtering}

func (d *User_Selector) CreatedTime_LT_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " created_time < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  CreatedTime_LE_Filtering}

func (d *User_Selector) CreatedTime_LE_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " created_time <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  CreatedTime_GT_Filtering}

func (d *User_Selector) CreatedTime_GT_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " created_time > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  CreatedTime_GE_Filtering}

func (d *User_Selector) CreatedTime_GE_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " created_time >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering And = And And_CreatedTime_Eq_Filtering}

func (d *User_Selector) And_CreatedTime_Eq_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And created_time = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_CreatedTime_LT_Filtering}

func (d *User_Selector) And_CreatedTime_LT_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And created_time < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_CreatedTime_LE_Filtering}

func (d *User_Selector) And_CreatedTime_LE_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And created_time <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_CreatedTime_GT_Filtering}

func (d *User_Selector) And_CreatedTime_GT_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And created_time > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_CreatedTime_GE_Filtering}

func (d *User_Selector) And_CreatedTime_GE_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And created_time >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering Or = Or Or_CreatedTime_Eq_Filtering}

func (d *User_Selector) Or_CreatedTime_Eq_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or created_time = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_CreatedTime_LT_Filtering}

func (d *User_Selector) Or_CreatedTime_LT_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or created_time < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_CreatedTime_LE_Filtering}

func (d *User_Selector) Or_CreatedTime_LE_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or created_time <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_CreatedTime_GT_Filtering}

func (d *User_Selector) Or_CreatedTime_GT_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or created_time > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_CreatedTime_GE_Filtering}

func (d *User_Selector) Or_CreatedTime_GE_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or created_time >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  FullName_Eq_FILTERING}

func (d *User_Selector) FullName_Eq_FILTERING(val string) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " full_name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_FullName_Eq_FILTERING}

func (d *User_Selector) And_FullName_Eq_FILTERING(val string) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And full_name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_FullName_Eq_FILTERING}

func (d *User_Selector) Or_FullName_Eq_FILTERING(val string) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or full_name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq  =  UserId_Eq}

func (d *User_Selector) UserId_Eq(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  UserId_LT_Filtering}

func (d *User_Selector) UserId_LT_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  UserId_LE_Filtering}

func (d *User_Selector) UserId_LE_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  UserId_GT_Filtering}

func (d *User_Selector) UserId_GT_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  UserId_GE_Filtering}

func (d *User_Selector) UserId_GE_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq And = And And_UserId_Eq}

func (d *User_Selector) And_UserId_Eq(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_UserId_LT_Filtering}

func (d *User_Selector) And_UserId_LT_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_UserId_LE_Filtering}

func (d *User_Selector) And_UserId_LE_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_UserId_GT_Filtering}

func (d *User_Selector) And_UserId_GT_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_UserId_GE_Filtering}

func (d *User_Selector) And_UserId_GE_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq Or = Or Or_UserId_Eq}

func (d *User_Selector) Or_UserId_Eq(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_UserId_LT_Filtering}

func (d *User_Selector) Or_UserId_LT_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_UserId_LE_Filtering}

func (d *User_Selector) Or_UserId_LE_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_UserId_GT_Filtering}

func (d *User_Selector) Or_UserId_GT_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_UserId_GE_Filtering}

func (d *User_Selector) Or_UserId_GE_Filtering(val int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  UserName_Eq_FILTERING}

func (d *User_Selector) UserName_Eq_FILTERING(val string) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_UserName_Eq_FILTERING}

func (d *User_Selector) And_UserName_Eq_FILTERING(val string) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_UserName_Eq_FILTERING}

func (d *User_Selector) Or_UserName_Eq_FILTERING(val string) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_name = ? "
	d.wheres = append(d.wheres, w)

	return d
}

///////////////////////////////////////// ins for all //////////////////

func (d *User_Deleter) CreatedTime_In_FILTERING(val ...int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " created_time IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Deleter) And_CreatedTime_In_FILTERING(val ...int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And created_time IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Deleter) Or_CreatedTime_In_FILTERING(val ...int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or created_time IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Deleter) FullName_In_FILTERING(val ...string) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " full_name IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Deleter) And_FullName_In_FILTERING(val ...string) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And full_name IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Deleter) Or_FullName_In_FILTERING(val ...string) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or full_name IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Deleter) UserId_In(val ...int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " user_id IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Deleter) And_UserId_In(val ...int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And user_id IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Deleter) Or_UserId_In(val ...int) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or user_id IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Deleter) UserName_In_FILTERING(val ...string) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " user_name IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Deleter) And_UserName_In_FILTERING(val ...string) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And user_name IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Deleter) Or_UserName_In_FILTERING(val ...string) *User_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or user_name IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Updater) CreatedTime_In_FILTERING(val ...int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " created_time IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Updater) And_CreatedTime_In_FILTERING(val ...int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And created_time IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Updater) Or_CreatedTime_In_FILTERING(val ...int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or created_time IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Updater) FullName_In_FILTERING(val ...string) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " full_name IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Updater) And_FullName_In_FILTERING(val ...string) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And full_name IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Updater) Or_FullName_In_FILTERING(val ...string) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or full_name IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Updater) UserId_In(val ...int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " user_id IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Updater) And_UserId_In(val ...int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And user_id IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Updater) Or_UserId_In(val ...int) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or user_id IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Updater) UserName_In_FILTERING(val ...string) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " user_name IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Updater) And_UserName_In_FILTERING(val ...string) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And user_name IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Updater) Or_UserName_In_FILTERING(val ...string) *User_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or user_name IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Selector) CreatedTime_In_FILTERING(val ...int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " created_time IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Selector) And_CreatedTime_In_FILTERING(val ...int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And created_time IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Selector) Or_CreatedTime_In_FILTERING(val ...int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or created_time IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Selector) FullName_In_FILTERING(val ...string) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " full_name IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Selector) And_FullName_In_FILTERING(val ...string) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And full_name IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Selector) Or_FullName_In_FILTERING(val ...string) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or full_name IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Selector) UserId_In(val ...int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " user_id IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Selector) And_UserId_In(val ...int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And user_id IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Selector) Or_UserId_In(val ...int) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or user_id IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Selector) UserName_In_FILTERING(val ...string) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " user_name IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Selector) And_UserName_In_FILTERING(val ...string) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And user_name IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *User_Selector) Or_UserName_In_FILTERING(val ...string) *User_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or user_name IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

/////////////////////////////////////// End of Ins //////////////////////
///////////////////////////// start of where cluases

/////////////////////////////////////// Start of select //////////////////
func (u *User_Selector) _toSql() (string, []interface{}) {

	sqlWheres, whereArgs := whereClusesToSql(u.wheres, "")
	selectCols := "*"
	if len(u.selectCol) > 0 {
		selectCols = strings.Join(u.selectCol, ", ")
	}
	sqlstr := "SELECT " + selectCols + " FROM twitter.user"

	if len(strings.Trim(sqlWheres, " ")) > 0 { //2 for safty
		sqlstr += " WHERE " + sqlWheres
	}

	if len(u.orderBy) > 0 {
		orders := strings.Join(u.orderBy, ", ")
		sqlstr += " ORDER BY " + orders
	}

	if u.limit != 0 {
		sqlstr += " LIMIT " + strconv.Itoa(u.limit)
	}
	if u.allowFilter {
		sqlstr += "  ALLOW FILTERING"
	}

	return sqlstr, whereArgs
}

func (u *User_Selector) GetRow(session *gocql.Session) (*User, error) {
	var err error

	u.limit = 1
	sqlstr, whereArgs := u._toSql()

	if LogTableCqlReq.User {
		XCLog(sqlstr, whereArgs)
	}

	query := session.Query(sqlstr, whereArgs...)
	var row *User
	//by Sqlx
	// err = gocqlx.Get(row ,query)
	rows, err := User_Iter(query.Iter(), 1)
	if err != nil {
		if LogTableCqlReq.User {
			XCLogErr(err)
		}
		return nil, err
	}
	if len(rows) == 0 {
		return nil, errors.New("empty rows")
	} else {
		row = rows[0]
	}

	row._exists = true

	//OnUser_LoadOne(row)

	return row, nil
}

func (u *User_Selector) GetRows(session *gocql.Session) ([]*User, error) {
	var err error

	sqlstr, whereArgs := u._toSql()

	if LogTableCqlReq.User {
		XCLog(sqlstr, whereArgs)
	}

	query := session.Query(sqlstr, whereArgs...)

	rows, err := User_Iter(query.Iter(), -1)
	if err != nil {
		if LogTableCqlReq.User {
			XCLogErr(err)
		}
		return rows, err
	}

	for i := 0; i < len(rows); i++ {
		rows[i]._exists = true
	}

	// OnUser_LoadMany(rows)

	return rows, nil
}

func (u *User_Updater) Update(session *gocql.Session) error {
	var err error

	var updateArgs []interface{}
	var sqlUpdateArr []string
	for up, newVal := range u.updates {
		sqlUpdateArr = append(sqlUpdateArr, up)
		updateArgs = append(updateArgs, newVal)
	}
	sqlUpdate := strings.Join(sqlUpdateArr, ",")

	sqlWheres, whereArgs := whereClusesToSql(u.wheres, "")

	var allArgs []interface{}
	allArgs = append(allArgs, updateArgs...)
	allArgs = append(allArgs, whereArgs...)

	sqlstr := `UPDATE twitter.user SET ` + sqlUpdate

	if len(strings.Trim(sqlWheres, " ")) > 0 {
		sqlstr += " WHERE " + sqlWheres
	}
	if LogTableCqlReq.User {
		XCLog(sqlstr, allArgs)
	}
	err = session.Query(sqlstr, allArgs...).Exec()
	if err != nil {
		XCLogErr(err)
		return err
	}

	return nil
}

func (d *User_Deleter) Delete(session *gocql.Session) error {
	var err error

	var wheresArr []string
	var args []interface{}

	var delCols string
	if len(d.deleteCol) > 0 {
		delCols = strings.Join(d.deleteCol, ",")
	}

	for _, w := range d.wheres {
		wheresArr = append(wheresArr, w.condition)
		args = append(args, w.args...)
	}
	wheresStr := strings.Join(wheresArr, "")

	sqlstr := "DELETE" + delCols + " FROM twitter.user WHERE " + wheresStr

	// run query
	if LogTableCqlReq.User {
		XCLog(sqlstr, args)
	}
	err = session.Query(sqlstr, args...).Exec()
	if err != nil {
		XCLogErr(err)
		return err
	}

	return nil
}

/*
func MassInsert_User(rows []*User, session *gocql.Session) error {
    if len(rows) == 0 {
        return errors.New("rows slice should not be empty - inserted nothing")
    }
    var err error
    ln := len(rows)
    insVals := sqlManyDollars( 4 ,len(rows),true)

    sqlstr := "INSERT INTO twitter.user (" +
       " created_time,full_name,user_id,user_name " +
        ") VALUES " + insVals

    // run query
    vals := make([]interface{}, 0, ln*5) //5 fields

    for _, row := range rows {
    		vals = append(vals, row.CreatedTime)
    		vals = append(vals, row.FullName)
    		vals = append(vals, row.UserId)
    		vals = append(vals, row.UserName)
    }

    if LogTableCqlReq.User {
        XCLog(" MassInsert len = ", ln, sqlstr ,vals)
    }
    err = session.Query(sqlstr, vals...).Exec()
    if err != nil {
        XCLogErr(err)
        return err
    }

    return nil
}
*/
///////

func (r *User) Save(session *gocql.Session) error {
	var cols []string
	var q []string
	var vals []interface{}

	if r.CreatedTime != 0 {
		cols = append(cols, "created_time")
		q = append(q, "?")
		vals = append(vals, r.CreatedTime)
	}

	if r.FullName != "" {
		cols = append(cols, "full_name")
		q = append(q, "?")
		vals = append(vals, r.FullName)
	}

	if r.UserId != 0 {
		cols = append(cols, "user_id")
		q = append(q, "?")
		vals = append(vals, r.UserId)
	}

	if r.UserName != "" {
		cols = append(cols, "user_name")
		q = append(q, "?")
		vals = append(vals, r.UserName)
	}

	if len(cols) == 0 {
		return errors.New("can not insert empty row.")
	}

	colOut := strings.Join(cols, ",")
	qOut := strings.Join(q, ",")
	cql := "insert into twitter.user (" + colOut + ") values (" + qOut + ") "

	if LogTableCqlReq.User {
		XCLog(cql, vals)
	}
	err := session.Query(cql, vals...).Exec()
	if err != nil {
		if LogTableCqlReq.User {
			XCLogErr(err)
		}
	}
	r._exists = true
	return err
}

func (r *User) SaveBatch(batch *gocql.Session) error {
	var cols []string
	var q []string
	var vals []interface{}

	if r.CreatedTime != 0 {
		cols = append(cols, "created_time")
		q = append(q, "?")
		vals = append(vals, r.CreatedTime)
	}

	if r.FullName != "" {
		cols = append(cols, "full_name")
		q = append(q, "?")
		vals = append(vals, r.FullName)
	}

	if r.UserId != 0 {
		cols = append(cols, "user_id")
		q = append(q, "?")
		vals = append(vals, r.UserId)
	}

	if r.UserName != "" {
		cols = append(cols, "user_name")
		q = append(q, "?")
		vals = append(vals, r.UserName)
	}

	if len(cols) == 0 {
		return errors.New("can not insert empty row.")
	}

	colOut := strings.Join(cols, ",")
	qOut := strings.Join(q, ",")
	cql := "insert into twitter.user (" + colOut + ") values (" + qOut + ") "

	if LogTableCqlReq.User {
		XCLog("(in batch)", cql, vals)
	}
	err := batch.Query(cql, vals...).Exec()
	if err != nil {
		if LogTableCqlReq.User {
			XCLogErr(err)
		}
	}
	batch.Query(cql, vals...)

	return err
}

func (r *User) Delete(session *gocql.Session) error {
	var err error
	del := NewUser_Deleter()

	del.UserId_Eq(r.UserId)

	err = del.Delete(session)
	if err != nil {
		return err
	}
	r._exists = false
	return nil
}

func User_Iter(iter *gocql.Iter, limit int) ([]*User, error) {
	var rows []*User
	if limit < 1 {
		limit = 1e6
	}

	for i := 0; i < limit; i++ {
		m := make(map[string]interface{}, 10)
		row := &User{}
		if iter.MapScan(m) {

			if val, ok := m["created_time"]; ok {
				row.CreatedTime = int(val.(int64))
				//row.CreatedTime = val.(int)
			}

			if val, ok := m["full_name"]; ok {
				row.FullName = string(val.(string))
				//row.FullName = val.(string)
			}

			if val, ok := m["user_id"]; ok {
				row.UserId = int(val.(int))
				//row.UserId = val.(int)
			}

			if val, ok := m["user_name"]; ok {
				row.UserName = string(val.(string))
				//row.UserName = val.(string)
			}

			rows = append(rows, row)
		} else {
			break
		}
	}
	err := iter.Close()

	return rows, err
}
