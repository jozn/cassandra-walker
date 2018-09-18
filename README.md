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

