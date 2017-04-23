package main

func MakeNationSumer() {
	n := NewNation("Kingdom of Sumer", "SMR", nil, 255, 0, 0)
	Nations[n.Tag] = n
	Provences["GLG"].Owner = n
	Provences["URK"].Owner = n
	Provences["ENK"].Owner = n
	Provences["NNL"].Owner = n
	Provences["AN"].Owner = n
	Provences["EUP"].Owner = n
	Provences["EAM"].Owner = n
	Provences["TIG"].Owner = n
	Provences["EAN"].Owner = n
}

func MakeNationNorden() {
	n := NewNation("Norden Republic", "NRD", nil, 0, 0, 255)
	Nations[n.Tag] = n
	Provences["RNK"].Owner = n
	Provences["WTM"].Owner = n
	Provences["QLL"].Owner = n
	Provences["ARB"].Owner = n
}

func InitNations() {
	Nations = make(map[string]*Nation)
	MakeNationSumer()
	MakeNationNorden()
}