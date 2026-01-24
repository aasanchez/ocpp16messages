package types

// Measurand represents the type of measurement as defined in OCPP 1.6.
type Measurand string

// Alias for shorter constant names within this package.
type m = Measurand

// Measurand enumeration values as defined in OCPP 1.6.
const (
	// MeasurandCurrentExport is the instantaneous current flow from EV.
	MeasurandCurrentExport m = "Current.Export"
	// MeasurandCurrentImport is the instantaneous current flow to EV.
	MeasurandCurrentImport m = "Current.Import"
	// MeasurandCurrentOffered is the maximum current offered to EV.
	MeasurandCurrentOffered m = "Current.Offered"
	// MeasurandEnergyActiveExportRegister is the absolute amount of active
	// energy exported (to grid).
	MeasurandEnergyActiveExportRegister m = "Energy.Active.Export.Register"
	// MeasurandEnergyActiveImportRegister is the absolute amount of active
	// energy imported (from grid). Default measurand.
	MeasurandEnergyActiveImportRegister m = "Energy.Active.Import.Register"
	// MeasurandEnergyReactiveExportRegister is the absolute amount of reactive
	// energy exported.
	MeasurandEnergyReactiveExportRegister m = "Energy.Reactive.Export.Register"
	// MeasurandEnergyReactiveImportRegister is the absolute amount of reactive
	// energy imported.
	MeasurandEnergyReactiveImportRegister m = "Energy.Reactive.Import.Register"
	// MeasurandEnergyActiveExportInterval is the active energy exported
	// during interval.
	MeasurandEnergyActiveExportInterval m = "Energy.Active.Export.Interval"
	// MeasurandEnergyActiveImportInterval is the active energy imported
	// during interval.
	MeasurandEnergyActiveImportInterval m = "Energy.Active.Import.Interval"
	// MeasurandEnergyReactiveExportInterval is the reactive energy exported
	// during interval.
	MeasurandEnergyReactiveExportInterval m = "Energy.Reactive.Export.Interval"
	// MeasurandEnergyReactiveImportInterval is the reactive energy imported
	// during interval.
	MeasurandEnergyReactiveImportInterval m = "Energy.Reactive.Import.Interval"
	// MeasurandFrequency is the power line frequency.
	MeasurandFrequency m = "Frequency"
	// MeasurandPowerActiveExport is the instantaneous active power
	// exported by EV.
	MeasurandPowerActiveExport m = "Power.Active.Export"
	// MeasurandPowerActiveImport is the instantaneous active power
	// imported by EV.
	MeasurandPowerActiveImport m = "Power.Active.Import"
	// MeasurandPowerFactor is the instantaneous power factor of total
	// energy flow.
	MeasurandPowerFactor m = "Power.Factor"
	// MeasurandPowerOffered is the maximum power offered to EV.
	MeasurandPowerOffered m = "Power.Offered"
	// MeasurandPowerReactiveExport is the instantaneous reactive power
	// exported by EV.
	MeasurandPowerReactiveExport m = "Power.Reactive.Export"
	// MeasurandPowerReactiveImport is the instantaneous reactive power
	// imported by EV.
	MeasurandPowerReactiveImport m = "Power.Reactive.Import"
	// MeasurandRPM is the fan speed in revolutions per minute.
	MeasurandRPM m = "RPM"
	// MeasurandSoC is the EV battery state of charge.
	MeasurandSoC m = "SoC"
	// MeasurandTemperature is the temperature reading inside the charge point.
	MeasurandTemperature m = "Temperature"
	// MeasurandVoltage is the instantaneous AC RMS supply voltage.
	MeasurandVoltage m = "Voltage"
)

// IsValid checks if the Measurand value is valid per OCPP 1.6.
func (m Measurand) IsValid() bool {
	switch m {
	case MeasurandCurrentExport,
		MeasurandCurrentImport,
		MeasurandCurrentOffered,
		MeasurandEnergyActiveExportRegister,
		MeasurandEnergyActiveImportRegister,
		MeasurandEnergyReactiveExportRegister,
		MeasurandEnergyReactiveImportRegister,
		MeasurandEnergyActiveExportInterval,
		MeasurandEnergyActiveImportInterval,
		MeasurandEnergyReactiveExportInterval,
		MeasurandEnergyReactiveImportInterval,
		MeasurandFrequency,
		MeasurandPowerActiveExport,
		MeasurandPowerActiveImport,
		MeasurandPowerFactor,
		MeasurandPowerOffered,
		MeasurandPowerReactiveExport,
		MeasurandPowerReactiveImport,
		MeasurandRPM,
		MeasurandSoC,
		MeasurandTemperature,
		MeasurandVoltage:
		return true
	default:
		return false
	}
}

// String returns the string representation of Measurand.
func (m Measurand) String() string {
	return string(m)
}
