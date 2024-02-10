package labels

func Difficulty(d uint8) string {
	var difficulty = []string{
		"Normal",
		"Hard",
	}
	return difficulty[d]
}

func Month(m uint16) string {
	var months = []string{
		"",
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}
	return months[m]
}

// Rank byte to name.
func Rank(r uint16) string {
	var ranks = []string{
		"",
		"Lieutenant",
		"Captain",
		"Commander",
		"Admiral",
		"Grand Admiral",
	}
	return ranks[r]
}

func Owner(o uint8) string {
	var owners = []string{
		"",
		"Galactic Empire",
		"Garthog Empire",
		"Morgath Empire",
		"Ychom Empire",
		"Dribs Empire",
		"Sullep Empire",
		"Dargslan Kingdom",
		"Ecalep Republic",
		"Alliance of Free Traders", // humans
		"Free Nations Society",     // humans
		"Pirates",                  // humans
		"Traders",                  // humans
	}
	return owners[o]
}

// Race converts race byte to proper name.
func Race(r uint8) string {
	var races = []string{
		"",
		"Human",
		"Garthogs",
		"Morgath",
		"Ychom",
		"Dribs",
		"Sullep",
		"Dargslan",
		"Ecaleps",
	}
	return races[r]
}

func BuildingName(b uint8) string {
	var buildings = []string{
		"",
		"Colony Hub",
		"Prefab Housing",
		"Apartment Block",
		"Arcology",
		"Nuclear Plant",
		"Fusion Plant",
		"Solar Plant",
		"Water Vaporator",
		"Hydroponic Food Farm",
		"Phood(TM) Factory",
		"Spaceship Factory",
		"Equipment Factory",
		"Weapon Factory",
		"Civil Eng.Dev.Centre",
		"Mechanics Dev.Centre",
		"Computer Dev.Centre",
		"A.I. Dev. Centre",
		"Military Dev.Centre",
		"Traders' Spaceport",
		"Military Spaceport",
		"Bank",
		"Trade Centre",
		"Hospital",
		"Police Station",
		"Fire Brigade",
		"Radar Telescope",
		"Field Telescope",
		"Phased Telescope",
		"Bunker",
		"Ion Projector",
		"Plasma Projector",
		"Fusion Projector",
		"Meson Projector",
		"Inversion Shield",
		"HyperShield",
		"Barracks",
		"Fortress",
		"Stronghold",
		"Recreation Centre",
		"Park",
		"Church",
		"Bar",
		"Stadium",
	}
	return buildings[b]
}

func PlanetName(p uint8) string {
	planets := []string{
		"",
		"Outer  4",
		"Outer  3",
		"Outer  6",
		"Outer  5",
		"Outer 27",
		"Outer  8",
		"Outer 28",
		"Outer 18",
		"Outer 10",
		"Outer  2",
		"Outer 21",
		"Exterior 17",
		"Outer  7",
		"Outer 17",
		"Outer 30",
		"Outer 25",
		"Outer 12",
		"Outer  1",
		"Outer 16",
		"Outer 20",
		"Outer  9",
		"Outer 19",
		"Persol 3",
		"Outer 14",
		"Outer 11",
		"Persol 1",
		"Giantropolis",
		"Outer 22",
		"Outer 23",
		"Outer 15",
		"Outer 13",
		"Outer 26",
		"Earth",
		"Persol 2",
		"Center 6",
		"Center 5",
		"Center 13",
		"Magellan 1",
		"Magellan 2",
		"Myridan",
		"Center 14",
		"Outer 29",
		"Center 4",
		"Center 3",
		"Outer 24",
		"Edgepolis",
		"Center 7",
		"Magellan 3",
		"Center 2",
		"Center 12",
		"Center 1",
		"Center 9",
		"Center 11",
		"Magellan 6",
		"Center 8",
		"Center 15",
		"Magellan 4",
		"Magellan 7",
		"Zeuson",
		"Center 16",
		"New Caroline",
		"Magellan 8",
		"Garthog 1",
		"Centronom",
		"Magellan 5",
		"Achilles",
		"Garthog 3",
		"San Sterling",
		"Garthog 2",
		"Exterior 19",
		"Garthog 4",
		"Magellan 11",
		"Center 10",
		"Magellan 9",
		"Naxos",
		"Exterior 14",
		"Garthog 5",
		"Magellan 10",
		"Andromeda 5",
		"Exterior 15",
		"Exterior  1",
		"Exterior 13",
		"Exterior 11",
		"Exterior 21",
		"Exterior  2",
		"Andromeda 6",
		"Exterior  9",
		"Exterior  3",
		"Exterior  7",
		"Exterior 12",
		"Exterior 22",
		"Andromeda 7",
		"Andromeda 4",
		"Exterior 16",
		"Andromeda 3",
		"Exterior 19",
		"Exterior  5",
		"Exterior  8",
		"Exterior 18",
		"Exterior 20",
		"Andromeda 1",
		"Andromeda 2",
		"Exterior  6",
		"Exterior  4",
		"Exterior 10",
	}
	return planets[p]
}

