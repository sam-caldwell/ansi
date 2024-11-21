package ansi

// Error - print error message
func Error(msg ...interface{}) *Color {
	return (&Color{}).Error(msg...)
}

// Errorf - print error message
func Errorf(format string, msg ...interface{}) *Color {
	return (&Color{}).Errorf(format, msg...)
}

// Errorln - print error message
func Errorln(msg ...interface{}) *Color {
	return (&Color{}).Errorln(msg...)
}

// Error - print error message
func (c *Color) Error(msg ...interface{}) *Color {
	c.Red().Print("[ERROR]: ")
	for _, m := range msg {
		c.Print(m.(string))
	}
	return c
}

// Errorf - print error message
func (c *Color) Errorf(format string, msg ...interface{}) *Color {
	return Red().Printf("[ERROR]: "+format, msg...)
}

// Errorln - print error message
func (c *Color) Errorln(msg ...interface{}) *Color {
	return c.Errorln(msg...).LF()
}
