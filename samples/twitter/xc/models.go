package xc

type Twitt struct {
	Body       string // body  regular
	CreateTime int    // create_time  regular
	TwiitId    string // twiit_id  clustering
	UserId     int    // user_id  partition_key

	_exists, _deleted bool
}

/*
:= &xc.Twitt {
	Body: "",
	CreateTime: 0,
	TwiitId: "",
	UserId: 0,
*/

type User struct {
	CreatedTime int    // created_time  regular
	FullName    string // full_name  regular
	UserId      int    // user_id  partition_key
	UserName    string // user_name  regular

	_exists, _deleted bool
}

/*
:= &xc.User {
	CreatedTime: 0,
	FullName: "",
	UserId: 0,
	UserName: "",
*/

// logs tables
type LogTableCql struct {
	Twitt bool
	User  bool
}

var LogTableCqlReq = LogTableCql{
	Twitt: true,
	User:  true,
}
