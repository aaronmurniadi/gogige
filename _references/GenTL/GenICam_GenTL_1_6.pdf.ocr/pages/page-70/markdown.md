|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

##### 6.3.2.3 TLGetInterfaceID

GC_ERROR TLGetInterfaceID ( TL_HANDLE hSystem,
    uint32_t iIndex,
    char * sIfaceID,
    size_t * piSize )

Queries the unique ID of the interface at iIndex in the internal interface list. Prior to this call the TLUpdateInterfaceList function must be called. The list content will not change until the next call of the update function.

This function is not thread safe since it relies on an internal cache.

##### Parameters

[in] hSystem System module to work on.
[in] iIndex Zero-based index of the interface on this system.
[in,out] sIfaceID Pointer to a user allocated C string buffer to receive the Interface module ID at the given iIndex. If this parameter is NULL, piSize will contain the needed size of sIfaceID in bytes. The size includes the terminating 0.
[in,out] piSize sIfaceID equal NULL:
out: minimal size of sIfaceID in bytes to hold all information sIfaceID unequal NULL:
in: size of the provided sIfaceID in bytes
out: number of bytes filled by the function

##### Returns

GC_ERR_SUCCESS Operation was successful; no error occurred.
GC_ERR_NOT_INITIALIZED No preceding call to GCInitLib.
GC_ERR_INVALID_HANDLE The handle hSystem is invalid (NULL) or does not reference an open System module retrieved through a call to TLOpen.
GC_ERR_INVALID_INDEX iIndex is greater than the number of available Interface modules - 1 retrieved through a call to TLGetNumInterfaces.
GC_ERR_INVALID_PARAMETER Parameter piSize is an invalid pointer (NULL or ~0x0).
GC_ERR_BUFFER_TOO_SMALL sIfaceID is not NULL and the value of *piSize is too small to receive the expected amount of data.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.