// TechnologyName
// (Category and Subcategory)
// 1. Spaceships
//  1. Fighters
//  2. Cruisers
//  3. Flagships
//  4. Satellites
//  5. Space Bases
//
// 2. Equipment
//  1. Hyperdrives
//  2. Modules
//  3. Radars
//  4. Shields
//
// 3. Weapons
//  1. Lasers
//  2. Guns
//  3. Missiles/Bombs
//  4. Tanks
//  5. Vehicles
//
// 4. Buildings
// //
func TechnologyName(t uint8) string {
	technologies := []string{
		"",
		"Fighter 1",
		"Fighter 2",
		"Fighter 3",
		"Fighter 4",
		"Fighter 5",
		"Fighter 6",
		"Destroyer 1",
		"Destroyer 2",
		"Destroyer 3",
		"Cruiser 1",
		"Cruiser 2",
		"Cruiser 3",
		"Flagship 1",
		"Flagship 2",
		"Flagship 3",
		"Colonization Ship",
		"Leviathan",
		"",
		"Survey Satellite",
		"Spy Satellite",
		"Adv. Spy Satellite",
		"Hubble 2",
		"",
		"",
		"Space Base 1",
		"Orbital Factory",
		"Space Base 2",
		"Space Base 3",
		"",
		"",
		"Hyperdrive v1.0",
		"Hyperdrive v2.0",
		"Hyperdrive v3.0",
		"Hyperdrive v4.0",
		"Hyperdrive v5.0",
		"",
		"Fuzzbox ECM",
		"Shocker ECM",
		"Cargo Pod",
		"Heavy Cargo Pod",
		"",
		"",
		"Radar Array",
		"Field Array",
		"Phased Array",
		"",
		"",
		"",
		"Light Shield",
		"Medium Shield",
		"Heavy Shield",
		"Super Heavy Shield",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"Laser",
		"Pulse Laser",
		"UV Laser",
		"UV Pulse Laser",
		"",
		"",
		"Ion Gun",
		"Plasma Gun",
		"Neutron Gun",
		"Meson Gun",
		"",
		"",
		"Bomb v1.0",
		"Bomb v2.0",
		"Virus Bomb",
		"Missile v1.0",
		"Missile v2.0",
		"Mul-Head Missile",
		"Light Tank",
		"Medium Tank",
		"Heavy Tank",
		"Behemoth",
		"",
		"",
		"Radar Car",
		"Rocket Sled",
		"Heavy Rocket Sled",
		"",
		"",
		"",
		"Solar Plant",
		"Phood(TM) Factory",
		"Trade Center",
		"",
		"",
		"",
		"Inversion Shield",
		"HyperShield",
		"Fortress",
		"Stronghold",
		"Bunker",
		"",
		"Radar Telescope",
		"Field Telescope",
		"Phased Telescope",
		"",
		"",
		"",
		"Plasma Projector",
		"Fusion Projector",
		"Meson Projector",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
	}
	return technologies[t]
}
