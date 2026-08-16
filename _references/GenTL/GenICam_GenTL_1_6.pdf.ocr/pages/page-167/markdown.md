|  ![img-244.jpeg](img-244.jpeg) CAM |   | ![img-245.jpeg](img-245.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Impl | Value | Description  |
| --- | --- | --- | --- |
|   |  |  | “TLDataStream” for the Data Stream module.“TLBuffer” for the Buffer module.“Device” for the remote device.Data type: STRING  |
|  PORT_INFO_LITTLE_ENDIAN | M | 5 | Flag indicating that the port's data is little endian.Data type: BOOL8  |
|  PORT_INFO_BIG_ENDIAN | M | 6 | Flag indicating that the port's data is big endian.Data type: BOOL8  |
|  PORT_INFO_ACCESS_READ | M | 7 | Flag indicating that read access is allowed.Data type: BOOL8  |
|  PORT_INFO_ACCESS_WRITE | M | 8 | Flag indicating that write access is allowed.Data type: BOOL8  |
|  PORT_INFO_ACCESS_NA | M | 9 | Flag indicating that the port is currently not available.Data type: BOOL8  |
|  PORT_INFO_ACCESS_NI | M | 10 | Flag indicating that no port is implemented. This is only valid on the Buffer module since on all other modules the port is mandatory.Data type: BOOL8  |
|  PORT_INFO_VERSION | M | 11 | Version of the port.Data type: STRING  |
|  PORT_INFO_PORTNAME | M | 12 | Name of the port as referenced in the XML description. This name is used to connect this port to the nodemap instance of this module.Data type: STRING  |
|  PORT_INFO_CUSTOM_ID | O | 1000 | Starting value for GenTL Producer custom IDs which are implementation specific.If a generic GenTL Consumer is using custom PORT_INFO_CMDs provided through a specific GenTL Producer implementation it must differentiate the handling of different GenTL Producer implementations in case other implementations will not provide that custom id or will use a different meaning with it.  |