![img-4.jpeg](img-4.jpeg)

|  GenTLVersionMajor | M | All | IInteger | R | - | E | Major version number of the GenTL specification the GenTL Producer implementation complies with.  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  GenTLVersionMinor | M | All | IInteger | R | - | E | Minor version number of the GenTL specification the GenTL Producer implementation complies with.  |
|  GenTLSFNCVersionMajor | R | All | IInteger | R | - | E | Major version number of the GenTL Standard Features Naming Convention that was used to create the GenTL Producer's XML.  |
|  GenTLSFNCVersionMinor | R | All | IInteger | R | - | E | Minor version number of the GenTL Standard Features Naming Convention that was used to create the GenTL Producer's XML.  |
|  GevVersionMajor | O | GEV | IInteger | R | - | E | This feature is deprecated (See InterfaceTLVersionMajor).  |
|  GevVersionMinor | O | GEV | IInteger | R | - | E | This feature is deprecated (See InterfaceTLVersionMinor).  |

#### 2.1.2 Interface Enumeration

Contains the features related to the enumeration of available Interface modules within the System module of a GenTL Producer.

Table 2-2: Interface Enumeration Summary

|  Name | Level | TLType | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- | --- |
|  InterfaceEnumeration | R | All | ICategory | R | - | B | Category that contains all Interface Enumeration features of the System module.  |
|  InterfaceUpdateList | M | All | ICommand | (R)/W | - | B | Updates the internal list of the interfaces.  |
|  InterfaceUpdateTimeout | R | All | IInteger | R/W | ms | E | Specifies timeout for the InterfaceUpdateList Command.  |
|  InterfaceSelector | M | All | IInteger | R/W | - | B | Selector for the different GenTL Producer interfaces.  |
|  InterfaceID[InterfaceSelector] | M | All | IString | R | - | B | GenTL Producer wide unique identifier of the selected interface.  |