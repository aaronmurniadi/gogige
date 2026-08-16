|  ![img-113.jpeg](img-113.jpeg) CAM |   | ![img-114.jpeg](img-114.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

GC_ERR_INVALID_PARAMETER

Parameters sDeviceID and/or phDevice are invalid pointers (NULL or ~0x0) or iOpenFlag contains a non valid/unknown value.

GC_ERR_NOT_IMPLEMENTED

iOpenFlag contains a value, which is not implemented by this GenTL Producer.

GC_ERR_ACCESS_DENIED

The access to the requested device is denied.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.3.7 IFUpdateDeviceList

GC_ERROR IFUpdateDeviceList ( IF_HANDLE hIface,
bool8_t * pbChanged,
uint64_t iTimeout )

Updates the internal list of available devices. This may change the connection between a list index and a device ID. It is recommended to call IFUpdateDeviceList regularly from time to time and after reconfiguration of the Interface module to reflect possible changes.

A call to this function has implications on the thread safety of

- IFGetNumDevices
- IFGetDeviceID

##### Parameters

[in] hIface

[out] pbChanged

[in] iTimeout

Interface module to work on.

Contains true if the internal list was changed and false otherwise. If set to NULL nothing is written to this parameter.

Timeout in ms. If set to GENTL_INFINITE the timeout is infinite and the function will only return if the operation is completed. In any case the GenTL Producer must make sure that this operation is completed in a reasonable amount of time depending on the underlying technology. Please be aware that there is no defined way of terminating such an update operation. On the other hand it is the GenTL Consumer's responsibility to call this function with a reasonable timeout.

##### Returns

GC_ERR_SUCCESS

GC_ERR_NOT_INITIALIZED

Operation was successful; no error occurred.

No preceding call to GCInitLib.