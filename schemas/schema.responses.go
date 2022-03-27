// tidak akan muncul di database, karena hanya untuk pesan respons saja

package schemas

type Response struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
	Message string `json:"message"`
}