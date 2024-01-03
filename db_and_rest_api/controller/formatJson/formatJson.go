package formatjson

import "log"

type formatResult struct {
	Title string      `json:"title"`
	Data  interface{} `json:"data"`
}

func DataSend(title string, data interface{}) formatResult {
	var dataSend = formatResult{
		Title: title, Data: data,
	}

	// Tambahkan logika untuk mengecek kesalahan (contoh: jika data tidak valid)
	if data == nil {
		log.Fatal("data does't valid")
	}

	return dataSend
}
