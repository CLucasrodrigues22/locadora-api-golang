package enums

type FuelType string

const (
	FuelGasoline FuelType = "gasoline"
	FuelEthanol  FuelType = "ethanol"
	FuelFlex     FuelType = "flex"
	FuelDiesel   FuelType = "diesel"
	FuelHybrid   FuelType = "hybrid"
	FuelElectric FuelType = "electric"
)

func (f FuelType) IsValid() bool {
	return isValidEnum(f, []FuelType{
		FuelGasoline,
		FuelEthanol,
		FuelFlex,
		FuelDiesel,
		FuelHybrid,
		FuelElectric,
	})
}

type GearboxType string

const (
	GearboxManual    GearboxType = "manual"
	GearboxAutomatic GearboxType = "automatic"
)

func (g GearboxType) IsValid() bool {
	return isValidEnum(g, []GearboxType{
		GearboxManual,
		GearboxAutomatic,
	})
}

type VehicleStatus string

const (
	StatusAvailable   VehicleStatus = "available"
	StatusRented      VehicleStatus = "rented"
	StatusMaintenance VehicleStatus = "maintenance"
)

func (s VehicleStatus) IsValid() bool {
	return isValidEnum(s, []VehicleStatus{
		StatusAvailable,
		StatusRented,
		StatusMaintenance,
	})
}

func isValidEnum[T ~string](value T, allowed []T) bool {
	for _, v := range allowed {
		if value == v {
			return true
		}
	}
	return false
}
