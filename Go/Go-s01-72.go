package net

type Error interface{
	error
	Timeout() bool   // Is the error a timeout?
	Temporary() bool // Is the error temporary?
}
