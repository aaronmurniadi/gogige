|  ![img-165.jpeg](img-165.jpeg) CAM |   | ![img-166.jpeg](img-166.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|   | module is opened in a way that it does not allow write access.  |
| --- | --- |
|  GC_ERR_INVALID_ADDRESS | iAddressis invalid for example because the port's register space is only 32Bit wide and iAddressis in the 64Bit register space or because there is no register with the provided iAddress.  |
|  GC_ERR_NOT_IMPLEMENTED | The provided module handle does not have a Port module implemented.  |
|  GC_ERR_INVALID_VALUE | An invalid value has been written. This error code is to be returned if the underlying registermap provides that information. In case the underlying technology does not provide that level of information a GC_ERR_ACCESS_DENIED is to be returned.  |
|  GC_ERR_IO | Communication error or connection lost.  |

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.6.7 GCWritePortStacked

GC_ERROR GCWritePortStacked ( PORT_HANDLE hPort,
    PORT_REGISTER_STACK_ENTRY *
    pEntries,
    size_t * piNumEntries )

Writes a number of bytes to the given address on the specified hPort for every element in the pEntries array. The endianness of the data content is specified by the GCGetPortInfo function.

If the underlying technology has alignment restrictions on the port write the GenTL Provider implementation has to handle this internally. For example if the underlying technology only allows a uint32_t aligned access and the calling GenTL Consumer wants to write 5 bytes starting at address 2. The implementation has to read 8 bytes starting at address 0, replace the 5 bytes provided and then write the 8 bytes back (read modify write).

In case of an error the function returns the number of successful writes in piNumEntries even though it returns an error code as return value. This is an exception to the statement in the section Error Handling (see chapter 6.1.5).

##### Parameters

[in] hPort
[in] pEntries
[in,out] piNumEntries

Module or remote device port handle to access the Port from. Array of structures containing write address and data to write. In: Number of entries in the array, Out: Number of successful executed writes according to the entries in the pEntries array.