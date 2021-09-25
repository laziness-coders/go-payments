package main

type Payment interface {
	Authorize()
	Capture()
	Void()
	Cancel()
	Refund()
	GenerateToken()
}
