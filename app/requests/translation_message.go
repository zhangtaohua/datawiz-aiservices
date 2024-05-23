// Package requests 处理请求数据和表单验证
package requests

import (
	"datawiz-aiservices/pkg/translator"
	"strings"
)

func RequiredMsg(key string) string {
	transMsg := []string{
		"required:",
		translator.TransHandler.T(key),
		translator.TransHandler.T("v.required"),
	}

	return strings.Join(transMsg, "")
}

func MinMsg(key string, number string) string {
	transMsg := []string{
		"min:",
		translator.TransHandler.T(key),
		translator.TransHandler.T("v.atleast"),
		number,
		translator.TransHandler.T("v.szieCh"),
	}

	return strings.Join(transMsg, "")
}

func MinCnMsg(key string, number string) string {
	transMsg := []string{
		"min_cn:",
		translator.TransHandler.T(key),
		translator.TransHandler.T("v.atleast"),
		number,
		translator.TransHandler.T("v.szieCh"),
	}

	return strings.Join(transMsg, "")
}

func MaxMsg(key string, number string) string {
	transMsg := []string{
		"max:",
		translator.TransHandler.T(key),
		translator.TransHandler.T("v.nomore"),
		number,
		translator.TransHandler.T("v.szieCh"),
	}

	return strings.Join(transMsg, "")
}

func MaxCnMsg(key string, number string) string {
	transMsg := []string{
		"max_cn:",
		translator.TransHandler.T(key),
		translator.TransHandler.T("v.nomore"),
		number,
		translator.TransHandler.T("v.szieCh"),
	}

	return strings.Join(transMsg, "")
}

func InMsg(key string, ins []string) string {
	transMsg := []string{
		"in:",
		translator.TransHandler.T(key),
		translator.TransHandler.T("v.in"),
		" [",
		strings.Join(ins, ", "),
		"]",
	}

	return strings.Join(transMsg, "")
}

func NotExistMsg(key string) string {
	transMsg := []string{
		"not_exists:",
		translator.TransHandler.T(key),
		translator.TransHandler.T("v.exists"),
	}

	return strings.Join(transMsg, "")
}

func NotExistUnionMsg(keys []string) string {
	transMsg := []string{}
	for _, key := range keys {
		transMsg = append(transMsg, translator.TransHandler.T(key))
		transMsg = append(transMsg, " ")
	}
	transMsg = append(transMsg, translator.TransHandler.T("v.union"))
	transMsg = append(transMsg, translator.TransHandler.T("v.exists"))

	return strings.Join(transMsg, "")
}
