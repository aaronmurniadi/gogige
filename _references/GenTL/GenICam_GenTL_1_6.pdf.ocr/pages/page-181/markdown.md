|  ![img-272.jpeg](img-272.jpeg)CAN |   | ![img-273.jpeg](img-273.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Name | Interface | Access | Description  |
| --- | --- | --- | --- |
|   |  |  | including its path.  |
|  TLVersion | IString | R | Vendor specific version string.  |
|  TLPath | IString | R | Full path to the GenTL Producer driver including name and extension.  |
|  TLType | IEnumeration | R | Identifies the transport layer technology of the GenTL Producer implementation. See chapter 6.6.1 for possible values.  |
|  GenTLVersionMajor | IInteger | R | Major version number of the GenTL specification the GenTL Producer implementation complies with.  |
|  GenTLVersionMinor | IInteger | R | Minor version number of the GenTL specification the GenTL Producer implementation complies with.  |

Table 7-6: Interface enumeration features

|  Name | Interface | Access | Description  |
| --- | --- | --- | --- |
|  InterfaceUpdateList | ICommand | (R)/W | Updates the internal interface list. This feature should be readable if the execution cannot performed immediately. The command then returns and the status can be polled. This function interacts with the TLUpdateInterfaceList of the GenTL Producer. It is up to the GenTL Consumer to handle access in case both methods are used.  |
|  InterfaceSelector | IInteger | R/W | Selector for the different GenTL Producer interfaces.This interface list only changes on execution of InterfaceUpdateList.The selector is 0 based in order to match the index of the C interface.  |
|  InterfaceID [InterfaceSelector] | IString | R | GenTL Producer wide unique identifier of the selected interface.This interface list only changes on execution of InterfaceUpdateList.  |

#### 7.1.2 Interface Module

All features that must be accessible in the interface module are listed here. Port functions use the IF_HANDLE to access these features. The Port access for this module is mandatory.

Table 7-7: Interface information features

|  Name | Interface | Access | Description  |
| --- | --- | --- | --- |