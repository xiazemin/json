package main

import (
	"encoding/json"
	"go/src/fmt"
)

type DebugInfo struct {
	Level  string
	Msg    string
	author string // 未导出字段不会被json解析
}
func main() {
	data := `[{"level":"debug","msg":"File Not Found","author":"Cynhard"},` +
		`{"level":"","msg":"Logic error","author":"Gopher"}]`

	var dbgInfos []DebugInfo
	e:=json.Unmarshal([]byte(data), &dbgInfos)
	fmt.Println(e,dbgInfos)
	b,e:=json.Marshal(dbgInfos)
	fmt.Println(string(b),e)
}
