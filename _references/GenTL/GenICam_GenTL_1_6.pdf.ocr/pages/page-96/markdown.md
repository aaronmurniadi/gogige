|  ![img-135.jpeg](img-135.jpeg) CAM |   | ![img-136.jpeg](img-136.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

Returns

|  GC_ERR_SUCCESS | Operation was successful; no error occurred.  |
| --- | --- |
|  GC_ERR_NOT_INITIALIZED | No preceding call to GCInitLib.  |
|  GC_ERR_INVALID_HANDLE | The handle hDataStream is invalid (NULL) or does not reference an open Data Stream module retrieved through a call to DevOpenDataStream or the handle hBuffer is invalid (NULL) or does not reference an announced Buffer.  |
|  GC_ERR_NOT_IMPLEMENTED | Specified iInfoCmd is not implemented.  |
|  GC_ERR_INVALID_PARAMETER | Parameters piSize and/or piType are invalid pointers (NULL or ~0x0)  |
|  GC_ERR_BUFFER_TOO_SMALL | pBuffer is not NULL and the value of *piSize is too small to receive the expected amount of data.  |
|  GC_ERR_NO_DATA: | The buffer referenced by hBuffer contains structured data (GenDC or multi-part payload) and given iInfoCmd is not available at global buffer level (refer to BUFFER_INFO_CMD).  |
|  GC_ERR_NOT_AVAILABLE | The request is implemented but the requested information is currently not available for any reason.  |

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.5.8 DSGetInfo

|  GC_ERROR DSGetInfo | ( DS_HANDLE hDataStream,STREAM_INFO_CMD iInfoCmd,INFO_DATATYPE * piType,void * pBuffer,size_t * piSize )  |
| --- | --- |

Inquires information about the Data Stream module associated with hDataStream as defined in STREAM_INFO_CMD.

Parameters

|  [in] | hDataStream | Data Stream module to work on.  |
| --- | --- | --- |
|  [in] | iInfoCmd | Information to be retrieved as defined inSTREAM_INFO_CMD.  |
|  [out] | piType | Data type of the pBuffer content as defined in theSTREAM_INFO_CMD and INFO_DATATYPE.  |