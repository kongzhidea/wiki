package mnumber

import "strconv"

func ParseInt(data string) int64 {
	result, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return 0
	}
	return result
}

func FormatInt(data int64) string {
	result := strconv.FormatInt(data, 10)
	return result
}

func JoinInt64(a []int64, sep string) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return FormatInt(a[0])
	case 2:
		// Special case for common small values.
		// Remove if golang.org/issue/6714 is fixed
		return FormatInt(a[0]) + sep + FormatInt(a[1])
	case 3:
		// Special case for common small values.
		// Remove if golang.org/issue/6714 is fixed
		return FormatInt(a[0]) + sep + FormatInt(a[1]) + sep + FormatInt(a[2])
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(FormatInt(a[i]))
	}

	b := make([]byte, n)
	bp := copy(b, FormatInt(a[0]))
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], FormatInt(s))
	}
	return string(b)
}

func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
