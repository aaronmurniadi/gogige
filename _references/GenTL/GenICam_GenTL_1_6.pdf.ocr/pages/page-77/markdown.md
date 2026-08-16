|  ![img-109.jpeg](img-109.jpeg) CAM |   | ![img-110.jpeg](img-110.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

GC_ERR_INVALID_PARAMETER

Parameters piSize and/or piType are invalid pointers (NULL or ~0x0)

GC_ERR_BUFFER_TOO_SMALL

pBuffer is not NULL and the value of *piSize is too small to receive the expected amount of data.

GC_ERR_NOT_AVAILABLE

The request is implemented but the requested information is currently not available for any reason.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.3.3 IFGetDeviceID

GC_ERROR IFGetDeviceID

( IF_HANDLE hIface,
    uint32_t iIndex,
    char * sDeviceID,
    size_t * piSize )

Queries the unique ID of the device at iIndex in the internal device list. Prior to this call the IFUpdateDeviceList function must be called. The list content will not change until the next call of the update function.

This function is not thread safe since it relies on an internal cache.

##### Parameters

[in] hIface

Interface module to work on.

[in] iIndex

Zero-based index of the device on this interface.

[in,out] sDeviceID

Pointer to a user allocated C string buffer to receive the Device module ID at the given iIndex. If this parameter is NULL, piSize will contain the needed size of sDeviceID in bytes. The size includes the terminating 0.

[in,out] piSize

sDeviceID equal NULL:

out: minimal size of sDeviceID in bytes to hold all information

sDeviceID unequal NULL:

in: size of the provided sDeviceID in bytes

out: number of bytes filled by the function

##### Returns

GC_ERR_SUCCESS

Operation was successful; no error occurred.

GC_ERR_NOT_INITIALIZED

No preceding call to GCInitLib.

GC_ERR_INVALID_HANDLE

The handle hIface is invalid (NULL) or does not reference an open Interface module retrieved through a call to TLOpenInterface.