package vehicle

import "fmt"

type CollisionDamage struct {
	SidedDice    int
	DiceCount    int
	TypeOfDamage string
}

type Vehicle struct {
	Level             int
	Price             int
	Speed             int
	EAC               int
	KAC               int
	HP                int
	Hardness          int
	CollisionDC       int
	ModificationSlots int
	CollisionDamage
}

func (v *Vehicle) Crash(other *Vehicle) {
}

func NewVehicle(level int) *Vehicle {
	switch level {
	case 1:
		return &Vehicle{Level: 1, Price: 700, Speed: 25, EAC: 12, KAC: 14, HP: 12, Hardness: 5, CollisionDC: 10, ModificationSlots: 1,
			CollisionDamage: CollisionDamage{SidedDice: 4, DiceCount: 4, TypeOfDamage: "B"}}
	case 2:
		return &Vehicle{Level: 2, Price: 1800, Speed: 25, EAC: 13, KAC: 15, HP: 20, Hardness: 5, CollisionDC: 11, ModificationSlots: 1}
	case 3:
		return &Vehicle{Level: 3, Price: 2600, Speed: 25, EAC: 14, KAC: 16, HP: 30, Hardness: 5, CollisionDC: 12, ModificationSlots: 1}
	case 4:
		return &Vehicle{Level: 4, Price: 4000, Speed: 25, EAC: 15, KAC: 17, HP: 40, Hardness: 5, CollisionDC: 13, ModificationSlots: 1}
	case 5:
		return &Vehicle{Level: 5, Price: 8000, Speed: 25, EAC: 17, KAC: 19, HP: 55, Hardness: 5, CollisionDC: 13, ModificationSlots: 2}
	case 6:
		return &Vehicle{Level: 6, Price: 8500, Speed: 30, EAC: 18, KAC: 20, HP: 75, Hardness: 5, CollisionDC: 14, ModificationSlots: 2}
	case 7:
		return &Vehicle{Level: 7, Price: 1300, Speed: 30, EAC: 19, KAC: 21, HP: 95, Hardness: 7, CollisionDC: 15, ModificationSlots: 2}
	case 8:
		return &Vehicle{Level: 8, Price: 19000, Speed: 30, EAC: 20, KAC: 22, HP: 120, Hardness: 8, CollisionDC: 16, ModificationSlots: 2}
	case 9:
		return &Vehicle{Level: 9, Price: 27000, Speed: 35, EAC: 21, KAC: 23, HP: 135, Hardness: 9, CollisionDC: 16, ModificationSlots: 2}
	case 10:
		return &Vehicle{Level: 10, Price: 36000, Speed: 35, EAC: 23, KAC: 25, HP: 150, Hardness: 10, CollisionDC: 17, ModificationSlots: 3}
	case 11:
		return &Vehicle{Level: 11, Price: 50000, Speed: 35, EAC: 24, KAC: 26, HP: 165, Hardness: 11, CollisionDC: 18, ModificationSlots: 3}
	case 12:
		return &Vehicle{Level: 12, Price: 73000, Speed: 40, EAC: 25, KAC: 27, HP: 185, Hardness: 12, CollisionDC: 19, ModificationSlots: 3}
	case 13:
		return &Vehicle{Level: 13, Price: 100000, Speed: 40, EAC: 26, KAC: 28, HP: 205, Hardness: 13, CollisionDC: 19, ModificationSlots: 3}
	case 14:
		return &Vehicle{Level: 14, Price: 150000, Speed: 40, EAC: 27, KAC: 29, HP: 230, Hardness: 14, CollisionDC: 20, ModificationSlots: 3}
	case 15:
		return &Vehicle{Level: 15, Price: 230000, Speed: 45, EAC: 29, KAC: 31, HP: 255, Hardness: 15, CollisionDC: 21, ModificationSlots: 4}
	case 16:
		return &Vehicle{Level: 16, Price: 345000, Speed: 45, EAC: 30, KAC: 32, HP: 280, Hardness: 16, CollisionDC: 22, ModificationSlots: 4}
	case 17:
		return &Vehicle{Level: 17, Price: 520000, Speed: 45, EAC: 31, KAC: 33, HP: 310, Hardness: 17, CollisionDC: 22, ModificationSlots: 4}
	case 18:
		return &Vehicle{Level: 18, Price: 760000, Speed: 50, EAC: 32, KAC: 34, HP: 340, Hardness: 18, CollisionDC: 23, ModificationSlots: 4}
	case 19:
		return &Vehicle{Level: 19, Price: 1150000, Speed: 50, EAC: 33, KAC: 35, HP: 370, Hardness: 19, CollisionDC: 24, ModificationSlots: 4}
	case 20:
		return &Vehicle{Level: 20, Price: 1750000, Speed: 50, EAC: 35, KAC: 37, HP: 400, Hardness: 29, CollisionDC: 25, ModificationSlots: 5}
	default:
		return nil

	}
}

func DumpVehicle(v Vehicle) {
	fmt.Println("Level:", v.Level)
	fmt.Println("Price:", v.Price)
	fmt.Println("Speed:", v.Speed)
	fmt.Println("EAC:", v.EAC)
	fmt.Println("KAC:", v.KAC)
	fmt.Println("HP:", v.HP)
	fmt.Println("Hardness:", v.Hardness)
	fmt.Println("CollisionDC:", v.CollisionDC)
	fmt.Println("ModificationSlots:", v.ModificationSlots)
	fmt.Println("CollisionDamage.SidedDice:", v.SidedDice)
	fmt.Println("CollisionDamage.DiceCount", v.DiceCount)
	fmt.Println("CollisionDamage.TypeOfDamage", v.TypeOfDamage)
}
