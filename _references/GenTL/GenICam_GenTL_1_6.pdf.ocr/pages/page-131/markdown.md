|  ![img-179.jpeg](img-179.jpeg) CAM |   | ![img-180.jpeg](img-180.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Impl | Value | Description  |
| --- | --- | --- | --- |
|   |  |  | Data type: STRING  |
|  TL_INFO_VENDOR | M | 1 | GenTL Producer vendor name.Data type: STRING  |
|  TL_INFO_MODEL | M | 2 | GenTL Producer model name.For example: A vendor produces more than one GenTL Producer for different device classes or different technologies.The TL_INFO_MODEL references a single GenTL Producer implementation. The combination of Vendor and Model provides a unique reference of ONE GenTL Producer implementation.Data type: STRING  |
|  TL_INFO_VERSION | M | 3 | GenTL Producer version.Data type: STRING  |
|  TL_INFO_TLTYPE | M | 4 | Transport layer technology that is supported. See string constants in chapter 6.6.1.Data type: STRING  |
|  TL_INFO_NAME | M | 5 | File name including extension of the library.Data type: STRING  |
|  TL_INFO_PATHNAME | M | 6 | Full path including file name and extension of the library.Data type: STRING  |
|  TL_INFO_DISPLAYNAME | M | 7 | User readable name of the GenTL Producer.Data type: STRING  |
|  TL_INFO_CHAR_ENCODING | M | 8 | The char encoding of the GenTL Producer.Data type: INT32(TL_CHAR_ENCODING_LIST enumeration value)Data type: INT32  |
|  TL_INFO_GENTL_VER_MAJOR | M | 9 | Major version number of GenTL Standard Version this Producer complies with.Data type: UINT32  |
|  TL_INFO_GENTL_VER_MINOR | M | 10 | Minor version number of GenTL Standard Version this Producer complies with.Data type: UINT32  |
|  TL_INFO_CUSTOM_ID | O | 1000 | Starting value for GenTL Producer  |