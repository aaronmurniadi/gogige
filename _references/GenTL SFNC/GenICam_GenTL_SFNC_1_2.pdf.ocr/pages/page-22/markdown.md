![img-5.jpeg](img-5.jpeg)

|  InterfaceDisplayName[InterfaceSelector] | R | All | IString | R | - | B | A user-friendly name of the Interface.  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  GevInterfaceMACAddress[InterfaceSelector] | M | GEV | IInteger | R | - | E | 48-bit MAC address of the selected interface.  |
|  GevInterfaceDefaultIPAddress[InterfaceSelector] | M | GEV | IInteger | R | - | E | IP address of the first subnet of the selected interface.  |
|  GevInterfaceDefaultSubnetMask[InterfaceSelector] | M | GEV | IInteger | R | - | E | Subnet mask of the first subnet of the selected interface.  |
|  GevInterfaceDefaultGateway[InterfaceSelector] | R | GEV | IInteger | R | - | E | Gateway of the selected interface.  |

#### 2.1.3 GenICam Control

Contains the features related to GenICam control and access of the System module.

Table 2-3: GenICam Control Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  Root | M | All | ICategory | R | - | B | Provides the Root of the GenICam features tree.  |
|  TLPort | M | All | IPort | R/W | - | I | The GenICam port through which the System module is accessed.  |