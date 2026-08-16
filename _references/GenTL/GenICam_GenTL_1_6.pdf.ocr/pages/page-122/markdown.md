|  ![img-169.jpeg](img-169.jpeg) CAM |   | ![img-170.jpeg](img-170.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

allows a uint32_t aligned access and the calling GenTL Consumer wants to read 5 bytes starting at address 2. The implementation has to read 8 bytes starting at address 0 and to extract the 5 bytes requested.

In case of an error the function returns the number of successful reads in piNumEntries even though it returns an error code as return value. This is an exception to the statement in the section Error Handling.

#### Parameters

[in] hPort

[in] pEntries

[in,out] piNumEntries

Module or remote device port handle to access the Port from.

Array of structures containing read address and data to read.

In: Number of entries in the array, Out: Number of successful executed reads according to the entries in the pEntries array.

#### Returns

GC_ERR_SUCCESS

Operation was successful; no error occurred.

GC_ERR_NOT_INITIALIZED

No preceding call to GCInitLib.

GC_ERR_INVALID_HANDLE

The handle hPort is invalid (NULL) or does not reference an open module.

GC_ERR_INVALID_PARAMETER

Parameters pEntries and/or piNumEntries are invalid pointers (NULL or ~0x0).

GC_ERR_ACCESS_DENIED

The access to at least one of the requested registers is denied because the register is not readable or because the Port module is opened in a way that it does not allow read access.

GC_ERR_NOT_IMPLEMENTED

The provided module handle does not have a Port module implemented.

GC_ERR_INVALID_ADDRESS

One or more addresses in the entries in pEntries has an invalid address for example because the port's register space is only 32Bit wide and Address is in the 64Bit register space or because there is no register with the specified address.

GC_ERR_IO

Communication error or connection lost.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.