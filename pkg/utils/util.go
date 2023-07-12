package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	// 读取请求的主体内容
	if body, err := io.ReadAll(r.Body); err == nil {
		// 将JSON格式的数据解析并映射到指定的目标对象x上
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
