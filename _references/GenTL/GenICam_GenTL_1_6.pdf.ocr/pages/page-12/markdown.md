|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

## 1 Introduction

### 1.1 Purpose

The goal of the GenICam GenTL standard is to provide a generic way to enumerate devices known to a system, communicate with one or more devices and, if possible, stream data from the device to the host independent from the underlying transport technology. This allows a third party software to use different technologies to control cameras and to acquire data in a transport layer agnostic way.

The core of the GenICam GenTL standard is the definition of a generic Transport Layer Interface (TLI). This software interface between the transport technology and a third party software is defined by a C interface together with a defined behavior and a set of standardized feature names and their meaning. To access these features the GenICam GenApi module is used.

The GenICam GenApi module defines an XML description file format to describe how to access and control device features. The Standard Features Naming Convention defines the behavior of these features.

The GenTL software interface does not cover any device-specific functionality of the remote device except the one to establish communication. The GenTL provides a port to allow access to the remote device features via the GenApi module.

This makes the GenTL the generic software interface to communicate with devices and stream data from them. The combination of GenApi and GenTL provides a complete software architecture to access devices, for example cameras.

### 1.2 GenTL Subcommittee

The GenTL Subcommittee is part of the GenICam Standard Group hosted by the EMVA.

### 1.3 Acronyms and Definitions

1.3.1 Acronyms

|  Term | Description  |
| --- | --- |
|  CL | Camera Link  |
|  CTI | Common Transport Interface  |
|  GenApi | GenICam Module  |
|  GenICam | Generic Interface to Cameras  |
|  GenTL | Generic Transport Layer  |
|  GenTL SFNC | GenICam Module: GenTL Standard Features Naming Convention  |
|  SFNC | GenICam Module: Standard Features Naming Convention  |
|  PFNC | GenICam Module: Pixel Format Naming Convention  |
|  GenDC | GenICam Module: Generic Data Container  |