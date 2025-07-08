package sharedtypes

// Package sharedtypes provides a collection of common data types and utility
// functions that are used across various Open Charge Point Protocol (OCPP) 1.6J
// messages within this library.
//
// Purpose and Importance:
// This package serves as a central repository for fundamental OCPP data structures
// such as `Integer`, `DateTime`, and various `CiString` (Case-Insensitive String)
// types. Its primary goals are:
//
// 1.  **Consistency**: By defining these types once, it ensures that all OCPP
//     message implementations (e.g., Authorize, BootNotification) use a uniform
//     representation for common data elements, preventing discrepancies and errors.
//
// 2.  **Type Safety and Validation**: Each type encapsulates its specific OCPP 1.6J
//     constraints and validation logic (e.g., maximum string lengths, printable
//     ASCII character sets, date/time formats). This proactive validation helps
//     maintain protocol compliance and reduces the likelihood of invalid data
//     being transmitted or processed.
//
// 3.  **Code Reusability**: It eliminates redundant code by providing shared
//     constructors, accessors, and validation methods for frequently used data
//     patterns, simplifying the development and maintenance of new OCPP messages.
//
// 4.  **Clarity and Readability**: Centralizing these types improves the overall
//     readability of the codebase, making it easier for developers and AI agents
//     to understand the structure and expected behavior of OCPP message fields.
//
// 5.  **Adherence to Specification**: The types defined here directly map to the
//     data types specified in the OCPP 1.6J protocol, particularly those detailed
//     in Part 2, Appendix 3: "Data Types". This strict adherence is crucial for
//     interoperability with other OCPP-compliant systems.
//
// Developers should utilize the types and functions provided in this package
// whenever an OCPP 1.6J message field corresponds to one of these common data types.

