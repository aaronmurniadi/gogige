|  ![img-101.jpeg](img-101.jpeg) CAM |   | ![img-102.jpeg](img-102.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

GC_ERR_BUFFER_TOO_SMALL

pBuffer is not NULL and the value of *piSize is too small to receive the expected amount of data.

GC_ERR_INVALID_ID

The GenTL Producer is unable to interpret the provided ID string sIfaceID or is not able to match it to an existing Interface.

GC_ERR_NOT_AVAILABLE

The request is implemented but the requested information is currently not available for any reason.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.2.5 TLGetNumInterfaces

GC_ERROR TLGetNumInterfaces ( TL_HANDLE hSystem, uint32_t * piNumIfaces )

Queries the number of available interfaces on this System module. Prior to this call the TLUpdateInterfaceList function must be called. The list content will not change until the next call of the update function.

This function is not thread safe since it relies on an internal cache.

##### Parameters

[in] hSystem

System module to work on.

[out] piNumIfaces

Number of interfaces on this System module.

##### Returns

GC_ERR_SUCCESS

Operation was successful; no error occurred.

GC_ERR_NOT_INITIALIZED

No preceding call to GCInitLib.

GC_ERR_INVALID_HANDLE

Either the handle hSystem has an invalid value or the handle does not belong to a previously opened TL module through a call to TLOpen.

GC_ERR_INVALID_PARAMETER

Parameter piNumIfaces is an invalid pointer (NULL or ~0x0).

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.