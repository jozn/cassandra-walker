package cwalker

type ConfigArgs struct {
	Keyspaces []string `arg:"positional,-k,help:cassandra keyspaces to build "`
	Host      string   `arg:"-c,help:cassandra cluster address (default 127.0.0.1)"`
	Port      int      `arg:"-p,help:cassandra port (default 9042)"`
	Verbose   bool     `arg:"-v,help:verbosity Log"`
	Dir       string   `arg:"-d,help:output of generated codes (default './')"`
	Package   string   `arg:"help:package of go"`
	Minimize  bool     `arg:"-m,help:minimize docs"`
}
