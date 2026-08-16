|  ![img-260.jpeg](img-260.jpeg)CAN |   | ![img-261.jpeg](img-261.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Member | Type | Description  |
| --- | --- | --- |
|  ChunkID | uint64_t | Numeric representation of the chunk's ChunkID.  |
|  ChunkOffset | ptrdiff_t | Offset of the chunk's data from the start of the buffer (in bytes).  |
|  ChunkLength | size_t | Size of the given chunk data (in bytes).  |

#### 6.5.1.2 DS_BUFFER_INFO_STACKED

struct DS_BUFFER_INFO_STACKED

Layout of the array elements being used in the function DSGetBufferInfoStacked to carry information about multiple buffer infos as defined in BUFFER_INFO_CMD.

|  Member | Type | Description  |
| --- | --- | --- |
|  iInfoCmd [in] | BUFFER_INFO_CMD | The buffer info to be retrieved.  |
|  iType [out] | INFO_DATATYPE | Data type of the pBuffer content as defined in the BUFFER_INFO_CMD and INFO_DATATYPE.  |
|  pBuffer [in,out] | void* | Pointer to a user allocated buffer to receive the requested information. If this parameter is NULL, iSize will contain the minimal size of pBuffer in bytes. If the iType is a string the size includes the terminating 0.  |
|  iSize [in,out] | size_t | pBuffer equal NULL: out: minimal size of pBuffer in bytes to hold all information.pBuffer unequal NULL: in: size of the provided pBuffer in bytes. out: number of bytes filled by the function.  |
|  iResult [out] | GC_ERROR | The result of the buffer info query.  |

DSGetBufferInfoStacked queries multiple buffer infos as defined in BUFFER_INFO_CMD at once.

The purpose and “direction” (in/out) of the structure data members is same as corresponding parameters of the DSGetBufferInfo function. Similar as other get-info functions, the buffer size required to hold given information can be negotiated first as described for the pBuffer==NULL above.

When retrieving multiple infos through DSGetBufferInfoStacked call, each DS_BUFFER_INFO_STACKED structure is handled independently on the others. GenTL Consumer should receive identical output (iResult and the value in pBuffer), no matter if it used single DSGetBufferInfoStacked or sequence of DSGetBufferInfo calls.