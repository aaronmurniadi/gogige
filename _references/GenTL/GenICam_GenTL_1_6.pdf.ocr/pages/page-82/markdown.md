|  ![img-115.jpeg](img-115.jpeg) CAM |   | ![img-116.jpeg](img-116.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

GC_ERR_INVALID_HANDLE

The handle hIface is invalid (NULL) or does not reference an open Interface module retrieved through a call to TLOpenInterface.

GC_ERR_TIMEOUT

The specified iTimeout expired before the Producer was able to completely update the list. In this case the “old” list stays valid.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.3.8 IFGetParentTL

GC ERROR IFGetParentTL ( IF_HANDLE hIface, TL_HANDLE * phSystem )

Retrieves a handle to the parent TL module.

##### Parameters

[in] hIface Interface module to work on.
[out] phSystem Handle to the parent System module

##### Returns

GC_ERR_SUCCESS Operation was successful; no error occurred.

GC_ERR_NOT_INITIALIZED No preceding call to GCInitLib

GC_ERR_INVALID_HANDLE The handle hIface is invalid (NULL) or does not reference an open Interface module retrieved through a call to TLOpenInterface.

GC_ERR_INVALID_PARAMETER Parameter phSystem is an invalid pointer (NULL or ~0x0).

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

#### 6.3.4 Device Functions

##### 6.3.4.1 DevClose

GC ERROR DevClose (DEV_HANDLE hDevice)

Closes the Device module associated with the given hDevice handle. This frees all resources of the Device module and closes all dependent Data Stream module instances. If DevClose is