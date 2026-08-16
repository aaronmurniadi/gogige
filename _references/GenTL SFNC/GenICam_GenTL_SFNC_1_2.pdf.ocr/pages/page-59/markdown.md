|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 3.2.1.7 GevInterfaceGatewaySelector

|  Name | GevInterfaceGatewaySelector  |
| --- | --- |
|  Category | InterfaceInformation  |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | ≥0  |

Selector for the different gateway entries for this interface. The selector is 0-based. Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory.

### 3.2.1.8 GevInterfaceGateway

|  Name | GevInterfaceGateway[GevInterfaceGatewaySelector]  |
| --- | --- |
|  Category | InterfaceInformation  |
|  Level | Mandatory  |
|  TLType | GigEVision  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values |   |

IP address of the selected gateway entry of this interface. Note that for a GenTL Producer implementation supporting GigE Vision this feature is mandatory.

### 3.2.1.9 GevInterfaceMACAddress

|  Name | GevInterfaceMACAddress  |
| --- | --- |