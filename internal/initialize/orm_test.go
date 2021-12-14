package initialize

import "testing"

func TestInitOrm(t *testing.T) {
	InitConfig()
	dsn := dsn()
	t.Log(dsn)
}
