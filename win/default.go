package win

type DefaultHandler struct{}

func (_ DefaultHandler) Tick() error                                                 { return nil }
func (_ DefaultHandler) Anim(ticks int) error                                        { return nil }
func (_ DefaultHandler) Drop(ticks int) error                                        { return nil }
func (_ DefaultHandler) EventChar(char rune)                                         {}
func (_ DefaultHandler) EventClose()                                                 {}
func (_ DefaultHandler) EventRefresh()                                               {}
func (_ DefaultHandler) EventDragDrop(names []string)                                {}
func (_ DefaultHandler) EventCursorEnter()                                           {}
func (_ DefaultHandler) EventCursorExit()                                            {}
func (_ DefaultHandler) EventCursorPos(x, y float64)                                 {}
func (_ DefaultHandler) EventFocusGained()                                           {}
func (_ DefaultHandler) EventFocusLost()                                             {}
func (_ DefaultHandler) EventIconified()                                             {}
func (_ DefaultHandler) EventDeiconified()                                           {}
func (_ DefaultHandler) EventKeyPress(key Key, scancode int, mod ModifierKey)        {}
func (_ DefaultHandler) EventKeyRepeat(key Key, scancode int, mod ModifierKey)       {}
func (_ DefaultHandler) EventKeyRelease(key Key, scancode int, mod ModifierKey)      {}
func (_ DefaultHandler) EventMouseButtonPress(button MouseButton, mod ModifierKey)   {}
func (_ DefaultHandler) EventMouseButtonRelease(button MouseButton, mod ModifierKey) {}
func (_ DefaultHandler) EventScroll(x, y float64)                                    {}
func (_ DefaultHandler) EventSize(width, height int)                                 {}
