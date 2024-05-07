// Package main
package main

import (
	"fmt"
	"sort"
	"strconv"

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

// ComponentTypeKey is a sorted string of each integer in the slice
type ComponentTypeKey string

// Archetype is a unique collection of Component Ids that define it
type Archetype struct {
	ArchetypeID ArchetypeID
	Type        ComponentType
}

// ArchetypeSet is a set of ArchetypeIDs
type ArchetypeSet Set[ArchetypeID]

// World contains the EntityIndex and ComponentIndex
type World struct {
	EntityIndex    map[EntityID]*Archetype
	ComponentIndex map[ComponentID]Set[int] // Set[int] = ArchetypeSet
	ArchetypeIndex map[ComponentTypeKey]Archetype
}

func (p *World) HasComponent(entityID EntityID, componentID ComponentID) bool {
	archetype := p.EntityIndex[entityID]
	archetypeSet := p.ComponentIndex[componentID]

	return archetypeSet.Contains(int(archetype.ArchetypeID))
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

func sliceKey(ints []int) string {
	// Sort the integers to ensure the same order produces the same string
	sort.Ints(ints)

	// Convert each integer to a string and concatenate them
	var str string
	for _, num := range ints {
		str += strconv.Itoa(num) + ","
	}
	// Remove the trailing comma
	str = str[:len(str)-1]

	return str
}

// // A list of component ids
// using Type = vector<ComponentId>;// Type used to store each unique component list only once
// struct Archetype {
//   ArchetypeId id;
//   Type type;
// };// Find an archetype by its list of component ids
// unordered_map<Type, Archetype> archetype_index;// Find the archetype for an entity
// unordered_map<EntityId, Archetype&> entity_index;// Find the archetypes for a component
// using ArchetypeSet = unordered_set<ArchetypeId>;
// unordered_map<ComponentId, ArchetypeSet> component_index;
