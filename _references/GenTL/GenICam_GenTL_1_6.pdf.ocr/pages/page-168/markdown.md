|  ![img-246.jpeg](img-246.jpeg) CAM |   | ![img-247.jpeg](img-247.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

##### 6.4.5.2 URL_INFO_CMD

enum URL_INFO_CMD

This enumeration defines commands to retrieve information with the GCGetPortURLInfo function on a module or remote device handle.

The column labeled “Impl” in the following table lists if the implementation of a given command is mandatory (M), optional (O) or conditional mandatory (CM). Mandatory means that a GenTL Producer must implement the listed command even tough it might return NA. Optional means that it is up to the implementor if a given command is implemented or not. Conditional Mandatory means that command is to be implemented if possible.

|  Enumerator | Impl | Value | Description  |
| --- | --- | --- | --- |
|  URL_INFO_URL | M | 0 | URL as defined in chapter 4.1.2Data type: STRING  |
|  URL_INFO_SCHEMA_VER_MAJOR | CM | 1 | Major version of the schema this URL refers to.Data type: INT32  |
|  URL_INFO_SCHEMA_VER_MINOR | CM | 2 | Minor version of the schema this URL refers to.Data type: INT32  |
|  URL_INFO_FILE_VER_MAJOR | CM | 3 | Major version of the XML-file this URL refers to.Data type: INT32  |
|  URL_INFO_FILE_VER_MINOR | CM | 4 | Minor version of the XML-file this URL refers to.Data type: INT32  |
|  URL_INFO_FILE_VER_SUBMINOR | CM | 5 | Subminor version of the XML-file this URL refers to.Data type: INT32  |
|  URL_INFO_FILE_SHA1_HASH | CM | 6 | SHA1 Hash of the XML-file this URL refers to. The size of the provided buffer is 160Bit according to the SHA1 specification.Data type: BUFFER  |
|  URL_INFO_FILE_REGISTER_ADDRESS | CM | 7 | Register address of the XML-File in the device's register map. In case the XML is not locally stored in the device's register map the info function should return aGC_ERR_NOT_AVAILABLE.Data type: UINT64  |