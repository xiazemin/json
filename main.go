// go_json_test project main.go
package main

import (
"encoding/json"
"fmt"
)

func encode_from_map_test() {
	fmt.Printf("encode_from_map_test\n")

	m := map[string][]string{
		"level":   {"debug"},
		"message": {"File not found", "Stack overflow"},
	}

	if data, err := json.Marshal(m); err == nil {
		fmt.Printf("%s\n", data)
	}
}

type GameInfo struct {
	Game_id   int64
	Game_map  string
	Game_time int32
}

func encode_from_object_test() {

	fmt.Printf("encode_from_object_test\n")

	game_infos := []GameInfo{
		GameInfo{1, "map1", 20},
		GameInfo{2, "map2", 60},
	}

	if data, err := json.Marshal(game_infos); err == nil {
		fmt.Printf("%s\n", data)
	}
}

type GameInfo1 struct {
	Game_id   int64  `json:"game_id"`            // Game_id解析为game_id
	Game_map  string `json:"game_map,omitempty"` // GameMap解析为game_map, 忽略空置
	Game_time int32
	Game_name string `json:"-"` // 忽略game_name
}

func encode_from_object_tag_test() {

	fmt.Printf("encode_from_object_tag_test\n")

	game_infos := []GameInfo1{
		GameInfo1{1, "map1", 20, "name1"},
		GameInfo1{2, "map2", 60, "name2"},
		GameInfo1{3, "", 120, "name3"},
	}

	if data, err := json.Marshal(game_infos); err == nil {
		fmt.Printf("%s\n", data)
	}
}

type BaseObject struct {
	Field_a string
	Field_b string
}

type DeriveObject struct {
	BaseObject
	Field_c string
}

func encode_from_object_with_anonymous_field() {
	fmt.Printf("encode_from_object_with_anonymous_field\n")

	object := DeriveObject{BaseObject{"a", "b"}, "c"}

	if data, err := json.Marshal(object); err == nil {
		fmt.Printf("%s\n", data)
	}
}

// convert interface
// 在调用Marshal(v interface{})时，该函数会判断v是否满足json.Marshaler或者 encoding.Marshaler 接口，
// 如果满足，则会调用这两个接口来进行转换（如果两个都满足，优先调用json.Marshaler）
/*
// json.Marshaler
type Marshaler interface {
    MarshalJSON() ([]byte, error)
}

// encoding.TextMarshaler
type TextMarshaler interface {
    MarshalText() (text []byte, err error)
}
*/

type Point struct {
	X int
	Y int
}

func (pt Point) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"px" : %d, "py" : %d}`, pt.X, pt.Y)), nil
}

func encode_from_object_with_marshaler_interface() {

	fmt.Printf("encode_from_object_with_marshaler_interface\n")

	if data, err := json.Marshal(Point{50, 50}); err == nil {
		fmt.Printf("%s\n", data)
	} else {
		fmt.Printf("error %s\n", err.Error())
	}
}

type Point1 struct {
	X int
	Y int
}

func (pt Point1) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"px" : %d, "py" : %d}`, pt.X, pt.Y)), nil
}

func encode_from_object_with_text_marshaler_interface() {

	fmt.Printf("encode_from_object_with_text_marshaler_interface\n")

	if data, err := json.Marshal(Point1{50, 50}); err == nil {
		fmt.Printf("%s\n", data)
	} else {
		fmt.Printf("error %s\n", err.Error())
	}
}

// decode test
func decode_to_map_test() {

	fmt.Printf("decode_to_map_test\n")

	data := `[{"Level":"debug","Msg":"File: \"test.txt\" Not Found"},` +
		`{"Level":"","Msg":"Logic error"}]`

	var debug_infos []map[string]string
	json.Unmarshal([]byte(data), &debug_infos)

	fmt.Println(debug_infos)
}

type DebugInfo struct {
	Level  string
	Msg    string
	author string // 未导出字段不会被json解析
}

func (dbgInfo DebugInfo) String() string {
	return fmt.Sprintf("{Level: %s, Msg: %s}", dbgInfo.Level, dbgInfo.Msg)
}

func decode_to_object_test() {
	fmt.Printf("decode_to_object_test\n")

	data := `[{"level":"debug","msg":"File Not Found","author":"Cynhard"},` +
		`{"level":"","msg":"Logic error","author":"Gopher"}]`

	var dbgInfos []DebugInfo
	json.Unmarshal([]byte(data), &dbgInfos)

	fmt.Println(dbgInfos)
}

type DebugInfoTag struct {
	Level  string `json:"level"`   // level 解码为 Level
	Msg    string `json:"message"` // message 解码为 Msg
	Author string `json:"-"`       // 忽略Author
}

func (dbgInfo DebugInfoTag) String() string {
	return fmt.Sprintf("{Level: %s, Msg: %s}", dbgInfo.Level, dbgInfo.Msg)
}

func decode_to_object_tag_test() {
	fmt.Printf("decode_to_object_tag_test\n")

	data := `[{"level":"debug","message":"File Not Found","author":"Cynhard"},` +
		`{"level":"","message":"Logic error","author":"Gopher"}]`

	var dbgInfos []DebugInfoTag
	json.Unmarshal([]byte(data), &dbgInfos)

	fmt.Println(dbgInfos)

}

type Pointx struct{ X, Y int }

type Circle struct {
	Pointx
	Radius int
}

func (c Circle) String() string {
	return fmt.Sprintf("{Point{X: %d, Y :%d}, Radius: %d}",
		c.Pointx.X, c.Pointx.Y, c.Radius)
}

func decode_to_object_with_anonymous_field() {
	fmt.Printf("decode_to_object_with_anonymous_field\n")

	data := `{"X":80,"Y":80,"Radius":40}`

	var c Circle
	json.Unmarshal([]byte(data), &c)

	fmt.Println(c)
}

// decode convert interface
// 解码时根据参数是否满足json.Unmarshaler和encoding.TextUnmarshaler来调用相应函数（若两个函数都存在，则优先调用UnmarshalJSON）
/*
// json.Unmarshaler
type Unmarshaler interface {
    UnmarshalJSON([]byte) error
}

// encoding.TextUnmarshaler
type TextUnmarshaler interface {
    UnmarshalText(text []byte) error
}
*/

type Pointy struct{ X, Y int }

func (pt Pointy) MarshalJSON() ([]byte, error) {
	// no decode, just print
	return []byte(fmt.Sprintf(`{"X":%d,"Y":%d}`, pt.X, pt.Y)), nil
}

func decode_to_object_with_marshaler_interface() {
	fmt.Printf("decode_to_object_with_marshaler_interface\n")

	if data, err := json.Marshal(Pointy{50, 50}); err == nil {
		fmt.Printf("%s\n", data)
	}
}

func main() {
	fmt.Println("json test!")

	fmt.Printf("ecode test\n")

	encode_from_map_test()

	encode_from_object_test()

	encode_from_object_tag_test()

	encode_from_object_with_anonymous_field()

	encode_from_object_with_marshaler_interface()

	encode_from_object_with_text_marshaler_interface()

	fmt.Printf("decode test\n")

	decode_to_map_test()

	decode_to_object_test()

	decode_to_object_tag_test()

	decode_to_object_with_anonymous_field()

	decode_to_object_with_marshaler_interface()

}