|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

|  Category | InterfaceInformation  |
| --- | --- |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

48-bit MAC address of this interface. Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory, and that the Representation element should be used in the XML to facilitate understanding the data.

3.2.1.10 GevInterfaceSubnetSelector

|  Name | GevInterfaceSubnetSelector  |
| --- | --- |
|  Category | InterfaceInformation  |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Selector for the subnet of this interface. Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory.

3.2.1.11 GevInterfaceSubnetIPAddress

|  Name | GevInterfaceSubnetIPAddress[GevInterfaceSubnetSelector]  |
| --- | --- |
|  Category | InterfaceInformation  |
|  Level | Mandatory  |
|  TLType | GigEVision  |