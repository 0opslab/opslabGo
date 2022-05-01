package main

import (
	"encoding/json"

	logger "github.com/sirupsen/logrus"
)

// func parseJson(b string, data interface{}) error {
// 	err := json.Unmarshal([]byte(b), &data)
// 	return err
// }

func toJosn(data interface{}) string {
	res, err := json.Marshal(&data)
	if err != nil {
		logger.Error("toJson Error:", err, data)
		return ""
	}
	return string(res)
}
