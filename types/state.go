package types

import s "blind/ship"

type State struct {
	WindowHeight int32
	WindowWidth  int32
	Ship         *s.Ship
}
