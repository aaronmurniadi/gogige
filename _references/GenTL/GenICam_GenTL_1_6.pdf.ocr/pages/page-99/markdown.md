|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

composite buffer (DSAnnounceCompositeBuffer)
NULL is to be returned. If the parameter is set to NULL it is ignored.
[out] ppPrivate Pointer to the user data pointer given in the buffer announcement function. If the parameter is set to NULL it is ignored.

# Returns

GC_ERR_SUCCESS Operation was successful; no error occurred.

GC_ERR_NOT_INITIALIZED No preceding call to GCInitLib.

GC_ERR_INVALID_HANDLE The handle hDataStream is invalid (NULL) or does not reference an open Data Stream module retrieved through a call to DevOpenDataStream or hBuffer is invalid (NULL) or does not reference an announced Buffer.

GC_ERR_BUSY The buffer is currently queued and can not be revoked or the GenTL Consumer tried to revoke the buffer while the acquisition was in progress and the implementation or the underlying technology would not allow it.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.5.11 DSStartAcquisition

GC_ERROR DSStartAcquisition ( DS_HANDLE hDataStream, ACQ_START_FLAGS iStartFlags, uint64_t iNumToAcquire )

Starts the acquisition engine on the host. Each call to DSStartAcquisition must be accompanied by a call to DSStopAcquisition.

# Parameters

[in] hDataStream Data Stream module to work on.
[in] iStartFlags As defined in ACQ START FLAGS.
[in] iNumToAcquire Sets the number of filled/delivered buffers after which the acquisition engine stops automatically. Buffers which are internally discarded or missed are not counted. If set to GENTL_INFINITE the acquisition continues until a call to DSStopAcquisition is issued. If set to 0 a GC_ERR_INVALID_PARAMETER is returned.

# Returns