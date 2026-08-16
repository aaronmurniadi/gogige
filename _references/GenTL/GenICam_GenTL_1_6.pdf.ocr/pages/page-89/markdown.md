|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

#### 6.3.5 Data Stream Functions

##### 6.3.5.1 DSAllocAndAnnounceBuffer

GC_ERROR DSAllocAndAnnounceBuffer ( DS_HANDLE hDataStream, size_t iBufferSize, void * pPrivate, BUFFER_HANDLE * phBuffer )

This function allocates the memory for a single buffer and announces this buffer to the Data Stream associated with the hDataStream handle and returns a buffer handle which references that single buffer until the buffer is revoked. This will allocate internal resources which will be freed upon a call to DSRevokeBuffer.

Announcing a buffer to a data stream does not mean that this buffer will be automatically queued for acquisition. This is done through a separate call to DSQueueBuffer.

The memory referenced in this buffer must stay valid until a buffer is revoked with DSRevokeBuffer.

Every call of this function should be matched with a call of DSRevokeBuffer even though the resources are also freed when the module is closed.

Refer to chapter 5.2.1 in order to determine the right buffer size.

Parameters

|  [in] | hDataStream | Data Stream module to work on.  |
| --- | --- | --- |
|  [in] | iBufferSize | Size of the buffer in bytes.  |
|  [in] | pPrivate | Pointer to private data which will be passed to the GenTL Consumer on New Buffer events. This parameter may be NULL.  |
|  [out] | phBuffer | Buffer module handle of the newly announced buffer. It is recommended to initialize *phBuffer to GENTL_INVALID_HANDLE before calling DSAllocAndAnnounceBuffer to indicate an invalid handle.  |

Returns

|  GC_ERR_SUCCESS | Operation was successful; no error occurred.  |
| --- | --- |
|  GC_ERR_NOT_INITIALIZED | No preceding call to GCInitLib.  |
|  GC_ERR_INVALID_HANDLE | The handle hDataStream is invalid (NULL) or does not reference an open Data Stream module retrieved through a call to DevOpenDataStream.  |
|  GC_ERR_INVALID_PARAMETER | Parameter phBuffer is an invalid pointer (NULL or ~0x0).  |