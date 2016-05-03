package window

type Event interface {
	eventInterface()
}

type EventClose struct{}

func (_ EventClose) eventInterface() {}

type EventChar struct {
	Char rune
}

func (_ EventChar) eventInterface() {}

type EventCursorEnter struct{}

func (_ EventCursorEnter) eventInterface() {}

type EventCursorExit struct{}

func (_ EventCursorExit) eventInterface() {}

type EventCursorPos struct {
	X float64
	Y float64
}

func (_ EventCursorPos) eventInterface() {}

type EventDrop struct {
	Names []string
}

func (_ EventDrop) eventInterface() {}

type EventFocusGained struct{}

func (_ EventFocusGained) eventInterface() {}

type EventFocusLost struct{}

func (_ EventFocusLost) eventInterface() {}

type EventIconified struct{}

func (_ EventIconified) eventInterface() {}

type EventDeIconified struct{}

func (_ EventDeIconified) eventInterface() {}

type EventKeyPress struct {
	Key      Key
	ScanCode int
	Mod      ModifierKey
}

func (_ EventKeyPress) eventInterface() {}

type EventKeyRelease struct {
	Key      Key
	ScanCode int
	Mod      ModifierKey
}

func (_ EventKeyRelease) eventInterface() {}

type EventKeyRepeat struct {
	Key      Key
	ScanCode int
	Mod      ModifierKey
}

func (_ EventKeyRepeat) eventInterface() {}

type EventMouseButtonPress struct {
	Mod    ModifierKey
	Button MouseButton
}

func (_ EventMouseButtonPress) eventInterface() {}

type EventMouseButtonRelease struct {
	Mod    ModifierKey
	Button MouseButton
}

func (_ EventMouseButtonRelease) eventInterface() {}

type EventRefresh struct{}

func (_ EventRefresh) eventInterface() {}

type EventScroll struct {
	X float64
	Y float64
}

func (_ EventScroll) eventInterface() {}

type EventSize struct {
	Width  int
	Height int
}

func (_ EventSize) eventInterface() {}
