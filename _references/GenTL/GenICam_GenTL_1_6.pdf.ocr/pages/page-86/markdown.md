|  ![img-120.jpeg](img-120.jpeg) CAM |   | ![img-121.jpeg](img-121.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

GC_ERR_INVALID_HANDLE

The handle hDevice is invalid (NULL) or does not reference an open Device module retrieved through a call to IFOpenDevice.

GC_ERR_INVALID_PARAMETER

Parameter piNumDataStreams is an invalid pointer (NULL or ~0x0).

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.4.5 DevGetPort

GC ERROR DevGetPort

( DEV_HANDLE hDevice,
PORT_HANDLE * phRemoteDev )

Retrieves the port handle for the associated remote device.

This function does not return the handle for the Port functions for the Device module but for the physical remote device.

The phRemoteDev handle must not be closed explicitly. This is done automatically when DevClose is called on this Device module.

The remote device Port handle is no valid source for Events. Therefore it must not be used to register Events through GCRegisterEvent.

##### Parameters

[in] hDevice
[out] phRemoteDev

Device module to work on.
Port handle for the remote device. It is recommended to initialize *phRemoteDev to GENTL_INVALID_HANDLE before calling DevGetPort to indicate an invalid handle.

##### Returns

GC_ERR_SUCCESS

Operation was successful; no error occurred.

GC_ERR_NOT_INITIALIZED

No preceding call to GCInitLib.

GC_ERR_INVALID_HANDLE

The handle hDevice is invalid (NULL) or does not reference an open Device module retrieved through a call to IFOpenDevice.

GC_ERR_INVALID_PARAMETER

Parameter phRemoteDev is an invalid pointer (NULL or ~0x0).

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.