package sonyflake

import "github.com/sony/sonyflake"

var sf = sonyflake.NewSonyflake(sonyflake.Settings{})

func GenID() (uint64, error) {
	return sf.NextID()
}
