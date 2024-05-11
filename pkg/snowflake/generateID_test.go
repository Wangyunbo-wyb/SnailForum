package snowflake

import "testing"

func TestSnowflakeSeqGenerator_GenerateId1(t *testing.T) {
	var dataCenterId, workId int64 = 1, 1
	generator, err := NewSnowflakeSeqGenerator(dataCenterId, workId)
	if err != nil {
		t.Error(err)
		return
	}
	var x, y string
	for i := 0; i < 100; i++ {
		y = generator.GenerateId1("", "")
		if x == y {
			t.Errorf("x(%s) & y(%s) are the same", x, y)
		}
		x = y
	}
}
