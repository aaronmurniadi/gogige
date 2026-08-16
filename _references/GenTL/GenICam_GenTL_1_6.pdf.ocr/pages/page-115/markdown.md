|  ![img-156.jpeg](img-156.jpeg) CAM |   | ![img-157.jpeg](img-157.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

[in,out] sURL

Pointer to a user allocated string buffer to receive the list of URLs If this parameter is NULL, piSize will contain the needed size of sURL in bytes. Each entry in the list is 0 terminated. After the last entry there is an additional 0. The size includes the terminating 0 characters.

[in,out] piSize

sURL equal NULL:
out: minimal size of sURL in bytes to hold all information.
sURL unequal NULL:
in: size of the provided sURL in bytes.
out: number of bytes filled by the function.

# Returns

GC_ERR_SUCCESS

Operation was successful; no error occurred.

GC_ERR_NOT_INITIALIZED

No preceding call to GCInitLib.

GC_ERR_INVALID_HANDLE

The handle hPort is invalid (NULL) or does not reference an open module.

GC_ERR_BUFFER_TOO_SMALL

sURL is not NULL and the value of *piSize is too small to receive the expected amount of data.

GC_ERR_INVALID_PARAMETER

Parameter piSize is an invalid pointer (NULL or ~0x0).

GC_ERR_NOT_IMPLEMENTED

The provided module handle does not have a Port module implemented.

GC_ERR_IO

Communication error or connection lost.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.6.3 GCGetNumPortURLs

GC ERROR

GCGetNumPortURLs

( PORT_HANDLE

hPort,

uint32_t *

piNumURLs )

Inquires the number of available URLs for this port.

# Parameters

[in] hPort

Module or remote device port handle to access Port from.

[out] piNumURLs

Number of available URL entries.

# Returns

GC_ERR_SUCCESS

Operation was successful; no error occurred.

GC_ERR_NOT_INITIALIZED

No preceding call to GCInitLib.