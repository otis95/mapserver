package httpserver

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/macaron.v1"
)

func StartHttpServer(addr string) {

	macaron.Env = macaron.PROD
	m := macaron.Classic()
	m.Use(macaron.Renderer())

	m.Get("/", index)
	m.Get("/788865972/:z/:x/:y", gaode) //静态地图请求接口
	m.Get("/favicon.ico", favicon)
	m.RunAddr(addr)
}

func index(ctx *macaron.Context) {
	ctx.HTML(200, "map_hz")
}

//解析前端请求参数，然后解析返回所需要的瓦片
func gaode(ctx *macaron.Context) {
	zoom := "L" + ctx.Params(":z")
	x, _ := strconv.Atoi(ctx.Params(":x"))
	y, _ := strconv.Atoi(ctx.Params(":y"))
	img_name := "C" + Tool_DecimalByteSlice2HexString(IntToBytes(x)) + ".png"
	path := "R" + Tool_DecimalByteSlice2HexString(IntToBytes(y))
	fmt.Println("public/" + zoom + "/" + path + "/" + img_name)
	ctx.ServeFileContent("public/" + zoom + "/" + path + "/" + img_name)
}

func Tool_DecimalByteSlice2HexString(DecimalSlice []byte) string {
	var sa = make([]string, 0)
	for _, v := range DecimalSlice {
		parse_str := fmt.Sprintf("%02x", v)
		sa = append(sa, parse_str)
	}
	ss := strings.Join(sa, "")
	return ss
}

func IntToBytes(n int) []byte {
	tmp := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

func favicon(ctx *macaron.Context) {
	ctx.ServeFileContent("public/pig.png")
}
