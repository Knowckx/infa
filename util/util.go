package util

func ShortStr(in string) string {
	leng := 7 // length of short git commitID
	return ShortStrLen(in, leng)
}

func ShortStrLen(in string, leng int) string {
	out := in
	if len(out) > leng {
		out = out[0:leng]
	}
	return out
}
