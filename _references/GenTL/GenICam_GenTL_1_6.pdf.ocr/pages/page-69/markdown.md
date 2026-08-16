|  ![img-97.jpeg](img-97.jpeg) CAM |   | ![img-98.jpeg](img-98.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

##### 6.3.2.2 TLGetInfo

GC ERROR TLGetInfo

( TL_HANDLE hSystem,
TL_INFO_CMD iInfoCmd,
INFO_DATATYPE * piType,
void * pBuffer,
size_t * piSize )

Inquire information about the System module as defined in TL_INFO_CMD.

##### Parameters

[in] hSystem System module to work on.
[in] iInfoCmd Information to be retrieved as defined in TL_INFO_CMD.
[out] piType Data type of the pBuffer content as defined in the TL_INFO_CMD and INFO_DATATYPE.
[in,out] pBuffer Pointer to a user allocated buffer to receive the requested information. If this parameter is NULL, piSize will contain the minimal size of pBuffer in bytes. If the piType is a string the size includes the terminating 0.
[in,out] piSize pBuffer equal NULL:
out: minimal size of pBuffer in bytes to hold all information
pBuffer unequal NULL:
in: size of the provided pBuffer in bytes
out: number of bytes filled by the function

##### Returns

GC_ERR_SUCCESS Operation was successful; no error occurred.
GC_ERR_NOT_INITIALIZED No preceding call to GCInitLib.
GC_ERR_INVALID_HANDLE The handle hSystem is invalid (NULL) or does not reference an open System module retrieved through a call to TLOpen.
GC_ERR_NOT_IMPLEMENTED Specified iInfoCmd is not implemented.
GC_ERR_INVALID_PARAMETER Parameters piSize and/or piType are invalid pointers (NULL or ~0x0).
GC_ERR_BUFFER_TOO_SMALL pBuffer is not NULL and the value of *piSize is too small to receive the expected amount of data.
GC_ERR_NOT_AVAILABLE The request is implemented but the requested information is currently not available for any reason.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.