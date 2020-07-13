package bit

func IsSet(bitmap uint8, flag uint8) bool {
	return (bitmap & flag) == flag
}
