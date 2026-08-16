|  ![img-143.jpeg](img-143.jpeg) CAM |   | ![img-144.jpeg](img-144.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

Returns

|  GC_ERR_SUCCESS | Operation was successful; no error occurred.  |
| --- | --- |
|  GC_ERR_NOT_INITIALIZED | No preceding call to GCInitLib.  |
|  GC_ERR_INVALID_HANDLE | The handle hDataStream is invalid (NULL) or does not reference an open Data Stream module retrieved through a call to DevOpenDataStream or the handle hBuffer is invalid (NULL) or does not reference an announced Buffer.  |
|  GC_ERR_INVALID_PARAMETER | Parameter piNumChunks is an invalid pointer (NULL or ~0x0)  |
|  GC_ERR_NO_DATA | The Buffer referenced by hBuffer does not contain chunk data or the buffer contains a standard self-described container (such as GenDC), i.e. the chunk parsing is expected to be handled by the GenTL Consumer.  |
|  GC_ERR_BUFFER_TOO_SMALL | pChunkData is not NULL and the value of *piNumChunks is too small to receive the expected amount of data.  |
|  GC_ERR_PARSING_CHUNK_DATA | An error occurred during the parsing of the chunk buffer.  |
|  GC_ERR_NOT_AVAILABLE | The request is implemented but the requested information is currently not available for any reason.  |

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.5.14 DSGetParentDev

|  GC_ERROR | DSGetParentDev | ( DS_HANDLE DEV_HANDLE * | hDataStream, phDevice )  |
| --- | --- | --- | --- |

Retrieves a handle to the parent Device module.

Parameters

|  [in] | hDataStream | Data Stream module to work on.  |
| --- | --- | --- |
|  [out] | phDevice | Handle to the parent Device module.  |

Returns

|  GC_ERR_SUCCESS | Operation was successful; no error occurred.  |
| --- | --- |
|  GC_ERR_NOT_INITIALIZED | No preceding call to GCInitLib.  |