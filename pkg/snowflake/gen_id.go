package snowflake

import (
	"fmt"
	"time"

	"go.uber.org/zap"

	sf "github.com/sony/sonyflake"
)

var (
	sonyFlake     *sf.Sonyflake
	sonyStartTime string
	sonyMachineID uint16
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

// Init 需传入当前的机器ID
func Init(startTime string, machineId uint16) (err error) {
	sonyStartTime = startTime
	sonyMachineID = machineId
	t, _ := time.Parse(sonyStartTime, "2023-01-01")
	settings := sf.Settings{
		StartTime: t,
		MachineID: getMachineID,
	}
	sonyFlake = sf.NewSonyflake(settings)
	zap.L().Info("[snowflake] init success")

	return
}

// GenID 返回生成的id值
func GenID() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("sonyflake not inited")
		return
	}

	id, err = sonyFlake.NextID()
	return
}
