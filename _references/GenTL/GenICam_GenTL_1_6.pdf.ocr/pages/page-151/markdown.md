|  ![img-214.jpeg](img-214.jpeg) CAM |   | ![img-215.jpeg](img-215.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Value | Description  |
| --- | --- | --- |
|   |  | format values is referencing GigE Vision 1.x.  |
|  PIXELFORMAT_NAMESPACE_IIDC | 2 | The interpretation of the pixel format values is referencing IIDC 1.x.  |
|  PIXELFORMAT_NAMESPACE_PFNC_16BIT | 3 | The interpretation of the pixel format values is referencing PFNC 16Bit ValuesIt is recommended to use the PFNC32 namespace when ever possible or even do the translation in the GenTL Producer since the support in GenTL consumers for it is expected to be much broader.  |
|  PIXELFORMAT_NAMESPACE_PFNC_32BIT | 4 | The interpretation of the pixel format values is referencing PFNC 32Bit Values.  |
|  PIXELFORMAT_NAMESPACE_CUSTOM_ID | 1000 | The interpretation of the pixel format values is GenTL Producer specific.  |

##### 6.4.4.7 PIXELENDIANNESS_IDS

enum PIXELENDIANNESS_IDS

This enumeration defines constants describing endianness of multi-byte pixel data in a buffer. These values are returned by a call to DSGetBufferInfo with the command BUFFER_INFO_PIXELENDIANNESS.

|  Enumerator | Value | Description  |
| --- | --- | --- |
|  PIXELENDIANNESS_UNKNOWN | 0 | Endianness of the pixel data is unknown to the GenTL Producer.  |
|  PIXELENDIANNESS_LITTLE | 1 | The pixel data is stored in little endian format.  |
|  PIXELENDIANNESS_BIG | 2 | The pixel data is stored in big endian format.  |

##### 6.4.4.8 STREAM_INFO_CMD

enum STREAM_INFO_CMD

This enumeration defines commands to retrieve information with the DSGetInfo function on a data stream handle.