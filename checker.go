package imparse

type Checker interface {
	CheckFrame(frame Frame) error
}
