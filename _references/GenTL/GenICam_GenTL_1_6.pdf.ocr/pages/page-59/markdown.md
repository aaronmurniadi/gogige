|  ![img-85.jpeg](img-85.jpeg) CAM |   | ![img-86.jpeg](img-86.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

## 6 Software Interface

### 6.1 Overview

A GenTL Producer implementation is provided as a platform dependent dynamic loadable library; under Microsoft Windows platform this would be a dynamic link library (DLL). The file extension of the library is “cti” for “Common Transport Interface”.

To enable easy dynamical loading and to support a wide range of languages a C interface is defined. It is designed to be minimal and complete regarding enumeration and the access to Configuration and Signaling. This enables a quick implementation and reduces the workload on testing.

All functions defined in this chapter are mandatory and must be implemented and exported in the libraries interface even if no implementation for a function is necessary.

#### 6.1.1 Installation

In order to install a GenTL Producer an installer needs to add the path where the GenTL Producer implementation can be found to a path variable with the name GENICAM_GENTL{32/64}_PATH. The entries within the variable are separated by ‘;’ on Windows and ‘:’ on UNIX based systems. In order to allow different directories for 32Bit and 64Bit implementations residing on the same system two variables are defined: GENICAM_GENTL32_PATH for 32Bit GenTL Producer implementations and GENICAM_GENTL64_PATH for 64Bit GenTL Producer implementations. A consumer may pick the appropriate version of the environment variable.

#### 6.1.2 Function Naming Convention

All functions of the TLI follow a common naming scheme:

Prefix Operation Specifier

Entries in italics are replaced by an actual value as follows:

Table 6-3: Function naming convention

|  Entry | Description  |
| --- | --- |
|  Prefix | Specifies the handle the function works on. The handle represents the module used.Values:GC if applicable for no or all modules (GC for GenICam)TL for System module (TL for Transport Layer)IF for Interface module (IF for Interface)Dev for Device module (Dev for Device)DS for Data Stream module (DS for Data Stream)Event for Event Objects  |
|  Operation | Specifies the operation done on a certain module.Values (choice):Open to instantiate a module  |