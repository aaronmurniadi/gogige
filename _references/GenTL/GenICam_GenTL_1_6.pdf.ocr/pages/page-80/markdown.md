|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

GC_ERR_INVALID_PARAMETER

Parameter piNumDevices is an invalid pointer (NULL or ~0x0).

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.3.6 IFOpenDevice

GC_ERROR IFOpenDevice

( IF_HANDLE hIface,
const char * sDeviceID,
DEVICE_ACCESS_FLAGS iOpenFlag,
DEV_HANDLE * phDevice )

Opens the given sDeviceID with the given iOpenFlag on the given hIface.

Any subsequent call to IFOpenDevice with an sDeviceID which has already been opened will return the error GC_ERR_RESOURCE_IN_USE.

The device ID need not match the one returned from IFGetDeviceID. As long as the GenTL Producer knows how to interpret that ID it will return a valid handle. For example, if in a specific implementation the device has a user-defined name, this function will return a valid handle as long as the provided name refers to an internally known device.

##### Parameters

[in] hIface Interface module to work on.
[in] sDeviceID Unique device ID to open as a 0-terminated C string.
[in] iOpenFlag Configures the open process as defined in the DEVICE_ACCESS_FLAGS.
[out] phDevice Device handle of the newly created Device module. It is recommended to initialize *phDevice to GENTL_INVALID_HANDLE before calling IFOpenDevice to indicate an invalid handle.

##### Returns

GC_ERR_SUCCESS

Operation was successful; no error occurred.

GC_ERR_NOT_INITIALIZED

No preceding call to GCInitLib.

GC_ERR_INVALID_HANDLE

The handle hIface is invalid (NULL) or does not reference an open Interface module retrieved through a call to TLOpenInterface.

GC_ERR_INVALID_ID

The GenTL Producer is unable to interpret the provided ID string sDeviceID or is not able to match it to an existing Device.

GC_ERR_RESOURCE_IN_USE

The Device module has already been instantiated through a previous call to IFOpenDevice.