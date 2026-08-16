|  ![img-163.jpeg](img-163.jpeg) CAM |   | ![img-164.jpeg](img-164.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

##### 6.3.6.6 GCWritePort

GC ERROR GCWritePort

( PORT_HANDLE hPort,
uint64_t iAddress,
const void * pBuffer,
size_t * piSize )

Writes a number of bytes at the given iAddress to the specified hPort. This is the global GenICam GenApi write access function for all ports implemented in the GenTL implementation. The endianness of the data content is specified by the GCGetPortInfo function.

If the underlying technology has alignment restrictions on the port write the GenTL Provider implementation has to handle this internally. For example if the underlying technology only allows a uint32_t aligned access and the calling GenTL Consumer wants to write 5 bytes starting at address 2. The implementation has to read 8 bytes starting at address 0, replace the 5 bytes provided and then write the 8 bytes back (read modify write).

The function is used to handle GenICam GenApi port write access when it is in general unknown which type of data is being written and whether it is acceptable to write only part of the requested piSize bytes. The operation is therefore considered successful only if all requested data has been written. The GenTL Producer is not allowed to report success (GC_ERR_SUCCESS) if the operation was finished only partially.

##### Parameters

[in] hPort

[in] iAddress

[in] pBuffer

[in,out] piSize

Module or remote device port handle to access the Port from.

Byte address to write to.

Pointer to a user allocated byte buffer containing the data to write; this must not be NULL.

Size of the provided pBuffer and thus the amount of bytes to write to the register map; after the write operation this parameter holds the information about the bytes actually written (if that is different from the requested size, the function return value should indicate the reason).

##### Returns

GC_ERR_SUCCESS

GC_ERR_NOT_INITIALIZED

GC_ERR_INVALID_HANDLE

GC_ERR_INVALID_PARAMETER

GC_ERR_ACCESS_DENIED

Operation was successful; no error occurred.

No preceding call to GCInitLib.

The handle hPort is invalid (NULL) or does not reference an open module.

Parameters pBuffer and/or piSize are invalid pointers (NULL or ~0x0).

The access to the requested register iAddress is denied because the register is not writable or because the Port