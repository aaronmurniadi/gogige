|  ![img-111.jpeg](img-111.jpeg) CAM |   | ![img-112.jpeg](img-112.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

Returns

|  GC_ERR_SUCCESS | Operation was successful; no error occurred.  |
| --- | --- |
|  GC_ERR_NOT_INITIALIZED | No preceding call to GCInitLib.  |
|  GC_ERR_INVALID_HANDLE | The handle hIface is invalid (NULL) or does not reference an open Interface module retrieved through a call to TLOpenInterface.  |
|  GC_ERR_INVALID_ID | The GenTL Producer is unable to interpret the provided ID string sDeviceID or is unable to match it to an existing Device.  |
|  GC_ERR_NOT_IMPLEMENTED | Specified iInfoCmd is not implemented.  |
|  GC_ERR_INVALID_PARAMETER | Parameters piSize and/or piType and/or sDeviceID are invalid pointers (NULL or ~0x0).  |
|  GC_ERR_BUFFER_TOO_SMALL | pBuffer is not NULL and the value of *piSize is too small to receive the expected amount of data.  |
|  GC_ERR_NOT_AVAILABLE | The request is implemented but the requested information is currently not available for any reason.  |

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.3.5 IFGetNumDevices

|  GC_ERROR | IFGetNumDevices | ( IF_HANDLE uint32_t * | hIface, piNumDevices )  |
| --- | --- | --- | --- |

Queries the number of available devices on this Interface module. Prior to this call the IFUpdateDeviceList function must be called. The list content will not change until the next call of the update function.

This function is not thread safe since it relies on an internal cache.

Parameters

|  [in] | hIface | Interface module to work on.  |
| --- | --- | --- |
|  [out] | piNumDevices | Number of devices on this Interface module.  |

Returns

|  GC_ERR_SUCCESS | Operation was successful; no error occurred.  |
| --- | --- |
|  GC_ERR_NOT_INITIALIZED | No preceding call to GCInitLib  |
|  GC_ERR_INVALID_HANDLE | The handle hIface is invalid (NULL) or does not reference an open Interface module retrieved through a call to TLOpenInterface.  |