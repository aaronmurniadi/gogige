### 2.2 Interface Module

#### 2.2.1 Interface Information

Contains the features related to general information about a specific Interface module.

Table 2-5: Interface Information Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  InterfaceInformation | R | All | ICategory | R | - | E | Category that contains all Interface Information features of the Interface module.  |
|  InterfaceID | M | All | IString | R | - | E | GenTL Producer wide unique identifier of the selected interface.  |
|  InterfaceDisplayName | R | All | IString | R | - | E | A user-friendly name of the Interface.  |
|  InterfaceType | M | All | IEnumeration | R | - | E | Transport layer type of the interface.  |
|  InterfaceTLVersionMajor | M | All | IInteger | R | - | E | Major version number of the transport layer specification the GenTL Producer interface complies with.  |
|  InterfaceTLVersionMinor | M | All | IInteger | R | - | E | Minor version number of the transport layer specification the GenTL Producer interface complies with.  |
|  GevInterfaceGatewaySelector | M | GEV | IInteger | R/W | - | E | Selector for the different gateway entries for this interface.  |
|  GevInterfaceGateway[GevInterfaceGatewaySelector] | M | GEV | IInteger | R | - | E | IP address of the selected gateway entry of this interface.  |
|  GevInterfaceMACAddress | M | GEV | IInteger | R | - | E | 48-bit MAC address of this interface.  |
|  GevInterfaceSubnetSelector | M | GEV | IInteger | R/W | - | E | Selector for the subnet of this interface.  |
|  GevInterfaceSubnetIPAddress[GevInterfaceSubnetSelector] | M | GEV | IInteger | R | - | E | IP address of the selected subnet of this interface.  |
|  GevInterfaceSubnetMask[GevInterfaceS | M | GEV | IInteger | R | - | E | Subnet mask of the selected subnet of this interface.  |