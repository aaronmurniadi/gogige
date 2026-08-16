|  GEN<ICAM |   | ![img-145.jpeg](img-145.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

[in,out] piSize

minimal size of pBuffer in bytes. If the piType is a string the size includes the terminating 0.

pBuffer equal NULL :

out: minimal size of pBuffer in bytes to hold all information.

pBuffer unequal NULL :

in: size of the provided pBuffer in bytes.

out: number of bytes filled by the function.

#### Returns

GC_ERR_SUCCESS:

Operation was successful; no error occurred.

GC_ERR_NOT_INITIALIZED:

No preceding call to GCInitLib.

GC_ERR_INVALID_HANDLE:

The handle hDataStream is invalid (NULL) or does not reference an open Data Stream module retrieved through a call to DevOpenDataStream or the handle hBuffer is invalid (NULL) or does not reference an announced Buffer.

GC_ERR_NOT_IMPLEMENTED:

Specified iInfoCmd is not implemented or the GenTL implementation does not support querying information about buffer parts.

GC_ERR_INVALID_PARAMETER

Parameters piSize and/or piType are invalid pointers (NULL or ~0x0).

GC_ERR_INVALID_INDEX:

iPartIndex is greater than the number of available buffer parts - 1 retrieved through a call to DSGetNumBufferParts.

GC_ERR_NO_DATA:

The buffer referenced by hBuffer does not contain data parts.

GC_ERR_BUFFER_TOO_SMALL:

pBuffer is not NULL and the value of *piSize is too small to receive the expected amount of data.

GC_ERR_NOT_AVAILABLE:

The request is implemented but the requested information is currently not available for any reason.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.