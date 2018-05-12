package src

import (
	"bytes"
	"io/ioutil"
	"ms/sun/shared/helper"
	"os"
	"os/exec"
	"text/template"
)

func build(gen *GenOut) {

	genModels(gen)
	writeOutput("z_xo.go", buildFromTemplate("xo.go.tpl", gen))
	//writeOutput("z_models.go", buildFromTemplate("models.go.tpl", gen))
	writeOutput("z_cache.go", buildFromTemplate("cache.go.tpl", gen))
	writeOutput("z_event.go", buildFromTemplate("event.go.tpl", gen))
	writeOutput("z_manual.go", buildFromTemplate("manual.go", gen))
	writeOutput("z_index.go", buildFromTemplate("index.go.tpl", gen))
	writeOutput("z_cache_secondary_index.go", buildFromTemplate("cache_secondary_index.go.tpl", gen))
	writeOutput("J.java", buildFromTemplate("J.java", gen))
	writeOutput("triggers.sql", buildFromTemplate("triggers.sql", gen))
	writeOutput("trigger.go", buildFromTemplate("trigger.go.tpl", gen))

	writeOutputConst("tables.go", buildFromTemplate("const.go.tpl", gen))

	genTablesOrma("orm.go.tpl", gen)

	PtMsgdef, converter := Gen_ProtosForTables(gen.Tables)
	writeOutput("TablePBCon.go", converter)
	ioutil.WriteFile(OUTPUT_PROTO_DIR+"pb_tables.proto", []byte(PtMsgdef), os.ModeType)

	if FORMAT {
        e1 := exec.Command("gofmt", "-w", OUTPUT_DIR_GO_X).Run()
        e2 := exec.Command("goimports", "-w", OUTPUT_DIR_GO_X).Run()
        helper.NoErr(e1)
        helper.NoErr(e2)
    }
}

func genTablesOrma(tplName string, gen *GenOut) {
	tpl := _getTemplate(tplName)

	for _, table := range gen.Tables {
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

func writeOutputConst(fileName, output string) {
    //println(output)
    ioutil.WriteFile(OUTPUT_DIR_GO_X_CONST+fileName, []byte(output), os.ModeType)

}

func buildFromTemplate(tplName string, gen *GenOut) string {
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

func genModels(gen *GenOut) {
	tpl := _getTemplate("models.go.tpl")
	tables := []*Table{}
	for _, t := range gen.Tables {
		if !skipTableModel(t.TableNameSql) {
			tables = append(tables, t)
		}
	}

	buffer := bytes.NewBufferString("")
	err := tpl.Execute(buffer, tables)
	helper.NoErr(err)
	writeOutput("z_models.go", buffer.String())
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
