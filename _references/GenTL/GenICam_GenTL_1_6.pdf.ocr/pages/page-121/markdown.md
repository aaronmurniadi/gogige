|  ![img-167.jpeg](img-167.jpeg) CAM |   | ![img-168.jpeg](img-168.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

Returns

|  GC_ERR_SUCCESS | Operation was successful; no error occurred.  |
| --- | --- |
|  GC_ERR_NOT_INITIALIZED | No preceding call to GCInitLib.  |
|  GC_ERR_INVALID_HANDLE | The handle hPort is invalid (NULL) or does not reference an open module.  |
|  GC_ERR_INVALID_PARAMETER | Parameters pEntries and/or piNumEntries are invalid pointers (NULL or ~0x0).  |
|  GC_ERR_ACCESS_DENIED | The access to at least one of the requested registers is denied because the register is not writable or because the Port module is opened in a way that it does not allow write access.  |
|  GC_ERR_NOT_IMPLEMENTED | The provided module handle does not have a Port module implemented.  |
|  GC_ERR_INVALID_ADDRESS | One or more entries in pEntries has an invalid address for example because the port's register space is only 32Bit wide and Address is in the 64Bit register space or because there is no register with the specified address.  |
|  GC_ERR_INVALID_VALUE | An invalid value has been written. This error code is to be returned if the underlying registermap provides that information. In case the underlying technology/registermap does not provide that level of information a GC_ERR_ACCESS_DENIED is to be returned.  |
|  GC_ERR_IO | Communication error or connection lost.  |

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.6.8 GCReadPortStacked

GC_ERROR GCReadPortStacked ( PORT_HANDLE hPort,
    PORT_REGISTER_STACK_ENTRY *
    pEntries,
    size_t * piNumEntries )

Reads a number of bytes from the given address on the specified hPort for every element in the pEntries array. The endianness of the data content is specified by the GCGetPortInfo function.

If the underlying technology has alignment restrictions on the port access the GenTL Provider implementation has to handle this internally. For example if the underlying technology only