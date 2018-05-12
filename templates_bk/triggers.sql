{{range .Tables }}
{{- if .PrimaryKey}}
{{- $TargetCol := ( ms_trigger_colmun .PrimaryKey.GoTypeOut ) }}
################################ {{.TableNameGo}} ######################################
{{ if .NeedTrigger}}
delimiter $$
{{- else}}
/* #### delimiter $$
{{- end}}
{{ $triggerName := (printf "%s%s" .TableName  "_OnCreateLogger") -}}
DROP TRIGGER IF EXISTS {{ $triggerName }} $$
CREATE TRIGGER {{ $triggerName }} AFTER INSERT ON {{.TableNameSql}}
  FOR EACH ROW
  BEGIN
    INSERT INTO trigger_log (ModelName,ChangeType,{{ $TargetCol }},CreatedSe) VALUES ("{{.TableNameGo}}","INSERT",NEW.{{.PrimaryKey.ColumnName}}, UNIX_TIMESTAMP(NOW()) );
  END;
$$

{{ $triggerName := (printf "%s%s" .TableName  "_OnUpdateLogger") -}}
DROP TRIGGER IF EXISTS {{ $triggerName }} $$
CREATE TRIGGER {{ $triggerName }} AFTER UPDATE ON {{.TableNameSql}}
  FOR EACH ROW
  BEGIN
  	INSERT INTO trigger_log (ModelName,ChangeType,{{ $TargetCol }},CreatedSe) VALUES ("{{.TableNameGo}}","UPDATE",NEW.{{.PrimaryKey.ColumnName}}, UNIX_TIMESTAMP(NOW()));
  END;
$$

{{ $triggerName := (printf "%s%s" .TableName  "_OnDeleteLogger") -}}
DROP TRIGGER IF EXISTS {{ $triggerName }} $$
CREATE TRIGGER {{ $triggerName }} AFTER DELETE ON {{.TableNameSql}}
  FOR EACH ROW
  BEGIN
   	INSERT INTO trigger_log (ModelName,ChangeType,{{ $TargetCol }},CreatedSe) VALUES ("{{.TableNameGo}}","DELETE",OLD.{{.PrimaryKey.ColumnName}}, UNIX_TIMESTAMP(NOW()));
  END;
$$

{{ if .NeedTrigger}}
delimiter ;
{{- else}}
 #### delimiter ;*/
{{- end}}

{{- end }}
{{- end }}

###############################################################################################
################################## Delete of all triggers #####################################
/*
{{range .Tables }}
{{- if .PrimaryKey}}
### {{.TableNameGo}} ##
DROP TRIGGER IF EXISTS {{ .TableName }}_OnCreateLogger ;
DROP TRIGGER IF EXISTS {{ .TableName }}_OnUpdateLogger ;
DROP TRIGGER IF EXISTS {{ .TableName }}_OnDeleteLogger ;
{{- end }}
{{- end }}
*/