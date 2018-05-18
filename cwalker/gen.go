package cwalker

import (
	"bytes"
	"io/ioutil"
	"ms/sun/shared/helper"
	"os"
	"os/exec"
	"text/template"
    "fmt"
)

func build(gen *GenOut) {

	//writeOutput("z_models.go", buildFromTemplate("models.go.tpl", gen))
	writeOutput("zc_models.go", buildFromTemplate("models_types.go.tpl", gen))

    for _,t := range gen.Tables{
        fileName := fmt.Sprintf("zc_%s.go",t.TableName)
        writeOutput(fileName, buildFromTemplate("model.go.tpl", t))
    }

	if true {
        e1 := exec.Command("gofmt", "-w", OUTPUT_DIR_GO_X).Run()
        e2 := exec.Command("goimports", "-w", OUTPUT_DIR_GO_X).Run()
        helper.NoErr(e1)
        helper.NoErr(e2)
    }
}

func genTablesOrma(tplName string, gen *GenOut) {
	tpl := _getTemplate(tplName)

	for _, table := range gen.TablesExtracted {
		buffer := bytes.NewBufferString("")
		err := tpl.Execute(buffer, table)
		helper.NoErr(err)
		writeOutput("zz_"+table.TableName+".go", buffer.String())
	}

}

func writeOutput(fileName, output string) {
	//println(output)
	ioutil.WriteFile(OUTPUT_DIR_GO_X+fileName, []byte(output), os.ModeType)

}

//func buildFromTemplate(tplName string, gen *GenOut) string {
func buildFromTemplate(tplName string, gen interface{}) string {
	tpl := template.New("" + tplName)
	tpl.Funcs(NewTemplateFuncs())
	tplGoInterface, err := ioutil.ReadFile(TEMPLATES_DIR + tplName)
	helper.NoErr(err)
	tpl, err = tpl.Parse(string(tplGoInterface))
	helper.NoErr(err)

	buffer := bytes.NewBufferString("")
	err = tpl.Execute(buffer, gen)
	helper.NoErr(err)

	return buffer.String()
}


func _getTemplate(tplName string) *template.Template {
	tpl := template.New("" + tplName)
	tpl.Funcs(NewTemplateFuncs())
	tplGoInterface, err := ioutil.ReadFile(TEMPLATES_DIR + tplName)
	helper.NoErr(err)
	tpl, err = tpl.Parse(string(tplGoInterface))
	helper.NoErr(err)
	return tpl
}
