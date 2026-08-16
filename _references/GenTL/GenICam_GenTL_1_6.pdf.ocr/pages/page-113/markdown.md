|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  GC_ERR_INVALID_HANDLE: | The handle hDataStream is invalid (NULL) or does not reference an open Data Stream module retrieved through a call to DevOpenDataStream or the handle hBuffer is invalid (NULL) or does not reference an announced Buffer.  |
| --- | --- |
|  GC_ERR_NOT_IMPLEMENTED: | Specified iInfoCmd is not implemented or the GenTL implementation does not support composite buffers.  |
|  GC_ERR_INVALID_PARAMETER: | Parameters piSize and/or piType are invalid pointers (NULL or ~0x0).  |
|  GC_ERR_INVALID_INDEX: | iSegmentIndex is greater than the number of buffer segments - 1 retrieved through a call to DSGetNumBufferSegments.  |
|  GC_ERR_BUFFER_TOO_SMALL: | pBuffer is not NULL and the value of *piSize is too small to receive the expected amount of data.  |
|  GC_ERR_NOT_AVAILABLE: | The request is implemented but the requested information is currently not available for any reason.  |

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

#### 6.3.6 Port Functions

##### 6.3.6.1 GCGetPortInfo

|  GC_ERROR | GCGetPortInfo | ( PORT_HANDLE | hPort,  |
| --- | --- | --- | --- |
|   |  | PORT_INFO_CMD | iInfoCmd,  |
|   |  | INFO_DATATYPE * | piType,  |
|   |  | void * | pBuffer,  |
|   |  | size_t * | piSize )  |

Queries detailed port information as defined in PORT_INFO_CMD.

##### Parameters

|  [in] | hPort | Module or remote device port handle to access Port from.  |
| --- | --- | --- |
|  [in] | iInfoCmd | Information to be retrieved as defined inPORT_INFO_CMD.  |
|  [out] | piType | Data type of the pBuffer content as defined in thePORT_INFO_CMD and INFO_DATATYPE.  |
|  [in,out] | pBuffer | Pointer to a user allocated buffer to receive the requested information. If this parameter is NULL, piSize will contain the minimal size of pBuffer in bytes. If the piType is a string the size includes the terminating 0.  |
|  [in,out] | piSize | pBuffer equal NULL:out: minimal size of pBuffer in bytes to hold all information.pBuffer unequal NULL:  |