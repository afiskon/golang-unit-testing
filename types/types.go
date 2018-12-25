package types

type Spell int

const (
	FIREBALL    Spell = iota
	THUNDERBOLT Spell = iota
)

type Weapon int

const (
	SWORD Weapon = iota
	BOW   Weapon = iota
)

type WarriorInfo struct {
	Weapon       Weapon
	ArrowsNumber int
}

type MageInfo struct {
	Spellbook []Spell
	Mana      int
}

type Hero struct {
	Name string
	HP   int
	XP   int
	info interface{}
}

//go:generate minimock -i github.com/afiskon/golang-unit-testing/types.CanTakeDamage -o . -s _mock.go
type CanTakeDamage interface {
	TakeDamage(num int) int
}

// IsDead return true if the hero has zero HP.
func (h *Hero) IsDead() bool {
	return h.HP == 0
}

// IsWarrior return true if the hero is a warrior.
func (h *Hero) IsWarrior() bool {
	switch h.info.(type) {
	case *WarriorInfo:
		return true
	default:
		return false
	}
}

// IsMage returns true if the hero is a mage.
func (h *Hero) IsMage() bool {
	switch h.info.(type) {
	case *MageInfo:
		return true
	default:
		return false
	}
}

// Attacks a given Hero
func (h *Hero) Attack(enemy CanTakeDamage) {
	if h.IsMage() {
		h.doMageAttack(enemy)
	} else if h.IsWarrior() {
		h.doWarriorAttack(enemy)
	} else {
		panic("Unknown class!")
	}
}

// TakeDamage takes the damage and returns the damage
// an attacker should take because of applied spells
func (h *Hero) TakeDamage(num int) int {
	h.HP -= num
	if h.HP < 0 {
		h.HP = 0
	}

	if (h.IsMage()) {
		// all mages are always protected
		return num / 10
	} else {
		return 0
	}
}

func (h *Hero) doMageAttack(enemy CanTakeDamage) {
	info := h.info.(*MageInfo)
	if info.Mana <= 5 {
		// there is no enough mana
		return
	}

	if len(info.Spellbook) == 0 {
		// there are no known spells
		return
	}

	info.Mana -= 5
	h.TakeDamage(enemy.TakeDamage(20))
}

func (h *Hero) doWarriorAttack(enemy CanTakeDamage) {
	info := h.info.(*WarriorInfo)
	if info.Weapon == BOW {
		if info.ArrowsNumber > 0 {
			// attack using a bow
			h.TakeDamage(enemy.TakeDamage(12))
			info.ArrowsNumber--
		}
	} else if info.Weapon == SWORD {
		h.TakeDamage(enemy.TakeDamage(8))
	} else {
		panic("unknown weapon")
	}
}
