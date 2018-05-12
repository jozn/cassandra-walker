package x

import (
	"ms/sun/shared/base"
	"strings"
	"time"
)

func init() {
	go triggerLoader()
}

type TriggerStringModel interface {
	OnInsert(ins []string)
	OnUpdate(ins []string)
	OnDelete(ins []string)
}

type TriggerIntModel interface {
	OnInsert(ins []int)
	OnUpdate(ins []int)
	OnDelete(ins []int)
}

type TriggerModelListener struct {
	{{- range .TablesTriggers }}
		{{- if eq .XPrimaryKeyGoType "int"}}
			{{.TableNameGo}} TriggerIntModel
		{{- else}}
			{{.TableNameGo}} TriggerStringModel
		{{- end}}
	{{- end}}
}

var lastLoaded int
var ArrTriggerListeners = make([]TriggerModelListener, 10)
var ActivateTrigger = false

func triggerLoader() {
	time.Sleep(time.Second * 10)
	for {
		time.Sleep(time.Second)
		if !ActivateTrigger {
			continue
		}

		selector := NewTriggerLog_Selector()
		if lastLoaded > 0 {
			selector.Id_GT(lastLoaded)
		}
		triggers, err := selector.OrderBy_Id_Asc().GetRows(base.DB)
		if err != nil || len(triggers) == 0 {
			continue
		}
		lastLoaded = triggers[len(triggers)-1].Id
		
		for _, listener := range ArrTriggerListeners {
			collect := triggerModelWalk{}
			for _, trig := range triggers {
		
				switch strings.ToUpper(trig.ChangeType) {
				case "INSERT":
					switch strings.ToUpper(trig.ModelName) {
					{{- range .TablesTriggers }}
					case "{{toUpper .TableNameGo}}": 
						if listener.{{.TableNameGo}} != nil {
							collect.{{.TableNameGo}}.OnInsert = append(collect.{{.TableNameGo}}.OnInsert, trig.{{ms_trigger_colmun .XPrimaryKeyGoType }} )
						}
					{{- end}}
					}
				case "UPDATE":
					switch strings.ToUpper(trig.ModelName) {
					{{- range .TablesTriggers }}
					case "{{toUpper .TableNameGo}}": 
						if listener.{{.TableNameGo}} != nil {
							collect.{{.TableNameGo}}.OnUpdate = append(collect.{{.TableNameGo}}.OnUpdate, trig.{{ms_trigger_colmun .XPrimaryKeyGoType }})
						}
					{{- end}}
					}
				case "DELETE":
					switch strings.ToUpper(trig.ModelName) {
					{{- range .TablesTriggers }}
					case "{{toUpper .TableNameGo}}": 
						if listener.{{.TableNameGo}} != nil {
							collect.{{.TableNameGo}}.OnDelete = append(collect.{{.TableNameGo}}.OnDelete, trig.{{ms_trigger_colmun .XPrimaryKeyGoType }})
						}
					{{- end}}
					}
				}
			}

			//each
			{{range .TablesTriggers }}
			if listener.{{.TableNameGo}} != nil {
				if len(collect.{{.TableNameGo}}.OnInsert) != 0 {
					listener.{{.TableNameGo}}.OnInsert(collect.{{.TableNameGo}}.OnInsert)
				}
				if len(collect.{{.TableNameGo}}.OnUpdate) != 0 {
					listener.{{.TableNameGo}}.OnUpdate(collect.{{.TableNameGo}}.OnUpdate)
				}
				if len(collect.{{.TableNameGo}}.OnDelete) != 0 {
					listener.{{.TableNameGo}}.OnDelete(collect.{{.TableNameGo}}.OnDelete)
				}
			}
			{{end}}
		}
	}
}

type triggerStringCollection struct {
	OnInsert []string
	OnUpdate []string
	OnDelete []string
}

type triggerIntCollection struct {
	OnInsert []int
	OnUpdate []int
	OnDelete []int
}

type triggerModelWalk struct {
	{{- range .TablesTriggers }}
		{{- if eq .XPrimaryKeyGoType "int"}}
			{{.TableNameGo}} triggerIntCollection
		{{- else}}
			{{.TableNameGo}} triggerStringCollection
		{{- end}}
	{{- end}}
}
