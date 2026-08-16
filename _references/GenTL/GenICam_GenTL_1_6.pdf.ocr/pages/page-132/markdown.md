|  ![img-181.jpeg](img-181.jpeg) CAM |   | ![img-182.jpeg](img-182.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Impl | Value | Description  |
| --- | --- | --- | --- |
|   |  |  | custom IDs which are implementation specific.If a generic GenTL Consumer is using custom TL_INFO_CMDs provided through a specific GenTL Producer implementation it must differentiate the handling of different GenTL Producer implementations in case other implementations will not provide that custom id or will use a different meaning with it.  |

#### 6.4.2 Interface Enumerations

##### 6.4.2.1 INTERFACE_INFO_CMD

enum INTERFACE_INFO_CMD

This enumeration defines commands to retrieve information with the IFGetInfo function from the Interface module or through TLGetInterfaceInfo.

The column labeled “Impl” in the following table lists if the implementation of a given command is mandatory (M) or optional (O).

|  Enumerator | Impl | Value | Description  |
| --- | --- | --- | --- |
|  INTERFACE_INFO_ID | M | 0 | Unique ID of the interface.Data type: STRING  |
|  INTERFACE_INFO_DISPLAYNAME | M | 1 | User readable name of the interface.Data type: STRING  |
|  INTERFACE_INFO_TLTYPE | M | 2 | Transport layer technology that is supported. See string constants in chapter 6.6.1.Data type: STRING  |
|  INTERFACE_INFO_CUSTOM_ID | O | 1000 | Starting value for GenTL Producer custom IDs which are implementation specific.If a generic GenTL Consumer is using custom INTERFACE_INFO_CMDs provided through a specific GenTL Producer implementation it must differentiate the handling of different GenTL Producer implementations in case other implementations will not provide that custom id or will use a different meaning with it.  |