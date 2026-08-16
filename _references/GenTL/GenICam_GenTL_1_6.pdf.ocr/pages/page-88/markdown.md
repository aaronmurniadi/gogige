|  GENICAM |   | ![img-124.jpeg](img-124.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

GC_ERR_ACCESS_DENIED

The access to the requested Data Stream module is denied. This may be because it is already opened by another Process but it might have other reasons as well.

GC_ERR_NOT_AVAILABLE

The sDataStreamID of the stream is generally valid but the stream is not available.

GC_ERR_NOT_IMPLEMENTED

The Producer does not implement streaming or the remote device does not provide a stream. DevGetNumDataStreams reports zero.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.4.7 DevGetParentIF

GC_ERROR DevGetParentIF ( DEV_HANDLE hDevice, IF_HANDLE * phIface )

Retrieves a handle to the parent Interface module.

##### Parameters

[in] hDevice Device module to work on.
[out] phIface Handle to the parent Interface module.

##### Returns

GC_ERR_SUCCESS Operation was successful; no error occurred.

GC_ERR_NOT_INITIALIZED No preceding call to GCInitLib.

GC_ERR_INVALID_HANDLE The handle hDevice is invalid (NULL) or does not reference an open Device module retrieved through a call to IFOpenDevice.

GC_ERR_INVALID_PARAMETER Parameter phIface is an invalid pointer (NULL or ~0x0).

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.