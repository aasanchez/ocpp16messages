// GetConfiguration handling rules:
//
// The Central System SHALL request configuration settings from a Charge Point
// by sending GetConfiguration.req.
//
// If the request contains no keys or the keys list is missing, the Charge Point
// SHALL return all configuration settings in GetConfiguration.conf.
//
// If specific keys are requested, the Charge Point SHALL return the recognized
// keys with their values and read-only status. Unrecognized keys SHALL be listed
// in the optional unknown key list element of the response.
//
// The number of keys that can be requested in a single PDU MAY be limited by
// the Charge Point. The maximum number of keys can be retrieved via the
// GetConfigurationMaxKeys configuration key.
package getConfiguration
