# Cassandra walker
Cassandra walker is and ORM (Object Relation Mapper) like, for cassandra.

`cassandra-walker` is command-line tool to generate Golang code based for cassandra keyspaces (databases).

with pointing `cassandra-walker` to cassandra cluster, `cassandra-walker` find tables in each keyspaces and for each tables create golang types, and idiom go codes.

This tools makes working with Cassandra a bleaz, it works great with IDEs, catch most of bugs at compile time. Makes code base much more scalable and maintainable.

#### Demo of how you will code with the result of this tool

Create, Update, Delete for a Row:
![](https://raw.githubusercontent.com/jozn/cassandra-walker/master/art/1.gif)

Building query against Cassandra:
![](https://raw.githubusercontent.com/jozn/cassandra-walker/master/art/2.gif)

## Quickstart

Install `cassandra-walker` with:
```
go get -u github.com/jozn/cassandra-walker
```

Then point to cassandra node to genrate code for a keyspace (ex `twitter`):

```
cassandra-walker twitter
```

This will produce `xc` folder in current directory, and puts generated golang codes in this folder.

## Command Line Parameters

Use ` cassandra-walker -h` to see parameter options.

```
Usage: cassandra-walker [--host HOST] [--port PORT] [--verbose] [--dir DIR] [--package PACKAGE] [--minimize] [KEYSPACES [KEYSPACES ...]]

Positional arguments:
  KEYSPACES              cassandra keyspaces to build

Options:
  --host HOST, -c HOST   cassandra cluster address (default 127.0.0.1)
  --port PORT, -p PORT   cassandra port (default 9042)
  --verbose, -v          verbosity Log
  --dir DIR, -d DIR      output of generated codes (default './')
  --package PACKAGE      package of go
  --minimize, -m         minimize docs
  --help, -h             display this help and exit
```

## Guides
Lets see how to use this tool.
We will follow twitter sample in [sample directory](https://github.com/jozn/cassandra-walker/tree/master/samples/twitter)

Assume you have this cassandra keyspace:
```cql
CREATE KEYSPACE twitter
  WITH REPLICATION = {
   'class' : 'SimpleStrategy',
   'replication_factor' : 1
 };

CREATE TABLE twitter.twitt (
	user_id bigint,
	tweet_id varchar,
	body varchar,
	create_time int,
	PRIMARY KEY (user_id,tweet_id)
);

CREATE TABLE twitter.user (
	user_id int,
	user_name varchar,
	full_name varchar,
	created_time bigint,
	PRIMARY KEY (user_id)
);
```

Run the following command:
```
cassandra-walker twitter --host 127.0.01 --port 9042
```

This will generates codes in `xc` directory.

See the result in [godoc](https://godoc.org/github.com/jozn/cassandra-walker/samples/twitter/xc) or in [go files](https://github.com/jozn/cassandra-walker/tree/master/samples/twitter/xc).

```go
package xc

type Twitt struct {
	Body       string // body  regular
	CreateTime int    // create_time  regular
	TweetId    string // tweet_id  clustering
	UserId     int    // user_id  partition_key

	_exists, _deleted bool
}

/*
:= &xc.Twitt {
	Body: "",
	CreateTime: 0,
	TweetId: "",
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
```
Now we have in type-safe can build queries, look at this simple script:

```go
package main

import (
	"github.com/gocql/gocql"
	"github.com/jozn/cassandra-walker/samples/twitter/xc"
)

func main() {
	// create cassandra session
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "twitter"
	cluster.Consistency = gocql.One
	session, _ := cluster.CreateSession()
	defer session.Close()

	// Create
	twitt1 := xc.Twitt{
		Body:       "Hello World",
		CreateTime: 1566000000,
		TweetId:    1,
		UserId:     1,
	}

	err := twitt1.Save(session)

	// Delete one object
	twitt1.Delete(session)

	//////////////// For Selector
	twitts, err := xc.NewTwitt_Selector().UserId_Eq(1).Limit(5).GetRows(session) // returns and array of twitts ( []*twitt ,err )

	twitt, err := xc.NewTwitt_Selector().UserId_Eq(1).Limit(5).GetRows(session) // returns a single twitt ( *twitt ,err )

	//can use clustering columns too
	twitts, err = xc.NewTwitt_Selector().UserId_Eq(1).And_TweetId_In(1, 25, 68).GetRows(session)

	//can select just some columns, it will returns *[]Twitt, with just selected columns sets
	twitts, err = xc.NewTwitt_Selector().Select_UserId().Select_Body().UserId_Eq(1).And_TweetId_In(1, 25, 68).Limit(12).GetRows(session)

	//for when need to use filtering
	twitts, err = xc.NewTwitt_Selector().UserId_LT_Filtering(100).Limit(10).AllowFiltering().GetRows(session)

	//////////////// For Updater
	err = xc.NewTwitt_Updater().
		Body("new twitt text").UserId_Eq(1).And_TweetId_In(1, 2, 3).Update(session)

	//////////////// For Deleter
	err = xc.NewTwitt_Deleter().UserId_Eq(1).And_TweetId_In(1, 2, 3).Delete(session)
	err = xc.NewTwitt_Deleter().UserId_Eq(1).Delete(session)

	_ = err
	_ = twitts
	_ = twitt
}

/* log output - this is produced CQL queries to cassandra:

2018/09/18 22:35:54 CQL:  [insert into twitter.twitt (body,create_time,tweet_id,user_id) values (?,?,?,?)  [Hello World 1566000000 1 1]]
2018/09/18 22:35:54 CQL:  [DELETE FROM twitter.twitt WHERE  user_id = ? And tweet_id = ?  [1 1]]
2018/09/18 22:35:54 CQL:  [SELECT * FROM twitter.twitt WHERE  user_id = ?  LIMIT 5 [1]]
2018/09/18 22:35:54 CQL:  [SELECT * FROM twitter.twitt WHERE  user_id = ?  LIMIT 5 [1]]
2018/09/18 22:35:54 CQL:  [SELECT * FROM twitter.twitt WHERE  user_id = ? And tweet_id IN (?,?,?)  [1 1 25 68]]
2018/09/18 22:35:54 CQL:  [SELECT user_id, body FROM twitter.twitt WHERE  user_id = ? And tweet_id IN (?,?,?)  LIMIT 12 [1 1 25 68]]
2018/09/18 22:35:54 CQL:  [SELECT * FROM twitter.twitt WHERE  user_id < ?  LIMIT 10  ALLOW FILTERING [100]]
2018/09/18 22:35:54 CQL:  [UPDATE twitter.twitt SET body = ?  WHERE  user_id = ? And tweet_id IN (?,?,?)  [new twitt text 1 1 2 3]]
2018/09/18 22:35:54 CQL:  [DELETE FROM twitter.twitt WHERE  user_id = ? And tweet_id IN (?,?,?)  [1 1 2 3]]
2018/09/18 22:35:54 CQL:  [DELETE FROM twitter.twitt WHERE  user_id = ?  [1]]

*/

```

### Todos
- [x] Add twitter sample play code.
- [ ] Add docs for logging , and better docs for `*_Selector()` ,`*_Updaterr()` and `*_Deleter()`.
- [ ] Add for Batching.
- [ ] Add `AllowFiltering` to Deleter
- [ ] Modify `.Save(...)` and add `.SaveCompact(...)`
- [ ] Do final cleanups ( remove double cql whitespaces, some unused codes, ... )

### Related Projects
Similar projects for SQL databases.
- [db_walker](https://github.com/jozn/db_walker)
- [xo](https://github.com/xo/xo)
- [Android Orma](https://github.com/maskarade/Android-Orma)