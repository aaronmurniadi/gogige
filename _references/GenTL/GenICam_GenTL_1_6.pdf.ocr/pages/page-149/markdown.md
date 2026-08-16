|  ![img-210.jpeg](img-210.jpeg) CAM |   | ![img-211.jpeg](img-211.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

##### 6.4.4.5 PAYLOADTYPE_INFO_IDS

enum PAYLOADTYPE_INFO_IDS

This enumeration defines constants to give a hint on the payload type to be expected in the buffer. These values are returned by a call to DSGetBufferInfo with the command BUFFER_INFO_PAYLOADTYPE. The interpretation of the PAYLOADTYPE_INFO_IDS is depending on the TLType of the device which streams the data.

|  Enumerator | Value | Description  |
| --- | --- | --- |
|  PAYLOAD_TYPE_UNKNOWN | 0 | The GenTL Producer is not aware of the payload type of the data in the provided buffer. For the GenTL Consumer perspective this can be handled as raw data.  |
|  PAYLOAD_TYPE_IMAGE | 1 | The buffer payload contains image data. The GenTL Consumer can check if additional chunk data is available via the BUFFER_INFO commands.  |
|  PAYLOAD_TYPE_RAW_DATA | 2 | The buffer payload contains raw and further unspecified data. This can be used to send acquisition statistics.  |
|  PAYLOAD_TYPE_FILE | 3 | The buffer payload contains data of a file. It is used to transfer files such as JPEG compressed images which can be stored by the GenTL Producer directly to a hard disk. The user might get a hint how to interpret the buffer by the filename provided through a call to DSGetBufferInfo with the command BUFFER_INFO_FILENAME.  |
|  PAYLOAD_TYPE_CHUNK_DATA | 4 | The buffer payload contains chunk data which can be parsed. The chunk data type might be reported through SFNC or deduced from the technology the device is based on. This constant is for backward compatibility with GEV 1.2 and is deprecated since GenTL version 1.5. From now on ChunkData can be part or any other payload type. Use the BUFFER_INFO_CONTAINS_CHUNKDATA commads to query if a given buffer content contains chunk data.  |
|  PAYLOAD_TYPE_JPEG | 5 | The buffer payload contains JPEG data in the format described in GEV 2.0. The GenTL Producer should report additional information through the corresponding BUFFER_INFO_CMD commands.  |