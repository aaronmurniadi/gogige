|  ![img-133.jpeg](img-133.jpeg) CAM |   | ![img-134.jpeg](img-134.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

GC_ERR_INVALID_HANDLE

The handle hDataStream is invalid (NULL) or does not reference an open Data Stream module retrieved through a call to DevOpenDataStream.

GC_ERR_INVALID_INDEX

iIndex is greater than the number of announced buffers through calls to one of the buffer announcement functions.

GC_ERR_INVALID_PARAMETER

Parameter phBuffer is an invalid pointer (NULL or ~0x0).

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.5.7 DSGetBufferInfo

GC_ERROR DSGetBufferInfo

( DS_HANDLE hDataStream,
BUFFER_HANDLE hBuffer,
BUFFER_INFO_CMD iInfoCmd,
INFO_DATATYPE * piType,
void * pBuffer,
size_t * piSize )

Inquire information about the Buffer module associated with hBuffer on the hDataStream instance as defined in BUFFER_INFO_CMD.

To retrieve multiple infos about a buffer at once and reduce the number of calls from the GenTL Consumer to the GenTL Producer, DSGetBufferInfoStacked function can be used instead.

##### Parameters

[in] hDataStream

Data Stream module to work on.

[in] hBuffer

Buffer handle to retrieve information about.

[in] iInfoCmd

Information to be retrieved as defined in

BUFFER INFO CMD.

[out] piType

Data type of the pBuffer content as defined in the

BUFFER INFO CMD and INFO DATATYPE.

[in,out] pBuffer

Pointer to a user allocated buffer to receive the requested

information. If this parameter is NULL, piSize will contain the minimal size of pBuffer in bytes. If the piType is a string the size includes the terminating 0.

[in,out] piSize

pBuffer equal NULL:

out: minimal size of pBuffer in bytes to hold all information.

pBuffer unequal NULL:

in: size of the provided pBuffer in bytes.

out: number of bytes filled by the function.