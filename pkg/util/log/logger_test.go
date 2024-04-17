package log

import "testing"

func TestInitLog(t *testing.T) {
	InitLog()
	if LogrusObj == nil {
		t.Errorf("InitLog failed")
	}
	LogrusObj.WithField("test", "test").Info("test")
	LogrusObj.Info("test2")
}
