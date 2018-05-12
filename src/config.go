package src

const (
	TEMPLATES_DIR = `C:\Go\_gopath\src\ms\db_walker\templates\`
	//OUTPUT_DIR_GO_X       = `C:\Go\_gopath\src\ms\snake\play\`
	//OUTPUT_DIR_GO_X  = `C:\Go\_gopath\src\ms\sun\models\x\`
	OUTPUT_DIR_GO_X       = `C:\Go\_gopath\src\ms\sun\shared\x\`
	OUTPUT_DIR_GO_X_CONST = `C:\Go\_gopath\src\ms\sun\shared\x\xconst\`
	OUTPUT_PROTO_DIR      = `C:\Go\_gopath\src\ms\sun\shared\proto\`
	//DATABASE              = "sun"

	FORMAT = true
)

var DATABASES = []string{"sun", "sun_chat", "sun_file","sun_meta", "sun_push", "sun_log"}

var OutPutBuffer = &GenOut{
	PackageName: "x",
}
