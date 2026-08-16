|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

# Parameters

|  [in] | hDataStream | Data Stream module to work on.  |
| --- | --- | --- |
|  [in] | hBuffer | Buffer handle to queue.  |

# Returns

GC_ERR_SUCCESS Operation was successful; no error occurred.

GC_ERR_NOT_INITIALIZED No preceding call to GCInitLib.

GC_ERR_INVALID_HANDLE The handle hDataStream is invalid (NULL) or does not reference an open Data Stream module retrieved through a call to DevOpenDataStream or hBuffer is invalid (NULL) or does not reference an announced Buffer.

GC_ERR_RESOURCE_IN_USE The buffer is already queued.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.5.10 DSRevokeBuffer

|  GC_ERROR | DSRevokeBuffer | ( DS_HANDLE | hDataStream,  |
| --- | --- | --- | --- |
|   |  | BUFFER_HANDLE | hBuffer,  |
|   |  | void ** | ppBuffer,  |
|   |  | void ** | ppPrivate )  |

Removes an announced buffer from the acquisition engine. This function will free all internally allocated resources associated with this buffer. A buffer can only be revoked if it is not queued in any queue. A buffer is automatically revoked when the stream is closed. It is up to the implementation/technology if the buffer can be revoked during an ongoing acquisition if it is not queued.

Note that the ppBuffer parameter must be used with care, since it is intended only for use with buffers announced using DSAnnounceBuffer, in all other cases it returns NULL. In particular for buffers announced using DSAnnounceCompositeBuffer the GenTL Consumer has to keep track of the associated segments and their allocated memory itself. It might be therefore suitable if it keeps track of the buffer memory resources in all cases.

# Parameters

|  [in] | hDataStream | Data Stream module to work on.  |
| --- | --- | --- |
|  [in] | hBuffer | Buffer handle to revoke.  |
|  [out] | ppBuffer | Pointer to the buffer memory. This is for convenience if GenTL Consumer allocated contiguous memory (DSAnnounceBuffer) is used which is to be freed. If the buffer was allocated by the GenTL Producer (DSAllocAndAnnounceBuffer) or if the buffer is a  |