/*
@Time : 2022/3/25
@Author : jzd
@Project: msgsvc
*/
package manager

import (
	"io"
	"net/http"
)

func callback(url string, body io.Reader) {
	http.Post(url, "application/json", body)
}
