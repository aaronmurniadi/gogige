|  ![img-107.jpeg](img-107.jpeg)CAN |   | ![img-108.jpeg](img-108.jpeg)  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

GC_ERR_INVALID_HANDLE

The handle hIface is invalid (NULL) or does not reference an open Interface module retrieved through a call to TLOpenInterface.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.3.2 IFGetInfo

GC_ERROR IFGetInfo

( IF_HANDLE hIface,
INTERFACE_INFO_CMD iInfoCmd,
INFO_DATATYPE * piType,
void * pBuffer,
size_t * piSize )

Inquires information about the Interface module as defined in INTERFACE_INFO_CMD. The reported information should be in sync to information retrieved through the TLGetInterfaceInfo function.

##### Parameters

[in] hIface

Interface module to work on.

[in] iInfoCmd

Information to be retrieved as defined in INTERFACE_INFO_CMD.

[out] piType

Data type of the pBuffer content as defined in the INTERFACE_INFO_CMD and INFO_DATATYPE.

[in,out] pBuffer

Pointer to a user allocated buffer to receive the requested information. If this parameter is NULL, piSize will contain the minimal size of pBuffer in bytes. If the piType is a string the size includes the terminating 0.

[in,out] piSize

pBuffer equal NULL:
out: minimal size of pBuffer in bytes to hold all information
pBuffer unequal NULL:
in: size of the provided pBuffer in bytes
out: number of bytes filled by the function

##### Returns

GC_ERR_SUCCESS

Operation was successful; no error occurred.

GC_ERR_NOT_INITIALIZED

No preceding call to GCInitLib

GC_ERR_INVALID_HANDLE

The handle hIface is invalid (NULL) or does not reference an open Interface module retrieved through a call to TLOpenInterface.

GC_ERR_NOT_IMPLEMENTED

Specified iInfoCmd is not implemented.