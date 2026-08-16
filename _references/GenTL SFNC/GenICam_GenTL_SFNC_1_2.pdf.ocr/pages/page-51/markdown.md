|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.1.2.6 InterfaceDisplayName

|  Name | InterfaceDisplayName[InterfaceSelector]  |
| --- | --- |
|  Category | InterfaceEnumeration  |
|  Level | Recommended  |
|  Interface | IString  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Any NULL-terminated string  |

A user-friendly name of the Interface.

Corresponds to the TLGetInterfaceID function with the index corresponding to "InterfaceSelector".

### 3.1.2.7 GevInterfaceMACAddress

|  Name | GevInterfaceMACAddress[InterfaceSelector]  |
| --- | --- |
|  Category | InterfaceEnumeration  |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

48-bit MAC address of the selected interface. Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory, and that the Representation element should be used in the XML to facilitate understanding the data.

### 3.1.2.8 GevInterfaceDefaultIPAddress

|  Name | GevInterfaceDefaultIPAddress[InterfaceSelector]  |
| --- | --- |
|  Category | InterfaceEnumeration  |