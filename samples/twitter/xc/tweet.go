package xc

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gocql/gocql"
)

////////////////////////////////////////// Query seletor updater and deleter /////////////////////////

func (a *Tweet) Exists() bool {
	return a._exists
}

func (a *Tweet) Deleted() bool {
	return a._deleted
}

type Tweet_Selector struct {
	wheres      []whereClause
	selectCol   []string
	orderBy     []string //" order by id desc //for ints
	limit       int
	allowFilter bool
}

type Tweet_Updater struct {
	wheres  []whereClause
	updates map[string]interface{}
}

type Tweet_Deleter struct {
	wheres    []whereClause
	deleteCol []string
}

//////////////////// just Selector
func (u *Tweet_Selector) Limit(limit int) *Tweet_Selector {
	u.limit = limit
	return u
}

func (u *Tweet_Selector) AllowFiltering() *Tweet_Selector {
	u.allowFilter = true
	return u
}

func NewTweet_Selector() *Tweet_Selector {
	u := Tweet_Selector{}
	return &u
}

func NewTweet_Updater() *Tweet_Updater {
	u := Tweet_Updater{}
	u.updates = make(map[string]interface{})
	return &u
}

func NewTweet_Deleter() *Tweet_Deleter {
	u := Tweet_Deleter{}
	return &u
}

//each select columns

func (u *Tweet_Selector) Select_Body() *Tweet_Selector {
	u.selectCol = append(u.selectCol, "body")
	return u
}

//each column orders //just ints
func (u *Tweet_Selector) OrderBy_Body_Desc() *Tweet_Selector {
	u.orderBy = append(u.orderBy, " body DESC")
	return u
}

func (u *Tweet_Selector) OrderBy_Body_Asc() *Tweet_Selector {
	u.orderBy = append(u.orderBy, " body ASC")
	return u
}

func (u *Tweet_Selector) Select_CreateTime() *Tweet_Selector {
	u.selectCol = append(u.selectCol, "create_time")
	return u
}

//each column orders //just ints
func (u *Tweet_Selector) OrderBy_CreateTime_Desc() *Tweet_Selector {
	u.orderBy = append(u.orderBy, " create_time DESC")
	return u
}

func (u *Tweet_Selector) OrderBy_CreateTime_Asc() *Tweet_Selector {
	u.orderBy = append(u.orderBy, " create_time ASC")
	return u
}

func (u *Tweet_Selector) Select_TweetId() *Tweet_Selector {
	u.selectCol = append(u.selectCol, "tweet_id")
	return u
}

//each column orders //just ints
func (u *Tweet_Selector) OrderBy_TweetId_Desc() *Tweet_Selector {
	u.orderBy = append(u.orderBy, " tweet_id DESC")
	return u
}

func (u *Tweet_Selector) OrderBy_TweetId_Asc() *Tweet_Selector {
	u.orderBy = append(u.orderBy, " tweet_id ASC")
	return u
}

func (u *Tweet_Selector) Select_UserId() *Tweet_Selector {
	u.selectCol = append(u.selectCol, "user_id")
	return u
}

//each column orders //just ints
func (u *Tweet_Selector) OrderBy_UserId_Desc() *Tweet_Selector {
	u.orderBy = append(u.orderBy, " user_id DESC")
	return u
}

func (u *Tweet_Selector) OrderBy_UserId_Asc() *Tweet_Selector {
	u.orderBy = append(u.orderBy, " user_id ASC")
	return u
}

//////////////////// just Deleter
//each column delete

func (u *Tweet_Deleter) Delete_Body() *Tweet_Deleter {
	u.deleteCol = append(u.deleteCol, "body")
	return u
}

func (u *Tweet_Deleter) Delete_CreateTime() *Tweet_Deleter {
	u.deleteCol = append(u.deleteCol, "create_time")
	return u
}

func (u *Tweet_Deleter) Delete_TweetId() *Tweet_Deleter {
	u.deleteCol = append(u.deleteCol, "tweet_id")
	return u
}

