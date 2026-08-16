|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

- Retrieval of the correct GenICam XML file: for the device configuration XML there is no unique way a GenTL Producer can create a node map that will be always identical to the one used by the application. Even if in most cases the XML is retrieved from the device, it cannot be assumed that it will always be the case.
- GenICam XML description implementation: there is no standardized implementation. The GenApi is only a reference implementation, not a mandatory standard. User implementations in the same or in a different language may be used to interpret GenApi XML files. Even if the same implementation is used, the GenTL Producer and Consumer may not even use the same version of the implementation.
- Caching: when using another instance of an XML description inside the GenTL Producer, unwanted cache behavior may occur because both instances will be maintaining their own local, disconnected caches.

#### 2.1.3 GenICam GenTL SFNC

In order to allow configuration of a GenTL Producer each module implements a virtual register map and provides a GenApi compliant XML file (see chapter 2.3.2). Only mandatory features of these XML files are described in this document in chapter 7. All features (mandatory and non-mandatory) are defined in the GenTL SFNC document.

### 2.2 GenTL Modules

The GenTL standard defines a layered structure for libraries implementing the GenTL Interface. Each layer is defined in a module. The modules are presented in a tree structure with the System module as its root.