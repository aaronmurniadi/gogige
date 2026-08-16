|  ![img-117.jpeg](img-117.jpeg) CAM |   | ![img-118.jpeg](img-118.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

# Returns

|  GC_ERR_SUCCESS | Operation was successful; no error occurred.  |
| --- | --- |
|  GC_ERR_NOT_INITIALIZED | No preceding call to GCInitLib.  |
|  GC_ERR_INVALID_HANDLE | The handle hDevice is invalid (NULL) or does not reference an open Device module retrieved through a call to IFOpenDevice.  |
|  GC_ERR_NOT_IMPLEMENTED | Specified iInfoCmd is not implemented.  |
|  GC_ERR_INVALID_PARAMETER | Parameters piSize and/or piType are invalid pointers (NULL or ~0x0).  |
|  GC_ERR_BUFFER_TOO_SMALL | pBuffer is not NULL and the value of *piSize is too small to receive the expected amount of data.  |
|  GC_ERR_NOT_AVAILABLE | The request is implemented but the requested information is currently not available for any reason.  |

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.4.3 DevGetDataStreamID

|  GC_ERROR | DevGetDataStreamID | ( DEV_HANDLE | hDevice,  |
| --- | --- | --- | --- |
|   |  | uint32_t | iIndex,  |
|   |  | char * | sDataStreamID,  |
|   |  | size_t * | piSize )  |

Queries the unique ID of the data stream at iIndex in the internal data stream list.

For GenTL Producers which do not provide a data stream the number of available data streams is zero. Calls to DevGetDataStreamID or DevOpenDataStream will fail. Nevertheless a GenTL Producer must export all functions of the public interface.

# Parameters

|  [in] | hDevice | Device module to work on.  |
| --- | --- | --- |
|  [in] | iIndex | Zero-based index of the data stream on this device.  |
|  [in,out] | sDataStreamID | Pointer to a user allocated C string buffer to receive the Datastream module ID at the given iIndex. If this parameter is NULL, piSize will contain the needed size of sDataStreamID in bytes. The size includes the terminating 0.  |
|  [in,out] | piSize | sDataStreamID equal NULL:out: minimal size of sDataStreamID in bytes to hold all information sDataStreamID unequal NULL:  |