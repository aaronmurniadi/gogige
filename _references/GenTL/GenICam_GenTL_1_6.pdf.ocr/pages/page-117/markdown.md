|  ![img-160.jpeg](img-160.jpeg) CAM |   | ![img-161.jpeg](img-161.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  GC_ERR_SUCCESS | Operation was successful; no error occurred.  |
| --- | --- |
|  GC_ERR_NOT_INITIALIZED | No preceding call to GCInitLib.  |
|  GC_ERR_INVALID_HANDLE | The handle hPort is invalid (NULL) or does not reference an open module.  |
|  GC_ERR_INVALID_INDEX | iURLIndex is greater than the number available URLs -1.  |
|  GC_ERR_NOT_IMPLEMENTED | Specified iInfoCmd is not implemented or the provided module handle does not have a Port module implemented.  |
|  GC_ERR_NOT_AVAILABLE | The module does not provide the requested information.  |
|  GC_ERR_INVALID_PARAMETER | Parameters piSize and/or piType are invalid pointers (NULL or ~0x0).  |
|  GC_ERR_BUFFER_TOO_SMALL | pBuffer is not NULL and the value of *piSize is too small to receive the expected amount of data.  |
|  GC_ERR_IO | Communication error or connection lost.  |

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.6.5 GCReadPort

GC_ERROR GCReadPort

( PORT_HANDLE

hPort,

uint64_t

iAddress,

void *

pBuffer,

size_t *

piSize )

Reads a number of bytes from a given iAddress from the specified hPort. This is the global GenICam GenApi read access function for all ports implemented in the GenTL implementation. The endianness of the data content is specified by the GCGetPortInfo function.

If the underlying technology has alignment restrictions on the port read, the GenTL Provider implementation has to handle this internally. For example if the underlying technology only allows a 4-byte aligned access and the calling GenTLConsumer wants to read 5 bytes starting at address 2. The implementation has to read 8 bytes starting at address 0 and then it must only return the requested 5 bytes.

The function is used to handle GenICam GenApi port read access when it is in general unknown which type of data is being read and whether it is acceptable to read only part of the requested piSize bytes. The operation is therefore considered successful only if all requested