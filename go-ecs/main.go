// Package main
package main

import (
	"fmt"

	"github.com/klauspost/cpuid"
)

// ComponentID is the Component Identifier
type ComponentID int

// EntityID is the Entity Identifier
type EntityID int

// ArchetypeID is the Archetype Identifier
type ArchetypeID int

// ComponentType is a slice of ComponentIDs
type ComponentType []ComponentID

// Archetype is a unique collection of Component Ids that define it
type Archetype struct {
	ArchetypeID ArchetypeID
	Type        ComponentType
}

// ArchetypeSet is a set of ArchetypeIDs
type ArchetypeSet Set[ArchetypeID]

// World contains the EntityIndex and ComponentIndex
type World struct {
	EntityIndex map[EntityID]*Archetype
	ComponentIndex map[ComponentID]ArchetypeSet
}

func (p *World) HasComponent(entityID EntityID, componentID ComponentID) {
	archetype := p.EntityIndex[entityID]
	archetypeSet := p.ComponentIndex[componentID]

	return ((Set)archetypeSet).Contains()
}

func main() {
	ss := NewSS(5, 10)

	//fmt.Println(ss.capacity)
	fmt.Println(ss.capacity)

	if cpuid.CPU.AVX() {
		fmt.Println("AVX: supported")
	} else {
		fmt.Println("AVX: unsupported")
	}

}
