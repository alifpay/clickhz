package storages

//ClickDB - clickhouse
type ClickDB interface {
	AppLog(data [][]interface{})
	Close()
}
