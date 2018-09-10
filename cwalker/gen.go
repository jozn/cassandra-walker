package cwalker

import (
	"bytes"
	"fmt"
	"github.com/jozn/cassandra_walker/bind"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"text/template"
)

func build(gen *GenOut) {

	//writeOutput("z_models.go", buildFromTemplate("models.go.tpl", gen))
	writeOutput("zc_models.go", buildFromTemplate("models_types.go.tpl", gen))

	for _, t := range gen.Tables {
		fileName := fmt.Sprintf("zc_%s.go", t.TableName)
		writeOutput(fileName, buildFromTemplate("model.go.tpl", t))
	}

	if true {
		/*e1 := exec.Command("gofmt", "-w", OUTPUT_DIR_GO_X).Run()
		  e2 := exec.Command("goimports", "-w", OUTPUT_DIR_GO_X).Run()*/
		e1 := exec.Command("gofmt", "-w", args.OutputDir).Run()
		e2 := exec.Command("goimports", "-w", args.OutputDir).Run()
		NoErr(e1)
		NoErr(e2)
	}
}

func genTablesOrma(tplName string, gen *GenOut) {
	tpl := _getTemplate(tplName)

	for _, table := range gen.TablesExtracted {
		buffer := bytes.NewBufferString("")
		err := tpl.Execute(buffer, table)
		NoErr(err)
		writeOutput("zz_"+table.TableName+".go", buffer.String())
	}

}

func writeOutput(fileName, output string) {
	//println(output)
	//ioutil.WriteFile(OUTPUT_DIR_GO_X+fileName, []byte(output), os.ModeType)
	os.MkdirAll(args.OutputDir, os.ModeDir)
	file := path.Join(args.OutputDir, fileName)

	ioutil.WriteFile(file, []byte(output), os.ModeType)

}

//func buildFromTemplate(tplName string, gen *GenOut) string {
func buildFromTemplate(tplName string, gen interface{}) string {
	tpl := template.New("" + tplName)
	tpl.Funcs(NewTemplateFuncs())

	//tplGoInterface, err := ioutil.ReadFile(TEMPLATES_DIR + tplName)
	tplGoInterface, err := bind.Asset(tplName)
	NoErr(err)
	tpl, err = tpl.Parse(string(tplGoInterface))
	NoErr(err)

	buffer := bytes.NewBufferString("")
	err = tpl.Execute(buffer, gen)
	NoErr(err)

	return buffer.String()
}

func _getTemplate(tplName string) *template.Template {
	tpl := template.New("" + tplName)
	tpl.Funcs(NewTemplateFuncs())
	//tplGoInterface, err := ioutil.ReadFile(TEMPLATES_DIR + tplName)
	tplGoInterface, err := bind.Asset(tplName)
	NoErr(err)
	tpl, err = tpl.Parse(string(tplGoInterface))
	NoErr(err)
	return tpl
}
