|  ![img-229.jpeg](img-229.jpeg) CAM |   | ![img-230.jpeg](img-230.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Impl | Value | Description  |
| --- | --- | --- | --- |
|   |  |  | ComponentIDValue/ChunkComponentIDValue features from SFNC). The receiver can use this tag to associate parts that convey the same purpose. In many cases, the data purpose is defined by the type of component the data represents.This information refers for example to the information provided in the GigE Vision image stream multi-part data leader.When given part is not tagged with a specific data purpose ID, the function should returnGC_ERR_NOT_AVAILABLE.Data type: UINT64  |
|  BUFFER_PART_INFO_CUSTOM_ID | O | 1000 | Starting value for GenTL Producer custom IDs which are implementation specific.If a generic GenTL Consumer is using custom BUFFER_PART_INFO_CMDs provided through a specific GenTL Producer implementation it must differentiate the handling of different GenTL Producer implementations in case other implementations will not provide that custom id or will use a different meaning with it.  |

##### 6.4.4.10 PARTDATATYPE_IDS

enum PARTDATATYPE_IDS

This enumeration defines constants to give a hint on the data type to be expected in the buffer part. These values are returned by a call to DSGetBufferPartInfo with the command BUFFER_PART_INFO_DATA_TYPE. The part data type is intended to describe data in individual parts of a multi-part buffer.

|  Enumerator | Value | Description  |
| --- | --- | --- |
|  PART_DATATYPE_UNKNOWN | 0 | The GenTL Producer is not aware of  |