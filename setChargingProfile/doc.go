// Package setChargingProfile implements the OCPP 1.6 SetChargingProfile
// message.
//
// # Handling Rules
//
// The Central System MAY send SetChargingProfile.req to a Charge Point to set
// a charging profile.
//
// This can occur in the following situations:
//   - At the start of a transaction to set the profile for that transaction.
//   - In a RemoteStartTransaction.req message.
//   - During an ongoing transaction to update the active profile.
//   - Outside a transaction to set a default profile for a connector, local
//     controller, or Charge Point.
//
// # General Rules
//
//   - To avoid mismatch, transactionId SHALL be included if the profile
//     applies to a specific transaction.
//
// # Setting a Charging Profile at Start of a Transaction
//
//   - Central System MAY send SetChargingProfile.req after receiving
//     StartTransaction.req/conf.
//   - It is RECOMMENDED to check the timestamp in StartTransaction.req to
//     ensure the transaction is still active.
//
// # Setting a Charging Profile in a RemoteStartTransaction Request
//
//   - ChargingProfile MAY be included. The ChargingProfilePurpose MUST be
//     TxProfile.
//   - transactionId SHALL NOT be set.
//   - Charge Point SHALL apply this profile to the newly started transaction,
//     which will receive a transactionId via StartTransaction.conf.
//   - If a SetChargingProfile.req is received with the same transactionId and
//     StackLevel, it SHALL replace the existing profile; otherwise it SHALL
//     stack next to existing profiles.
//
// # Setting a Charging Profile During a Transaction
//
//   - Central System MAY send a SetChargingProfile.req to update the active
//     profile.
//   - If a profile with the same chargingProfileId or stackLevel/Purpose
//     exists, it SHALL replace it; otherwise, it SHALL be added.
//   - Charge Point SHALL re-evaluate its collection of profiles to determine
//     the active profile.
//   - ChargingProfilePurpose MUST be TxProfile to apply only to the current
//     transaction.
//
// # Setting a Charging Profile Outside of a Transaction
//
//   - Used to set default charging profiles at any time.
//   - Profiles with the same chargingProfileId or stackLevel/Purpose SHALL
//     replace existing profiles; otherwise, they SHALL be added.
//   - Charge Point SHALL re-evaluate its collection to determine the active
//     profile.
//   - It is NOT possible to set a TxProfile without an active transaction.
//
// # Additional Notes on Profile Execution
//
//   - When refreshing a profile, set startSchedule in the past to avoid a
//     gap in behavior. Existing profiles SHALL continue until the new profile
//     is installed.
//   - If chargingSchedulePeriod > duration, the remainder periods SHALL NOT
//     be executed.
//   - If duration > chargingSchedulePeriod, the last value SHALL persist
//     until duration ends.
//   - If recurrencyKind is used and duration > recurrence period, remainder
//     periods SHALL NOT execute.
//   - The StartSchedule of the first chargingSchedulePeriod SHALL always
//     be 0.
//   - If recurrencyKind period is longer than schedule duration, Charge
//     Point SHALL fall back to default behavior or lower stackLevel profile;
//     if none available, normal charging is allowed.
package setChargingProfile
