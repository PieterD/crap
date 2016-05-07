package win

type DefaultEventHandler struct{}

func (_ DefaultEventHandler) FrameTick() error                                            { return nil }
func (_ DefaultEventHandler) FrameAnim(ticks int) error                                   { return nil }
func (_ DefaultEventHandler) FrameDrop(ticks int) error                                   { return nil }
func (_ DefaultEventHandler) EventChar(char rune)                                         {}
func (_ DefaultEventHandler) EventClose()                                                 {}
func (_ DefaultEventHandler) EventRefresh()                                               {}
func (_ DefaultEventHandler) EventDragDrop(names []string)                                {}
func (_ DefaultEventHandler) EventCursorEnter()                                           {}
func (_ DefaultEventHandler) EventCursorExit()                                            {}
func (_ DefaultEventHandler) EventCursorPos(x, y float64)                                 {}
func (_ DefaultEventHandler) EventFocusGained()                                           {}
func (_ DefaultEventHandler) EventFocusLost()                                             {}
func (_ DefaultEventHandler) EventIconified()                                             {}
func (_ DefaultEventHandler) EventDeiconified()                                           {}
func (_ DefaultEventHandler) EventKeyPress(key Key, scancode int, mod ModifierKey)        {}
func (_ DefaultEventHandler) EventKeyRepeat(key Key, scancode int, mod ModifierKey)       {}
func (_ DefaultEventHandler) EventKeyRelease(key Key, scancode int, mod ModifierKey)      {}
func (_ DefaultEventHandler) EventMouseButtonPress(button MouseButton, mod ModifierKey)   {}
func (_ DefaultEventHandler) EventMouseButtonRelease(button MouseButton, mod ModifierKey) {}
func (_ DefaultEventHandler) EventScroll(x, y float64)                                    {}
func (_ DefaultEventHandler) EventSize(width, height int)                                 {}
