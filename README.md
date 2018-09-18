# Cassandra walker
Cassandra walker is and ORM (Object Relation Mapper) like, for cassandra.

`cassandra_walker` is command-line tool to generate Golang code based for cassandra keyspaces (databases).

with pointing `cassandra_walker` to cassandra cluster, `cassandra_walker` find tables in each keyspaces and for each tables create golang types, and idiom go codes.

## Quickstart

Install `cassandra_walker` with:
```
go get -u github.com/jozn/cassandra_walker
```

Then point to cassandra node to genrate code for a keyspace (ex `twitter`):

```
cassandra_walker twitter
```

This will produce `xc` folder in current directory, and puts generated golang codes in this folder.

## Command Line Parameters

Use ` cassandra_walker -h` to see parameter options.

```
Usage: cassandra_walker [--host HOST] [--port PORT] [--verbose] [--dir DIR] [--package PACKAGE] [--minimize] [KEYSPACES [KEYSPACES ...]]

Positional arguments:
  KEYSPACES              cassandra keyspaces to build

Options:
  --host HOST, -c HOST   cassandra cluster address (default 127.0.0.1)
  --port PORT, -p PORT   cassandra port (default 9042)
  --verbose, -v          verbosity Log
  --dir DIR, -d DIR      output of generated codes (default './')
  --package PACKAGE      package of go
  --minimize, -m          minimize docs
  --help, -h             display this help and exit
```

## Guides
Lets see how to use this tool.
We will follow twitter sample in [sample directory](https://github.com/jozn/cassandra_walker/tree/master/samples/twitter)

Assume you have this cassandra keyspace:
```cql
CREATE KEYSPACE twitter
  WITH REPLICATION = {
   'class' : 'SimpleStrategy',
   'replication_factor' : 1
 };

CREATE TABLE twitter.twitt (
	user_id bigint,
	twiit_id varchar,
	body varchar,
	create_time int,
	PRIMARY KEY (user_id,twiit_id)
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
cassandra_walker twitter --host 127.0.01 --port 9042
```

This will generates codes in `xc` directory.

```
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
```
From now we just work with `Twiit` type. But
For each table, we now have

```
// create


```
