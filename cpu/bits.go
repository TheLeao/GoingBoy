package cpu

func GetMsb(word int) int {
	return word >> 8
}

func GetLsb(word int) int {
	return word & 0xff
}

func ToWord(msb int, lsb int) int {
	return msb << 8 | lsb
}

func ToWord(bytes []int) int {
	return ToWord(bytes[1], bytes[0])
}

func GetBit(byteValue int, position int) bool {
	return (byteValue & (1 << position)) != 0
}

func SetBit(byteValue int, position int) {
	return (byteValue | (1 << position)) & 0xff;
}

func SetBit(byteValue int, position int, value bool) {
	if (value) {
		SetBit(byteValue, position)
	} else {
		ClearBit(byteValue, position)
	}
}

func ClearBit(byteValue int, position int) int {
	return ~(1 << position) & byteValue & 0xff
}

func ToSigned(byteValue) int {
	if byteValue & (1 << 7)) == 0 {
		return byteValue
	} else {
		return byteValue - 0x100
	}
}