func (u *Tweet_Deleter) Delete_UserId() *Tweet_Deleter {
	u.deleteCol = append(u.deleteCol, "user_id")
	return u
}

//////////////////// End of just Deleter

//////////////////// just Updater
//each column delete

func (u *Tweet_Updater) Body(newVal string) *Tweet_Updater {
	u.updates["body = ? "] = newVal
	return u
}

func (u *Tweet_Updater) CreateTime(newVal int) *Tweet_Updater {
	u.updates["create_time = ? "] = newVal
	return u
}

func (u *Tweet_Updater) TweetId(newVal string) *Tweet_Updater {
	u.updates["tweet_id = ? "] = newVal
	return u
}

func (u *Tweet_Updater) UserId(newVal int) *Tweet_Updater {
	u.updates["user_id = ? "] = newVal
	return u
}

//////////////////// End just Updater

//{_Eq_FILTERING  =  Body_Eq_FILTERING}

func (d *Tweet_Deleter) Body_Eq_FILTERING(val string) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " body = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_Body_Eq_FILTERING}

func (d *Tweet_Deleter) And_Body_Eq_FILTERING(val string) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And body = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_Body_Eq_FILTERING}

func (d *Tweet_Deleter) Or_Body_Eq_FILTERING(val string) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or body = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering  =  CreateTime_Eq_Filtering}

func (d *Tweet_Deleter) CreateTime_Eq_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " create_time = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  CreateTime_LT_Filtering}

func (d *Tweet_Deleter) CreateTime_LT_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " create_time < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  CreateTime_LE_Filtering}

func (d *Tweet_Deleter) CreateTime_LE_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " create_time <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  CreateTime_GT_Filtering}

func (d *Tweet_Deleter) CreateTime_GT_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " create_time > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  CreateTime_GE_Filtering}

func (d *Tweet_Deleter) CreateTime_GE_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " create_time >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering And = And And_CreateTime_Eq_Filtering}

func (d *Tweet_Deleter) And_CreateTime_Eq_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And create_time = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_CreateTime_LT_Filtering}

func (d *Tweet_Deleter) And_CreateTime_LT_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And create_time < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_CreateTime_LE_Filtering}

func (d *Tweet_Deleter) And_CreateTime_LE_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And create_time <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_CreateTime_GT_Filtering}

func (d *Tweet_Deleter) And_CreateTime_GT_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And create_time > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_CreateTime_GE_Filtering}

func (d *Tweet_Deleter) And_CreateTime_GE_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And create_time >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering Or = Or Or_CreateTime_Eq_Filtering}

func (d *Tweet_Deleter) Or_CreateTime_Eq_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or create_time = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_CreateTime_LT_Filtering}

func (d *Tweet_Deleter) Or_CreateTime_LT_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or create_time < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_CreateTime_LE_Filtering}

func (d *Tweet_Deleter) Or_CreateTime_LE_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or create_time <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_CreateTime_GT_Filtering}

func (d *Tweet_Deleter) Or_CreateTime_GT_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or create_time > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_CreateTime_GE_Filtering}

func (d *Tweet_Deleter) Or_CreateTime_GE_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or create_time >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq  =  TweetId_Eq}

func (d *Tweet_Deleter) TweetId_Eq(val string) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " tweet_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq And = And And_TweetId_Eq}

func (d *Tweet_Deleter) And_TweetId_Eq(val string) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And tweet_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq Or = Or Or_TweetId_Eq}

func (d *Tweet_Deleter) Or_TweetId_Eq(val string) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or tweet_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq  =  UserId_Eq}

func (d *Tweet_Deleter) UserId_Eq(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  UserId_LT_Filtering}

func (d *Tweet_Deleter) UserId_LT_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  UserId_LE_Filtering}

func (d *Tweet_Deleter) UserId_LE_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  UserId_GT_Filtering}

func (d *Tweet_Deleter) UserId_GT_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  UserId_GE_Filtering}

func (d *Tweet_Deleter) UserId_GE_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq And = And And_UserId_Eq}

