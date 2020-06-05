package curl_test

import (
	"../curl"
	"testing"
)

func TestSetUrl(t *testing.T) {
	ch := curl.Init()
	t.Logf("%+v", ch)
	ch.SetUrl("http://baidu.com?q=xxx")
	t.Logf("%+v", ch)
}
