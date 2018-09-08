package cwalker

const (
	TEMPLATES_DIR = `C:\Go\_gopath\src\ms\cassandra_walker\templates\`
	//OUTPUT_DIR_GO_X       = `C:\Go\_gopath\src\ms\cassandra_walker\play\out\`
	OUTPUT_DIR_GO_X = `C:\Go\_gopath\src\ms\sun\shared\xc\`
)

var DATABASES = []string{"sunc_file", "sunc_msg"}

type ConfigArgs struct {
	keyspaces []string `arg:"-k,help:cassandra keyspaces to build "`
	Host      string   `arg:"-c,help:cassandra cluster address (default 127.0.0.1)"`
	Port      int      `arg:"-p,help:cassandra port (default 9042)"`
	Verbose   bool     `arg:"-v,help:verbosity Log"`
	OutputDir string   `arg:"help:output of generated codes"`
}