func (d *Tweet_Deleter) And_UserId_Eq(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_UserId_LT_Filtering}

func (d *Tweet_Deleter) And_UserId_LT_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_UserId_LE_Filtering}

func (d *Tweet_Deleter) And_UserId_LE_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_UserId_GT_Filtering}

func (d *Tweet_Deleter) And_UserId_GT_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_UserId_GE_Filtering}

func (d *Tweet_Deleter) And_UserId_GE_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq Or = Or Or_UserId_Eq}

func (d *Tweet_Deleter) Or_UserId_Eq(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_UserId_LT_Filtering}

func (d *Tweet_Deleter) Or_UserId_LT_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_UserId_LE_Filtering}

func (d *Tweet_Deleter) Or_UserId_LE_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_UserId_GT_Filtering}

func (d *Tweet_Deleter) Or_UserId_GT_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_UserId_GE_Filtering}

func (d *Tweet_Deleter) Or_UserId_GE_Filtering(val int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  Body_Eq_FILTERING}

func (d *Tweet_Updater) Body_Eq_FILTERING(val string) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " body = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_Body_Eq_FILTERING}

func (d *Tweet_Updater) And_Body_Eq_FILTERING(val string) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And body = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_Body_Eq_FILTERING}

func (d *Tweet_Updater) Or_Body_Eq_FILTERING(val string) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or body = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering  =  CreateTime_Eq_Filtering}

func (d *Tweet_Updater) CreateTime_Eq_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " create_time = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  CreateTime_LT_Filtering}

func (d *Tweet_Updater) CreateTime_LT_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " create_time < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  CreateTime_LE_Filtering}

func (d *Tweet_Updater) CreateTime_LE_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " create_time <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  CreateTime_GT_Filtering}

func (d *Tweet_Updater) CreateTime_GT_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " create_time > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  CreateTime_GE_Filtering}

func (d *Tweet_Updater) CreateTime_GE_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " create_time >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering And = And And_CreateTime_Eq_Filtering}

func (d *Tweet_Updater) And_CreateTime_Eq_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And create_time = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_CreateTime_LT_Filtering}

func (d *Tweet_Updater) And_CreateTime_LT_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And create_time < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_CreateTime_LE_Filtering}

func (d *Tweet_Updater) And_CreateTime_LE_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And create_time <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_CreateTime_GT_Filtering}

func (d *Tweet_Updater) And_CreateTime_GT_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And create_time > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_CreateTime_GE_Filtering}

func (d *Tweet_Updater) And_CreateTime_GE_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And create_time >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering Or = Or Or_CreateTime_Eq_Filtering}

func (d *Tweet_Updater) Or_CreateTime_Eq_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or create_time = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_CreateTime_LT_Filtering}

func (d *Tweet_Updater) Or_CreateTime_LT_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or create_time < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_CreateTime_LE_Filtering}

func (d *Tweet_Updater) Or_CreateTime_LE_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or create_time <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_CreateTime_GT_Filtering}

func (d *Tweet_Updater) Or_CreateTime_GT_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or create_time > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_CreateTime_GE_Filtering}

func (d *Tweet_Updater) Or_CreateTime_GE_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or create_time >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq  =  TweetId_Eq}

func (d *Tweet_Updater) TweetId_Eq(val string) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " tweet_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq And = And And_TweetId_Eq}

func (d *Tweet_Updater) And_TweetId_Eq(val string) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And tweet_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq Or = Or Or_TweetId_Eq}

func (d *Tweet_Updater) Or_TweetId_Eq(val string) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or tweet_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq  =  UserId_Eq}

func (d *Tweet_Updater) UserId_Eq(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  UserId_LT_Filtering}

func (d *Tweet_Updater) UserId_LT_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  UserId_LE_Filtering}

func (d *Tweet_Updater) UserId_LE_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  UserId_GT_Filtering}

func (d *Tweet_Updater) UserId_GT_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  UserId_GE_Filtering}

func (d *Tweet_Updater) UserId_GE_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq And = And And_UserId_Eq}

