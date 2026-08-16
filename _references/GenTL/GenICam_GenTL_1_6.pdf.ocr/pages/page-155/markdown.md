|  ![img-221.jpeg](img-221.jpeg)CAN |   | ![img-222.jpeg](img-222.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

##### 6.4.4.9 BUFFER_PART_INFO_CMD

enum BUFFER_PART_INFO_CMD

This enumeration defines commands to retrieve information with the DSGetBufferPartInfo function on a buffer handle. In case a BUFFER_PART_INFO_CMD is not available or not implemented the DSGetBufferPartInfo function must return the appropriate error return value.

The column labeled “Impl” in the following table lists if the implementation of a given command is mandatory (M), optional (O) or conditional mandatory (CM). Mandatory means that a GenTL Producer must implement the listed command even tough it might return NI or NA under certain circumstances. Optional means that it is up to the implementor if a given command is implemented or not. Conditional Mandatory means that command is to be implemented if possible.

|  Enumerator | Impl | Value | Description  |
| --- | --- | --- | --- |
|  BUFFER_PART_INFO_BASE | M | 0 | Base address of the buffer part memory. This is the address where the valid buffer part data start, not considering any padding between data parts or buffer alignment.Data type: PTR  |
|  BUFFER_PART_INFO_DATA_SIZE | M | 1 | Size of the buffer part in bytes.Actual size of the data within this buffer part should be reported.Eventual padding between buffer parts is not included. In case of variable payload type only the size of valid data within the buffer part is reported.Data type: SIZET  |
|  BUFFER_PART_INFO_DATA_TYPE | M | 2 | Type of the data in given part.This information refers to the constants defined in PARTDATATYPE_IDs.Data type: SIZET  |
|  BUFFER_PART_INFO_DATA_FORMAT | M | 3 | Format of the individual items (such as pixels) in the buffer part.The interpretation of the format is specific to every data type (BUFFER_PART_INFO_DATA_TYPE), as specified in definitions of individual PARTDATATYPE_IDS.The actual meaning of the data format  |