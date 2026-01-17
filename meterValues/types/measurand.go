package types

// Measurand represents the type of measurement as defined in OCPP 1.6.
type Measurand string

// Alias for shorter constant names within this package.
type m = Measurand

// Measurand enumeration values as defined in OCPP 1.6.
const (
	// CurrentExport is the instantaneous current flow from EV.
	CurrentExport m = "Current.Export"
	// CurrentImport is the instantaneous current flow to EV.
	CurrentImport m = "Current.Import"
	// CurrentOffered is the maximum current offered to EV.
	CurrentOffered m = "Current.Offered"
	// EnergyActiveExportRegister is the absolute amount of active energy
	// exported (to grid).
	EnergyActiveExportRegister m = "Energy.Active.Export.Register"
	// EnergyActiveImportRegister is the absolute amount of active energy
	// imported (from grid). Default measurand.
	EnergyActiveImportRegister m = "Energy.Active.Import.Register"
	// EnergyReactiveExportRegister is the absolute amount of reactive energy
	// exported.
	EnergyReactiveExportRegister m = "Energy.Reactive.Export.Register"
	// EnergyReactiveImportRegister is the absolute amount of reactive energy
	// imported.
	EnergyReactiveImportRegister m = "Energy.Reactive.Import.Register"
	// EnergyActiveExportInterval is the active energy exported during interval.
	EnergyActiveExportInterval m = "Energy.Active.Export.Interval"
	// EnergyActiveImportInterval is the active energy imported during interval.
	EnergyActiveImportInterval m = "Energy.Active.Import.Interval"
	// EnergyReactiveExportInterval is the reactive energy exported
	// during interval.
	EnergyReactiveExportInterval m = "Energy.Reactive.Export.Interval"
	// EnergyReactiveImportInterval is the reactive energy imported
	// during interval.
	EnergyReactiveImportInterval m = "Energy.Reactive.Import.Interval"
	// Frequency is the power line frequency.
	Frequency m = "Frequency"
	// PowerActiveExport is the instantaneous active power exported by EV.
	PowerActiveExport m = "Power.Active.Export"
	// PowerActiveImport is the instantaneous active power imported by EV.
	PowerActiveImport m = "Power.Active.Import"
	// PowerFactor is the instantaneous power factor of total energy flow.
	PowerFactor m = "Power.Factor"
	// PowerOffered is the maximum power offered to EV.
	PowerOffered m = "Power.Offered"
	// PowerReactiveExport is the instantaneous reactive power exported by EV.
	PowerReactiveExport m = "Power.Reactive.Export"
	// PowerReactiveImport is the instantaneous reactive power imported by EV.
	PowerReactiveImport m = "Power.Reactive.Import"
	// RPM is the fan speed in revolutions per minute.
	RPM m = "RPM"
	// SoC is the EV battery state of charge.
	SoC m = "SoC"
	// Temperature is the temperature reading inside the charge point.
	Temperature m = "Temperature"
	// Voltage is the instantaneous AC RMS supply voltage.
	Voltage m = "Voltage"
)

// IsValid checks if the Measurand value is valid per OCPP 1.6.
func (m Measurand) IsValid() bool {
	switch m {
	case CurrentExport,
		CurrentImport,
		CurrentOffered,
		EnergyActiveExportRegister,
		EnergyActiveImportRegister,
		EnergyReactiveExportRegister,
		EnergyReactiveImportRegister,
		EnergyActiveExportInterval,
		EnergyActiveImportInterval,
		EnergyReactiveExportInterval,
		EnergyReactiveImportInterval,
		Frequency,
		PowerActiveExport,
		PowerActiveImport,
		PowerFactor,
		PowerOffered,
		PowerReactiveExport,
		PowerReactiveImport,
		RPM,
		SoC,
		Temperature,
		Voltage:
		return true
	default:
		return false
	}
}

// String returns the string representation of Measurand.
func (m Measurand) String() string {
	return string(m)
}
