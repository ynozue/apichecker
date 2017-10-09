package main

import (
	"testing"
	"regexp"
)

func TestGetAPI_パラメーターエラー(t *testing.T) {
	if getAPI("") != "not endpoint" {
		t.Error("failed validation check")
	}
}

func TestGetAPI_正常(t *testing.T) {
	r := regexp.MustCompile("OK \\(expire=[0-9]{2}/[0-9]{2}/[0-9]{2} [0-9]{2}:[0-9]{2}\\)\nhttps://www.yahoo.co.jp")
	if !r.MatchString(getAPI("https://www.yahoo.co.jp")) {
		t.Error("function format error")
	}
}

func TestGetAPI_証明書エラー(t *testing.T) {
	r := regexp.MustCompile("NG\n.+")
	if !r.MatchString(getAPI("https://www.yahoo.jp")) {
		t.Error("function format error")
	}
}
