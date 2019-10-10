package utils

// Python-compatible mod
func Mod(a int, b int) int {
	var mod int = a % b
	if (mod < 0 && b > 0) || (mod > 0 && b < 0) {
		return mod + b
	}
	return mod
}


// Python-compatible divmod
func Divmod(a int, b int) (int, int) {
	var div int = a / b
	var mod int = a % b
	if (mod < 0 && b > 0) || (mod > 0 && b < 0) {
		return div - 1, mod + b
	}
	return div, mod
}

