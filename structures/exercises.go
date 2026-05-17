package structures

type Weapon struct {
	On    bool
	Ammo  int
	Power int
}

func (w *Weapon) Shoot() bool {
	if w.Ammo <= 0 || !w.On {
		return false
	}
	w.Ammo -= 1
	return true
}

func (w *Weapon) RideBike() bool {
	if w.Power <= 0 || !w.On {
		return false
	}
	w.Power -= 1
	return true
}
