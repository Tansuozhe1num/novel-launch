package Biu

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data any) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}

func Failed(c *gin.Context, code int, msg string) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
	})
}

func JSON(ctx *gin.Context, status int, resp gin.H) {
	ctx.JSON(status, resp)
}

func ResponseMsg(ctx *gin.Context, status int, code int, msg string) {
	ctx.JSON(status, gin.H{
		"code": code,
		"msg":  msg,
	})
}
func ResponseData(ctx *gin.Context, status int, code, msg string, data interface{}) {
	ctx.JSON(status, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// ReturnGeneric 将任意 slice 转换成你需要的通用返回格式。
// list 必须是 slice/array；slice 每个元素会尝试转换为 map[string]interface{}。
func ReturnList(ctx *gin.Context, code int, msg string, list interface{}) {
	var data []map[string]interface{}

	if list == nil {
		data = []map[string]interface{}{}
	} else {
		rv := reflect.ValueOf(list)
		kind := rv.Kind()
		if kind != reflect.Slice && kind != reflect.Array {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": -1,
				"msg":  "list must be a slice or array",
				"data": []map[string]interface{}{},
			})
			return
		}

		for i := 0; i < rv.Len(); i++ {
			elem := rv.Index(i).Interface()

			// 直接是 map[string]interface{} 的情况
			if m, ok := elem.(map[string]interface{}); ok {
				data = append(data, m)
				continue
			}

			// 元素为 nil 的情况 -> 空对象
			if elem == nil {
				data = append(data, map[string]interface{}{})
				continue
			}

			// 其它类型：尝试 marshal -> unmarshal 为 map[string]interface{}
			b, err := json.Marshal(elem)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"code": -1,
					"msg":  "marshal element failed: " + err.Error(),
					"data": nil,
				})
				return
			}

			var mm map[string]interface{}
			if err := json.Unmarshal(b, &mm); err != nil {
				// 如果不是 JSON object（比如是数组或基本类型），将其包一层，保证 data 的每项都是 object
				data = append(data, map[string]interface{}{"value": elem})
			} else {
				data = append(data, mm)
			}
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