func (d *Tweet_Updater) And_UserId_Eq(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_UserId_LT_Filtering}

func (d *Tweet_Updater) And_UserId_LT_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_UserId_LE_Filtering}

func (d *Tweet_Updater) And_UserId_LE_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_UserId_GT_Filtering}

func (d *Tweet_Updater) And_UserId_GT_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_UserId_GE_Filtering}

func (d *Tweet_Updater) And_UserId_GE_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq Or = Or Or_UserId_Eq}

func (d *Tweet_Updater) Or_UserId_Eq(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_UserId_LT_Filtering}

func (d *Tweet_Updater) Or_UserId_LT_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_UserId_LE_Filtering}

func (d *Tweet_Updater) Or_UserId_LE_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_UserId_GT_Filtering}

func (d *Tweet_Updater) Or_UserId_GT_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_UserId_GE_Filtering}

func (d *Tweet_Updater) Or_UserId_GE_Filtering(val int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING  =  Body_Eq_FILTERING}

func (d *Tweet_Selector) Body_Eq_FILTERING(val string) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " body = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING And = And And_Body_Eq_FILTERING}

func (d *Tweet_Selector) And_Body_Eq_FILTERING(val string) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And body = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_FILTERING Or = Or Or_Body_Eq_FILTERING}

func (d *Tweet_Selector) Or_Body_Eq_FILTERING(val string) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or body = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering  =  CreateTime_Eq_Filtering}

func (d *Tweet_Selector) CreateTime_Eq_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " create_time = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  CreateTime_LT_Filtering}

func (d *Tweet_Selector) CreateTime_LT_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " create_time < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  CreateTime_LE_Filtering}

func (d *Tweet_Selector) CreateTime_LE_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " create_time <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  CreateTime_GT_Filtering}

func (d *Tweet_Selector) CreateTime_GT_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " create_time > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  CreateTime_GE_Filtering}

func (d *Tweet_Selector) CreateTime_GE_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " create_time >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering And = And And_CreateTime_Eq_Filtering}

func (d *Tweet_Selector) And_CreateTime_Eq_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And create_time = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_CreateTime_LT_Filtering}

func (d *Tweet_Selector) And_CreateTime_LT_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And create_time < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_CreateTime_LE_Filtering}

func (d *Tweet_Selector) And_CreateTime_LE_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And create_time <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_CreateTime_GT_Filtering}

func (d *Tweet_Selector) And_CreateTime_GT_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And create_time > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_CreateTime_GE_Filtering}

func (d *Tweet_Selector) And_CreateTime_GE_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And create_time >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq_Filtering Or = Or Or_CreateTime_Eq_Filtering}

func (d *Tweet_Selector) Or_CreateTime_Eq_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or create_time = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_CreateTime_LT_Filtering}

func (d *Tweet_Selector) Or_CreateTime_LT_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or create_time < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_CreateTime_LE_Filtering}

func (d *Tweet_Selector) Or_CreateTime_LE_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or create_time <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_CreateTime_GT_Filtering}

func (d *Tweet_Selector) Or_CreateTime_GT_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or create_time > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_CreateTime_GE_Filtering}

func (d *Tweet_Selector) Or_CreateTime_GE_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or create_time >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq  =  TweetId_Eq}

func (d *Tweet_Selector) TweetId_Eq(val string) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " tweet_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq And = And And_TweetId_Eq}

func (d *Tweet_Selector) And_TweetId_Eq(val string) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And tweet_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq Or = Or Or_TweetId_Eq}

func (d *Tweet_Selector) Or_TweetId_Eq(val string) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or tweet_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq  =  UserId_Eq}

func (d *Tweet_Selector) UserId_Eq(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering  <  UserId_LT_Filtering}

func (d *Tweet_Selector) UserId_LT_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering  <=  UserId_LE_Filtering}

func (d *Tweet_Selector) UserId_LE_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering  >  UserId_GT_Filtering}

func (d *Tweet_Selector) UserId_GT_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering  >=  UserId_GE_Filtering}

