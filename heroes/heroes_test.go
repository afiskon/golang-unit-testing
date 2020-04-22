package heroes

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func heroMage() Hero {
	return Hero{"Alex", 100, 25, &MageInfo{[]Spell{THUNDERBOLT}, 100}}
}

func heroWarriorWithASword() Hero {
	return Hero{"Bob", 100, 25, &WarriorInfo{SWORD, 0}}
}

func heroWarriorWithABow() Hero {
	return Hero{"Charlie", 100, 25, &WarriorInfo{BOW, 10}}
}

func TestHeroIsDead(t *testing.T) {
	// t.Skip()
	t.Parallel()
	h := heroMage()
	require.False(t, h.IsDead())

	h.HP = 0
	require.True(t, h.IsDead())
}

// Basic test for a mage
func TestHeroIsMageIsWarrior(t *testing.T) {
	t.Parallel()
	m := heroMage()
	require.True(t, m.IsMage())
	require.False(t, m.IsWarrior())

	ws := heroWarriorWithASword()
	require.True(t, ws.IsWarrior())
	require.False(t, ws.IsMage())

	wb := heroWarriorWithABow()
	require.True(t, wb.IsWarrior())
	require.False(t, wb.IsMage())
}

// Mage attack with -10 HP effect
func TestMageAttack(t *testing.T) {
	t.Parallel()
	h := heroMage()
	m := NewCanTakeDamageMock(t)
	m.TakeDamageMock.Expect(20).Return(10)
	h.Attack(m)
	require.Equal(t, h.HP, 90)
	require.Equal(t, h.info.(*MageInfo).Mana, 95)
}

// Warrior with a sword attack with -10 HP effect
func TestWarriorWithASwrodAttack(t *testing.T) {
	t.Parallel()
	h := heroWarriorWithASword()
	m := NewCanTakeDamageMock(t)
	m.TakeDamageMock.Expect(8).Return(10)
	h.Attack(m)
	require.Equal(t, h.HP, 90)
}

// Warrior with a bow attack with -10 HP effect
func TestWarriorWithABowAttack(t *testing.T) {
	t.Parallel()
	h := heroWarriorWithABow()
	m := NewCanTakeDamageMock(t)
	m.TakeDamageMock.Expect(12).Return(10)
	h.Attack(m)
	require.Equal(t, h.HP, 90)
	require.Equal(t, h.info.(*WarriorInfo).ArrowsNumber, 9)
}

// Mage attack, no mana case
func TestMageAttackNoMana(t *testing.T) {
	t.Parallel()
	h := heroMage()
	m := NewCanTakeDamageMock(t)
	// No mana case
	h.info.(*MageInfo).Mana = 0
	// m.TakeDamageCounter = 0 // redundant since we used ExpectOnce
	h.Attack(m)
	require.Equal(t, h.HP, 100)
	require.Equal(t, h.info.(*MageInfo).Mana, 0)
	// require.Zero(t, m.TakeDamageCounter) // redundant since we used ExpectOnce
}

// Mage attack, no spells case
func TestMageAttackNoSpells(t *testing.T) {
	t.Parallel()
	h := heroMage()
	m := NewCanTakeDamageMock(t)
	// No spells case
	h.info = &MageInfo{
		Mana:      100,
		Spellbook: []Spell{},
	}
	h.Attack(m)
	require.Equal(t, h.HP, 100)
	require.Equal(t, h.info.(*MageInfo).Mana, 100)
}

// Warrior with a bow attack, no arrows case
func TestWarriorWithABowAttackNoArrows(t *testing.T) {
	t.Parallel()
	h := heroWarriorWithABow()
	h.info.(*WarriorInfo).ArrowsNumber = 0
	m := NewCanTakeDamageMock(t)
	h.Attack(m)
	require.Equal(t, h.HP, 100)
}
