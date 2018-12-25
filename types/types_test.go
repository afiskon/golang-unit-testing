package types

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func heroMage() Hero {
	return Hero{"Alex", 100, 25, &MageInfo{[]Spell{THUNDERBOLT}, 100}}
}

func heroWariorWithASword() Hero {
	return Hero {"Bob", 100, 25, &WariorInfo{ SWORD, 0 }}
}

func heroWariorWithABow() Hero {
	return Hero {"Charlie", 100, 25, &WariorInfo{ BOW, 10 }}
}

func TestHeroIsDead(t *testing.T) {
	h := heroMage()
	require.False(t, h.IsDead())

	h.HP = 0
	require.True(t, h.IsDead())
}

// Basic test for a mage
func TestHeroIsMageIsWarior(t *testing.T) {
	m := heroMage()
	require.True(t, m.IsMage())
	require.False(t, m.IsWarior())

	ws := heroWariorWithASword()
	require.True(t, ws.IsWarior())
	require.False(t, ws.IsMage())

	wb := heroWariorWithABow()
	require.True(t, wb.IsWarior())
	require.False(t, wb.IsMage())
}

// Mage attack with -10 HP effect
func TestMageAttack(t *testing.T) {
	h := heroMage()
	m := NewCanTakeDamageMock(t)
	m.TakeDamageMock.ExpectOnce(20).Return(10)
	h.Attack(m)
	require.Equal(t, h.HP, 90)
	require.Equal(t, h.info.(*MageInfo).Mana, 95)
}

// Warior with a sword attack with -10 HP effect
func TestWariorWithASwrodAttack(t *testing.T) {
	h := heroWariorWithASword()
	m := NewCanTakeDamageMock(t)
	m.TakeDamageMock.ExpectOnce(8).Return(10)
	h.Attack(m)
	require.Equal(t, h.HP, 90)
}

// Warior with a bow attack with -10 HP effect
func TestWariorWithABowAttack(t *testing.T) {
	h := heroWariorWithABow()
	m := NewCanTakeDamageMock(t)
	m.TakeDamageMock.ExpectOnce(12).Return(10)
	h.Attack(m)
	require.Equal(t, h.HP, 90)
	require.Equal(t, h.info.(*WariorInfo).ArrowsNumber, 9)
}

// Mage attack, no mana case
func TestMageAttackNoMana(t *testing.T) {
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
	h := heroMage()
	m := NewCanTakeDamageMock(t)
	// No spells case
	h.info = &MageInfo{
		Mana: 100,
		Spellbook: []Spell{},
	}
	h.Attack(m)
	require.Equal(t, h.HP, 100)
	require.Equal(t, h.info.(*MageInfo).Mana, 100)
}

// Warrior with a bow attack, no arrows case
func TestWariorWithABowAttackNoArrows(t *testing.T) {
	h := heroWariorWithABow()
	h.info.(*WariorInfo).ArrowsNumber = 0
	m := NewCanTakeDamageMock(t)
	h.Attack(m)
	require.Equal(t, h.HP, 100)
}
