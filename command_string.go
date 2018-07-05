// Code generated by "stringer -type Command"; DO NOT EDIT.

package medtronic

import "strconv"

const _Command_name = "acknakcgmWriteTimestampsetBasalPatternAsetBasalPatternBsetClocksetMaxBolusbolusselectBasalPatternsetAbsoluteTempBasalsuspendbuttonwakeupsetPercentTempBasalsetMaxBasalsetBasalRatesclockpumpIDbatteryreservoirfirmwareVersionerrorStatushistoryPagecarbUnitsglucoseUnitscarbRatiosinsulinSensitivitiesglucoseTargets512modelsettings512basalRatesbasalPatternAbasalPatternBtempBasalglucosePageisigPagecalibrationFactorhistoryPageCountglucoseTargetssettingscgmPageCountstatusvcntrPage"

var _Command_map = map[Command]string{
	6:   _Command_name[0:3],
	21:  _Command_name[3:6],
	40:  _Command_name[6:23],
	48:  _Command_name[23:39],
	49:  _Command_name[39:55],
	64:  _Command_name[55:63],
	65:  _Command_name[63:74],
	66:  _Command_name[74:79],
	74:  _Command_name[79:97],
	76:  _Command_name[97:117],
	77:  _Command_name[117:124],
	91:  _Command_name[124:130],
	93:  _Command_name[130:136],
	105: _Command_name[136:155],
	110: _Command_name[155:166],
	111: _Command_name[166:179],
	112: _Command_name[179:184],
	113: _Command_name[184:190],
	114: _Command_name[190:197],
	115: _Command_name[197:206],
	116: _Command_name[206:221],
	117: _Command_name[221:232],
	128: _Command_name[232:243],
	136: _Command_name[243:252],
	137: _Command_name[252:264],
	138: _Command_name[264:274],
	139: _Command_name[274:294],
	140: _Command_name[294:311],
	141: _Command_name[311:316],
	145: _Command_name[316:327],
	146: _Command_name[327:337],
	147: _Command_name[337:350],
	148: _Command_name[350:363],
	152: _Command_name[363:372],
	154: _Command_name[372:383],
	155: _Command_name[383:391],
	156: _Command_name[391:408],
	157: _Command_name[408:424],
	159: _Command_name[424:438],
	192: _Command_name[438:446],
	205: _Command_name[446:458],
	206: _Command_name[458:464],
	213: _Command_name[464:473],
}

func (i Command) String() string {
	if str, ok := _Command_map[i]; ok {
		return str
	}
	return "Command(" + strconv.FormatInt(int64(i), 10) + ")"
}