func (d *Tweet_Selector) UserId_GE_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = " user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq And = And And_UserId_Eq}

func (d *Tweet_Selector) And_UserId_Eq(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering And < And And_UserId_LT_Filtering}

func (d *Tweet_Selector) And_UserId_LT_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering And <= And And_UserId_LE_Filtering}

func (d *Tweet_Selector) And_UserId_LE_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering And > And And_UserId_GT_Filtering}

func (d *Tweet_Selector) And_UserId_GT_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering And >= And And_UserId_GE_Filtering}

func (d *Tweet_Selector) And_UserId_GE_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "And user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_Eq Or = Or Or_UserId_Eq}

func (d *Tweet_Selector) Or_UserId_Eq(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id = ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LT_Filtering Or < Or Or_UserId_LT_Filtering}

func (d *Tweet_Selector) Or_UserId_LT_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id < ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_LE_Filtering Or <= Or Or_UserId_LE_Filtering}

func (d *Tweet_Selector) Or_UserId_LE_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id <= ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GT_Filtering Or > Or Or_UserId_GT_Filtering}

func (d *Tweet_Selector) Or_UserId_GT_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id > ? "
	d.wheres = append(d.wheres, w)

	return d
}

//{_GE_Filtering Or >= Or Or_UserId_GE_Filtering}

func (d *Tweet_Selector) Or_UserId_GE_Filtering(val int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	insWhere = append(insWhere, val)
	w.args = insWhere
	w.condition = "Or user_id >= ? "
	d.wheres = append(d.wheres, w)

	return d
}

///////////////////////////////////////// ins for all //////////////////

func (d *Tweet_Deleter) Body_In_FILTERING(val ...string) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " body IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Deleter) And_Body_In_FILTERING(val ...string) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And body IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Deleter) Or_Body_In_FILTERING(val ...string) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or body IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Deleter) CreateTime_In_FILTERING(val ...int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " create_time IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Deleter) And_CreateTime_In_FILTERING(val ...int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And create_time IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Deleter) Or_CreateTime_In_FILTERING(val ...int) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or create_time IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Deleter) TweetId_In(val ...string) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " tweet_id IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Deleter) And_TweetId_In(val ...string) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And tweet_id IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Deleter) Or_TweetId_In(val ...string) *Tweet_Deleter {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or tweet_id IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Deleter) UserId_In(val ...int) *Tweet_Deleter {
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

func (d *Tweet_Deleter) And_UserId_In(val ...int) *Tweet_Deleter {
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

func (d *Tweet_Deleter) Or_UserId_In(val ...int) *Tweet_Deleter {
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

func (d *Tweet_Updater) Body_In_FILTERING(val ...string) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " body IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Updater) And_Body_In_FILTERING(val ...string) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And body IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Updater) Or_Body_In_FILTERING(val ...string) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or body IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Updater) CreateTime_In_FILTERING(val ...int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " create_time IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Updater) And_CreateTime_In_FILTERING(val ...int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And create_time IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Updater) Or_CreateTime_In_FILTERING(val ...int) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or create_time IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Updater) TweetId_In(val ...string) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " tweet_id IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Updater) And_TweetId_In(val ...string) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And tweet_id IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Updater) Or_TweetId_In(val ...string) *Tweet_Updater {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or tweet_id IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Updater) UserId_In(val ...int) *Tweet_Updater {
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

func (d *Tweet_Updater) And_UserId_In(val ...int) *Tweet_Updater {
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

func (d *Tweet_Updater) Or_UserId_In(val ...int) *Tweet_Updater {
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

func (d *Tweet_Selector) Body_In_FILTERING(val ...string) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " body IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Selector) And_Body_In_FILTERING(val ...string) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And body IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Selector) Or_Body_In_FILTERING(val ...string) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or body IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Selector) CreateTime_In_FILTERING(val ...int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " create_time IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Selector) And_CreateTime_In_FILTERING(val ...int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And create_time IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Selector) Or_CreateTime_In_FILTERING(val ...int) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or create_time IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Selector) TweetId_In(val ...string) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = " tweet_id IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Selector) And_TweetId_In(val ...string) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "And tweet_id IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Selector) Or_TweetId_In(val ...string) *Tweet_Selector {
	w := whereClause{}
	var insWhere []interface{}
	for _, v := range val {
		insWhere = append(insWhere, v)
	}
	w.args = insWhere
	w.condition = "Or tweet_id IN (" + dbQuestionForSqlIn(len(val)) + ") "
	d.wheres = append(d.wheres, w)

	return d
}

