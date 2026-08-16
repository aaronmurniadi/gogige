|  GEN<ICAM |   | ![img-119.jpeg](img-119.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

in: size of the provided sDataStreamID in bytes

out: number of bytes filled by the function

Returns

|  GC_ERR_SUCCESS | Operation was successful; no error occurred.  |
| --- | --- |
|  GC_ERR_NOT_INITIALIZED | No preceding call to GCInitLib.  |
|  GC_ERR_INVALID_HANDLE | The handle hDevice is invalid (NULL) or does not reference an open Device module retrieved through a call to IFOpenDevice.  |
|  GC_ERR_NOT_IMPLEMENTED | The Producer does not implement streaming or the remote device does not provide a stream. DevGetNumDataStreams reports zero.  |
|  GC_ERR_INVALID_INDEX | iIndex is greater than the number of available Data Stream modules - 1 retrieved through a call to DevGetNumDataStreams.  |
|  GC_ERR_INVALID_PARAMETER | Parameter piSize is an invalid pointer (NULL or ~0x0).  |
|  GC_ERR_BUFFER_TOO_SMALL | sDataStreamID is not NULL and the value of *piSize is too small to receive the expected amount of data.  |

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.4.4 DevGetNumDataStreams

GC_ERROR DevGetNumDataStreams(DEV_HANDLE hDevice, uint32_t * piNumDataStreams)

Queries the number of available data streams on this Device module.

For GenTL Producers which do not provide a data stream the number of available data streams is zero. Calls to DevGetDataStreamID or DevOpenDataStream will fail with GC_ERR_NOT_IMPLEMENTED. Nevertheless a GenTL Producer must export all functions of the public interface

##### Parameters

[in] hDevice Device module to work on.
[out] piNumDataStreams Number of data stream on this Device module.

##### Returns

GC_ERR_SUCCESS Operation was successful; no error occurred.
GC_ERR_NOT_INITIALIZED No preceding call to GCInitLib.