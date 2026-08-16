|  ![img-122.jpeg](img-122.jpeg) CAM |   | ![img-123.jpeg](img-123.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

##### 6.3.4.6 DevOpenDataStream

|  GC_ERROR | DevOpenDataStream | ( DEV_HANDLE const char * DS_HANDLE * | hDevice, sDataStreamID, phDataStream )  |
| --- | --- | --- | --- |

Opens the given sDataStreamID on the given hDevice.

Any subsequent call to DevOpenDataStream with an sDataStreamID which has already been opened will return the error GC_ERR_RESOURCE_IN_USE.

The Data Stream ID need not match the one returned from DevGetDataStreamID. As long as the GenTL Producer knows how to interpret that ID it will return a valid handle. For example, if in a specific implementation the data stream has a user defined name, this function will return a valid handle as long as the provided name refers to an internally known data stream.

For GenTL Producers which do not provide a data stream the number of available data streams is zero. Calls to DevGetDataStreamID or DevOpenDataStream will fail. Nevertheless a GenTL Producer must export all functions of the public interface.

##### Parameters

|  [in] | hDevice | Device module to work on  |
| --- | --- | --- |
|  [in] | sDataStreamID | Unique data stream ID to open as a 0-terminated C string.  |
|  [out] | phDataStream | Data Stream module handle of the newly created stream. It is recommended to initialize *phDataStream to GENTL_INVALID_HANDLE before calling DevOpenDataStream to indicate an invalid handle.  |

##### Returns

|  GC_ERR_SUCCESS | Operation was successful; no error occurred.  |
| --- | --- |
|  GC_ERR_NOT_INITIALIZED | No preceding call to GCInitLib.  |
|  GC_ERR_INVALID_HANDLE | The handle hDevice is invalid (NULL) or does not reference an open Device module retrieved through a call to IFOpenDevice.  |
|  GC_ERR_RESOURCE_IN_USE | The Data Stream module has already been instantiated through a previous call to DevOpenDataStream.  |
|  GC_ERR_INVALID_ID | The GenTL Producer is unable to interpret the provided ID string sDataStreamID or is not able to match it to an existing Data Stream.  |
|  GC_ERR_INVALID_PARAMETER | Parameters phDataStream and/or sDataStreamID are invalid pointers (NULL or ~0x0).  |