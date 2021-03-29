

interface IF {}

type F1 struct {}
type F2 struct {}

factoryF(nameF) {
	switch nameF {
	case "F1":
		return F1{B: factoryB(), A: factoryA()}
	case "F2":
		return F2{B: factoryB(), A: factoryA()}
	}
}

interface IE {}

type E1 struct {}
type E2 struct {}

factoryE(nameE) {
	switch nameE {
	case "E1":
		return E1{}
	case "E2":
		return E2{}
	}
}

interface ID {}

type D1 struct {}
type D2 struct {}

factoryD(nameD) {
	switch nameD {
	case "D1":
		return D1{}
	case "D2":
		return D2{}
	}
}

interface IC {}

type C1 struct {}
type C2 struct {}

factoryC(nameC) {
	switch nameC {
	case "C1":
		return C1{}
	case "C2":
		return C2{}
	}
}

interface IB {}

type B1 struct {}
type B2 struct {}

factoryB(nameB) {
	switch nameB {
	case "B1":
		return B1{}
	case "B2":
		return B2{}
	}
}

interface IA {}

type A1 struct {}
type A2 struct {}

factoryA(nameA) {
	switch nameA {
	case "A1":
		return A1{}
	case "A2":
		return A2{}
	}
}