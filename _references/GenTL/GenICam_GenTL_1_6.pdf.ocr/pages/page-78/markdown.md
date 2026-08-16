|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

GC_ERR_INVALID_INDEX

iIndex is greater than the number of available Device modules - 1 retrieved through a call to IFGetNumDevices.

GC_ERR_INVALID_PARAMETER

Parameter piSize is an invalid pointer (NULL or ~0x0).

GC_ERR_BUFFER_TOO_SMALL

sIfaceID is not NULL and the value of *piSize is too small to receive the expected amount of data.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.3.4 IFGetDeviceInfo

GC ERROR IFGetDeviceInfo

( IF_HANDLE hIface,
const char * sDeviceID,
DEVICE_INFO_CMD iInfoCmd,
INFO_DATATYPE * piType,
void * pBuffer,
size_t * piSize )

Inquires information about a device on the given Interface module hIface as defined in DEVICE_INFO_CMD without the need to open the device. The reported information should be in sync to information returned through the DevGetInfo function.

##### Parameters

[in] hIface

Interface module to work on.

[in] sDeviceID

Unique ID of the device to inquire information about. Like with the IFOpenDevice function it is also possible to feed an alternative ID as long as the GenTL Producer knows how to interpret it.

[in] iInfoCmd

Information to be retrieved as defined in DEVICE_INFO_CMD.

[out] piType

Data type of the pBuffer content as defined in the DEVICE_INFO_CMD and INFO_DATATYPE.

[in,out] pBuffer

Pointer to a user allocated buffer to receive the requested information. If this parameter is NULL, piSize will contain the minimal size of pBuffer in bytes. If the piType is a string the size includes the terminating 0.

[in,out] piSize

pBuffer equal NULL:

out: minimal size of pBuffer in bytes to hold all information pBuffer unequal NULL:

in: size of the provided pBuffer in bytes

out: number of bytes filled by the function