func (d *Tweet_Selector) UserId_In(val ...int) *Tweet_Selector {
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

func (d *Tweet_Selector) And_UserId_In(val ...int) *Tweet_Selector {
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

func (d *Tweet_Selector) Or_UserId_In(val ...int) *Tweet_Selector {
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

/////////////////////////////////////// End of Ins //////////////////////
///////////////////////////// start of where cluases

/////////////////////////////////////// Start of select //////////////////
func (u *Tweet_Selector) _toSql() (string, []interface{}) {

	sqlWheres, whereArgs := whereClusesToSql(u.wheres, "")
	selectCols := "*"
	if len(u.selectCol) > 0 {
		selectCols = strings.Join(u.selectCol, ", ")
	}
	sqlstr := "SELECT " + selectCols + " FROM twitter.tweet"

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

func (u *Tweet_Selector) GetRow(session *gocql.Session) (*Tweet, error) {
	var err error

	u.limit = 1
	sqlstr, whereArgs := u._toSql()

	if LogTableCqlReq.Tweet {
		XCLog(sqlstr, whereArgs)
	}

	query := session.Query(sqlstr, whereArgs...)
	var row *Tweet
	//by Sqlx
	// err = gocqlx.Get(row ,query)
	rows, err := Tweet_Iter(query.Iter(), 1)
	if err != nil {
		if LogTableCqlReq.Tweet {
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

	//OnTweet_LoadOne(row)

	return row, nil
}

func (u *Tweet_Selector) GetRows(session *gocql.Session) ([]*Tweet, error) {
	var err error

	sqlstr, whereArgs := u._toSql()

	if LogTableCqlReq.Tweet {
		XCLog(sqlstr, whereArgs)
	}

	query := session.Query(sqlstr, whereArgs...)

	rows, err := Tweet_Iter(query.Iter(), -1)
	if err != nil {
		if LogTableCqlReq.Tweet {
			XCLogErr(err)
		}
		return rows, err
	}

	for i := 0; i < len(rows); i++ {
		rows[i]._exists = true
	}

	// OnTweet_LoadMany(rows)

	return rows, nil
}

func (u *Tweet_Updater) Update(session *gocql.Session) error {
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

	sqlstr := `UPDATE twitter.tweet SET ` + sqlUpdate

	if len(strings.Trim(sqlWheres, " ")) > 0 {
		sqlstr += " WHERE " + sqlWheres
	}
	if LogTableCqlReq.Tweet {
		XCLog(sqlstr, allArgs)
	}
	err = session.Query(sqlstr, allArgs...).Exec()
	if err != nil {
		XCLogErr(err)
		return err
	}

	return nil
}

func (d *Tweet_Deleter) Delete(session *gocql.Session) error {
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

	sqlstr := "DELETE" + delCols + " FROM twitter.tweet WHERE " + wheresStr

	// run query
	if LogTableCqlReq.Tweet {
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
func MassInsert_Tweet(rows []*Tweet, session *gocql.Session) error {
    if len(rows) == 0 {
        return errors.New("rows slice should not be empty - inserted nothing")
    }
    var err error
    ln := len(rows)
    insVals := sqlManyDollars( 4 ,len(rows),true)

    sqlstr := "INSERT INTO twitter.tweet (" +
       " body,create_time,tweet_id,user_id " +
        ") VALUES " + insVals

    // run query
    vals := make([]interface{}, 0, ln*5) //5 fields

    for _, row := range rows {
    		vals = append(vals, row.Body)
    		vals = append(vals, row.CreateTime)
    		vals = append(vals, row.TweetId)
    		vals = append(vals, row.UserId)
    }

    if LogTableCqlReq.Tweet {
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

func (r *Tweet) Save(session *gocql.Session) error {
	var cols []string
	var q []string
	var vals []interface{}

	if r.Body != "" {
		cols = append(cols, "body")
		q = append(q, "?")
		vals = append(vals, r.Body)
	}

	if r.CreateTime != 0 {
		cols = append(cols, "create_time")
		q = append(q, "?")
		vals = append(vals, r.CreateTime)
	}

	if r.TweetId != "" {
		cols = append(cols, "tweet_id")
		q = append(q, "?")
		vals = append(vals, r.TweetId)
	}

	if r.UserId != 0 {
		cols = append(cols, "user_id")
		q = append(q, "?")
		vals = append(vals, r.UserId)
	}

	if len(cols) == 0 {
		return errors.New("can not insert empty row.")
	}

	colOut := strings.Join(cols, ",")
	qOut := strings.Join(q, ",")
	cql := "insert into twitter.tweet (" + colOut + ") values (" + qOut + ") "

	if LogTableCqlReq.Tweet {
		XCLog(cql, vals)
	}
	err := session.Query(cql, vals...).Exec()
	if err != nil {
		if LogTableCqlReq.Tweet {
			XCLogErr(err)
		}
	}
	r._exists = true
	return err
}

func (r *Tweet) SaveBatch(batch *gocql.Session) error {
	var cols []string
	var q []string
	var vals []interface{}

	if r.Body != "" {
		cols = append(cols, "body")
		q = append(q, "?")
		vals = append(vals, r.Body)
	}

	if r.CreateTime != 0 {
		cols = append(cols, "create_time")
		q = append(q, "?")
		vals = append(vals, r.CreateTime)
	}

	if r.TweetId != "" {
		cols = append(cols, "tweet_id")
		q = append(q, "?")
		vals = append(vals, r.TweetId)
	}

	if r.UserId != 0 {
		cols = append(cols, "user_id")
		q = append(q, "?")
		vals = append(vals, r.UserId)
	}

	if len(cols) == 0 {
		return errors.New("can not insert empty row.")
	}

	colOut := strings.Join(cols, ",")
	qOut := strings.Join(q, ",")
	cql := "insert into twitter.tweet (" + colOut + ") values (" + qOut + ") "

	if LogTableCqlReq.Tweet {
		XCLog("(in batch)", cql, vals)
	}
	err := batch.Query(cql, vals...).Exec()
	if err != nil {
		if LogTableCqlReq.Tweet {
			XCLogErr(err)
		}
	}
	batch.Query(cql, vals...)

	return err
}

func (r *Tweet) Delete(session *gocql.Session) error {
	var err error
	del := NewTweet_Deleter()

	del.UserId_Eq(r.UserId)

	del.And_TweetId_Eq(r.TweetId)

	err = del.Delete(session)
	if err != nil {
		return err
	}
	r._exists = false
	return nil
}

func Tweet_Iter(iter *gocql.Iter, limit int) ([]*Tweet, error) {
	var rows []*Tweet
	if limit < 1 {
		limit = 1e6
	}

	for i := 0; i < limit; i++ {
		m := make(map[string]interface{}, 10)
		row := &Tweet{}
		if iter.MapScan(m) {

			if val, ok := m["body"]; ok {
				row.Body = string(val.(string))
				//row.Body = val.(string)
			}

			if val, ok := m["create_time"]; ok {
				row.CreateTime = int(val.(int))
				//row.CreateTime = val.(int)
			}

			if val, ok := m["tweet_id"]; ok {
				row.TweetId = string(val.(string))
				//row.TweetId = val.(string)
			}

			if val, ok := m["user_id"]; ok {
				row.UserId = int(val.(int64))
				//row.UserId = val.(int)
			}

			rows = append(rows, row)
		} else {
			break
		}
	}
	err := iter.Close()

	return rows, err
}
