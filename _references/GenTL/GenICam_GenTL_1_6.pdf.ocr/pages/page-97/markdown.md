|  ![img-137.jpeg](img-137.jpeg) CAM |   | ![img-138.jpeg](img-138.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

[in,out] pBuffer

Pointer to a user allocated buffer to receive the requested information. If this parameter is NULL, piSize will contain the minimal size of pBuffer in bytes. If the piType is a string the size includes the terminating 0.

[in,out] piSize

pBuffer equal NULL:
out: minimal size of pBuffer in bytes to hold all information.
pBuffer unequal NULL:
in: size of the provided pBuffer in bytes.
out: number of bytes filled by the function.

##### Returns

GC_ERR_SUCCESS

Operation was successful; no error occurred.

GC_ERR_NOT_INITIALIZED

No preceding call to GCInitLib.

GC_ERR_INVALID_HANDLE

The handle hDataStream is invalid (NULL) or does not reference an open Data Stream module retrieved through a call to DevOpenDataStream.

GC_ERR_NOT_IMPLEMENTED

Specified iInfoCmd is not implemented.

GC_ERR_INVALID_PARAMETER

Parameters piSize and/or piType are invalid pointers (NULL or ~0x0)

GC_ERR_BUFFER_TOO_SMALL

pBuffer is not NULL and the value of *piSize is too small to receive the expected amount of data.

GC_ERR_NOT_AVAILABLE

The request is implemented but the requested information is currently not available for any reason.

GC_ERR_NO_DATA

The request is implemented but the requested query is not applicable in current configuration (refer to guidelines provided with individual info commands).

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

#### 6.3.5.9 DSQueueBuffer

GC_ERROR

DSQueueBuffer

( DS_HANDLE

hDataStream,

BUFFER_HANDLE

hBuffer )

This function queues a particular buffer for acquisition. A buffer can be queued for acquisition any time after the buffer was announced (before or after the acquisition has been started) if it is not currently queued. Furthermore, a buffer which is already waiting to be delivered cannot be queued for acquisition. A queued buffer cannot be revoked.

The order of the delivered buffers is not necessarily the same as the order in which they have been queued.