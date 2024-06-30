package ansi

var debugMode bool

func init() {
	debugMode = false // Use color by default
}

// DisableDebug - Turn on ANSI Debug print functionality. (This is a global setting)
func DisableDebug() *Color {
	return (&Color{}).DisableDebug()
}

// DisableDebug - Turn off ANSI Debug print functionality. (This is a global setting)
func (c *Color) DisableDebug() *Color {
	debugMode = false
	return c
}

// EnableDebug - Turn on ANSI Debug print functionality. (This is a global setting)
func (c *Color) EnableDebug() *Color {
	debugMode = true
	return c
}

// EnableDebug - Turn on ANSI Debug print functionality. (This is a global setting)
func EnableDebug() *Color {
	return (&Color{}).EnableDebug()
}

// Debug - print debug message (if debugging is enabled)
func Debug(msg ...interface{}) *Color {
	return (&Color{}).Debug(msg...)
}

// Debugf - print debug message (if debugging is enabled)
func Debugf(format string, msg ...interface{}) *Color {
	return (&Color{}).Debugf(format, msg...)
}

// Debugln - print debug message (if debugging is enabled)
func Debugln(msg ...interface{}) *Color {
	return (&Color{}).Debugln(msg...)
}

// Debug - print debug message (if debugging is enabled)
func (c *Color) Debug(msg ...interface{}) *Color {
	return c.Debugf("%v ", msg)
}

// Debugf - print debug message (if debugging is enabled)
func (c *Color) Debugf(format string, msg ...interface{}) *Color {
	if debugMode {
		return Yellow().Printf("[DEBUG]: "+format, msg...)
	}
	return c
}

// Debugln - print debug message (if debugging is enabled)
func (c *Color) Debugln(msg ...interface{}) *Color {
	return c.Debug(msg...).LF()
}
