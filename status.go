package medtronic

const (
	Status CommandCode = 0xCE
)

type StatusInfo struct {
	Normal    bool
	Bolusing  bool
	Suspended bool
}

func (pump *Pump) Status() (StatusInfo, error) {
	result, err := pump.Execute(Status, func(data []byte) interface{} {
		if len(data) < 4 || data[0] != 3 {
			return nil
		}
		// Observed values for data[1]:
		//   0: rewinding
		//   1: preparing to prime
		//   2: priming
		//   3: normal
		return StatusInfo{
			Normal:    data[1] == 0x03,
			Bolusing:  data[2] == 0x01,
			Suspended: data[3] == 0x01,
		}
	})
	if err != nil {
		return StatusInfo{}, err
	}
	return result.(StatusInfo), nil
}