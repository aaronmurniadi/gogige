|  ![img-125.jpeg](img-125.jpeg)CAN |   | ![img-126.jpeg](img-126.jpeg)  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

GC_ERR_BUSY

The acquisition has been started and the GenTL Producer does not support announcing buffers while the acquisition is active.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.5.2 DSAnnounceBuffer

GC_ERROR DSAnnounceBuffer ( DS_HANDLE hDataStream,
    void * pBuffer,
    size_t iSize,
    void * pPrivate,
    BUFFER_HANDLE * phBuffer )

This announces a GenTL Consumer allocated memory to the Data Stream associated with the hDataStream handle and returns a buffer handle which references that single buffer until the buffer is revoked. This will allocate internal resources which will be freed upon a call to DSRevokeBuffer.

Announcing a buffer to a data stream does not mean that this buffer will be automatically queued for acquisition. This is done through a separate call to DSQueueBuffer.

The memory referenced in pBuffer must stay valid until the buffer is revoked with DSRevokeBuffer. Every call of this function must be matched with a call of DSRevokeBuffer.

A buffer can only be announced once to a given stream. If a GenTL Consumer tries to announce an already announced buffer the function will return the error GC_ERR_RESOURCE_IN_USE. A buffer may additionally be announced to one or more other data stream(s) which will then result in one or more additional handles. The Consumer needs to take care about synchronisation between these streams.

Refer to chapter 5.2.1 in order to determine the right buffer size.

Parameters

|  [in] | hDataStream | Data Stream module to work on.  |
| --- | --- | --- |
|  [in] | pBuffer | Pointer to buffer memory to announce.  |
|  [in] | iSize | Size of the pBuffer in bytes.  |
|  [in] | pPrivate | Pointer to private data which will be passed to the GenTL Consumer on New Buffer events. This parameter may be NULL.  |
|  [out] | phBuffer | Buffer module handle of the newly announced buffer. It is recommended to initialize *phBuffer to GENTL_INVALID_HANDLE before calling DSAnnounceBuffer to indicate an invalid handle.  |