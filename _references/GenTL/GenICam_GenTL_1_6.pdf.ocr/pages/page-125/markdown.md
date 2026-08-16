|  ![img-173.jpeg](img-173.jpeg) CAM |   | ![img-174.jpeg](img-174.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  [in] | hEvent | Event handle to parse data from.  |
| --- | --- | --- |
|  [in] | pInBuffer | Pointer to a buffer containing event data. This value must not be NULL.  |
|  [in] | iInSize | Size of the provided pInBuffer in bytes.  |
|  [in] | iInfoCmd | Information to be retrieved as defined in EVENT_DATA_INFO_CMD and EVENT_TYPE.  |
|  [out] | piType | Data type of the pOutBuffer content as defined in the EVENT_DATA_INFO_CMD, EVENT_TYPE and INFO_DATATYPE.  |
|  [in,out] | pOutBuffer | Pointer to a user allocated buffer to receive the requested information. If this parameter is NULL, piOutSize will contain the minimal size of pOutBuffer in bytes. If the piType is a string the size includes the terminating 0.  |
|  [in,out] | piOutSize | pOutBuffer equal NULL: out: minimal size of pOutBuffer in bytes to hold all information. pOutBuffer unequal NULL: in: size of the provided pOutBuffer in bytes. out: number of bytes filled by the function.  |

Returns

|  GC_ERR_SUCCESS | Operation was successful; no error occurred.  |
| --- | --- |
|  GC_ERR_NOT_INITIALIZED | No preceding call to GCInitLib.  |
|  GC_ERR_INVALID_HANDLE | The handle hEvent is invalid (NULL or ~ 0x0) or does not reference a previously registered event.  |
|  GC_ERR_NOT_IMPLEMENTED | Specified iInfoCmd is not implemented.  |
|  GC_ERR_INVALID_PARAMETER | Parameters pInBuffer, piOutSize and/or piType are invalid pointers (NULL or ~0x0) or iInSize is 0  |
|  GC_ERR_BUFFER_TOO_SMALL | pOutBuffer is not NULL and the value of *piOutSize is too small to receive the expected amount of data.  |
|  GC_ERR_NOT_AVAILABLE | The request is implemented but the requested information is currently not available for any reason.  |

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.