package mysql_heap

const (
	TypeUnspecified byte = 0
	TypeTiny        byte = 1
	TypeShort       byte = 2
	TypeLong        byte = 3
	TypeFloat       byte = 4
	TypeDouble      byte = 5
	TypeNull        byte = 6
	TypeTimestamp   byte = 7
	TypeLonglong    byte = 8
	TypeInt24       byte = 9
	TypeDate        byte = 10
	/* TypeDuration original name was TypeTime, renamed to TypeDuration to resolve the conflict with Go type Time.*/
	TypeDuration byte = 11
	TypeDatetime byte = 12
	TypeYear     byte = 13
	TypeNewDate  byte = 14
	TypeVarchar  byte = 15
	TypeBit      byte = 16

	TypeJSON       byte = 0xf5
	TypeNewDecimal byte = 0xf6
	TypeEnum       byte = 0xf7
	TypeSet        byte = 0xf8
	TypeTinyBlob   byte = 0xf9
	TypeMediumBlob byte = 0xfa
	TypeLongBlob   byte = 0xfb
	TypeBlob       byte = 0xfc
	TypeVarString  byte = 0xfd
	TypeString     byte = 0xfe
	TypeGeometry   byte = 0xff
)

func NeedQuotes(tp byte) bool {
	return !(IsNumberType(tp) || IsFloatType(tp))
}

// IsNumberType returns true if tp is number type
func IsNumberType(tp byte) bool {
	switch tp {
	case TypeTiny, TypeShort, TypeLong, TypeLonglong, TypeInt24, TypeYear:
		return true
	}

	return false
}

// IsFloatType returns true if tp is float type
func IsFloatType(tp byte) bool {
	switch tp {
	case TypeFloat, TypeDouble, TypeNewDecimal:
		return true
	}

	return false
}
