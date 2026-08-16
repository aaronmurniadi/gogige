|  GEN<I>CAM |   | ![img-162.jpeg](img-162.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

data has been read. The GenTL Producer is not allowed to report success (GC_ERR_SUCCESS) if the operation was finished only partially.

Parameters

|  [in] | hPort | Module or remote device port handle to access Port from.  |
| --- | --- | --- |
|  [in] | iAddress | Byte address to read from.  |
|  [out] | pBuffer | Pointer to a user allocated byte buffer to receive data; this must not be NULL.  |
|  [in,out] | piSize | Size of the provided pBuffer and thus the amount of bytes to read from the register map; after the read operation this parameter holds the information about the bytes actually read (if that is different from the requested size, the function return value should indicate the reason).  |

Returns

|  GC_ERR_SUCCESS | Operation was successful; no error occurred.  |
| --- | --- |
|  GC_ERR_NOT_INITIALIZED | No preceding call to GCInitLib.  |
|  GC_ERR_INVALID_HANDLE | The handle hPort is invalid (NULL) or does not reference an open module.  |
|  GC_ERR_INVALID_PARAMETER | Parameters pBuffer and/or piSize are invalid pointers (NULL or ~0x0).  |
|  GC_ERR_ACCESS_DENIED | The access to the requested register iAddress is denied because the register is not readable.  |
|  GC_ERR_INVALID_ADDRESS | iAddress is invalid for example because the port's register space is only 32Bit wide and iAddress is in the 64Bit register space or because there is no register with the provided iAddress.  |
|  GC_ERR_NOT_IMPLEMENTED | The provided module handle does not have a Port module implemented.  |
|  GC_ERR_IO | Communication error or connection lost.  |

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.