|  ![img-158.jpeg](img-158.jpeg)CAM |   | ![img-159.jpeg](img-159.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

GC_ERR_INVALID_HANDLE

The handle hPort is invalid (NULL) or does not reference an open module.

GC_ERR_INVALID_PARAMETER

Parameter piNumURLs is an invalid pointer (NULL or ~0x0).

GC_ERR_NOT_IMPLEMENTED

The provided module handle does not have a Port module implemented.

GC_ERR_IO

Communication error or connection lost.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.6.4 GCGetPortURLInfo

GC ERROR

GCGetPortURLInfo

( PORT_HANDLE hPort,
    uint32_t iURLIndex,
    URL_INFO_CMD iInfoCmd,
    INFO_DATATYPE * piType,
    void * pBuffer,
    size_t * piSize )

Queries detailed port information as defined in URL_INFO_CMD.

In case a module does not support multiple URLs and/or the related information the function will return GC_ERR_NOT_AVAILABLE for information which cannot be retrieved.

##### Parameters

[in] hPort

Module or remote device port handle to access Port from.

[in] iURLIndex

Zero based index of the URL to query.

[in] iInfoCmd

Information to be retrieved as defined in URL_INFO_CMD.

[out] piType

Data type of the pBuffer content as defined in the URL URL_INFO_CMD and INFO_DATATYPE.

[in,out] pBuffer

Pointer to a user allocated buffer to receive the requested information. If this parameter is NULL, piSize will contain the minimal size of pBuffer in bytes. If the piType is a string the size includes the terminating 0.

[in,out] piSize

pBuffer equal NULL:

out: minimal size of pBuffer in bytes to hold all information.
pBuffer unequal NULL:

in: size of the provided pBuffer in bytes.

out: number of bytes filled by the function.

##### Returns