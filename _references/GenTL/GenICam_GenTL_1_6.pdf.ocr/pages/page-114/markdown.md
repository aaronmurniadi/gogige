|  ![img-154.jpeg](img-154.jpeg) CAM |   | ![img-155.jpeg](img-155.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

in: size of the provided pBuffer in bytes.

out: number of bytes filled by the function.

Returns

|  GC_ERR_SUCCESS | Operation was successful; no error occurred.  |
| --- | --- |
|  GC_ERR_NOT_INITIALIZED | No preceding call to GCInitLib.  |
|  GC_ERR_INVALID_HANDLE | The handle hPort is invalid (NULL) or does not reference an open module.  |
|  GC_ERR_NOT_IMPLEMENTED | Specified iInfoCmd is not implemented or the provided module handle does not have a Port module implemented.  |
|  GC_ERR_INVALID_PARAMETER | Parameters piSize and/or piType are invalid pointers (NULL or ~0x0)  |
|  GC_ERR_BUFFER_TOO_SMALL | pBuffer is not NULL and the value of *piSize is too small to receive the expected amount of data.  |
|  GC_ERR_NOT_AVAILABLE | The request is implemented but the requested information is currently not available for any reason.  |
|  GC_ERR_IO | Communication error or connection lost.  |

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.6.2 GCGetPortURL

|  GC_ERROR | GCGetPortURL | ( PORT_HANDLE char * size_t * | hPort, sURL, piSize )  |
| --- | --- | --- | --- |

GCGetPortURL retrieves a URL list with the XML description for the given hPort. See 4.1.2 XML Description page 32 for more information about supported URLs. Each URL is terminated with a trailing '\0' and after the last URL are two '\0'.

In case of multiple XMLs in the device the GCGetNumPortURLs and GCGetPortURLInfo should be used.

This function has been deprecated. Producers should support the new functions GCGetNumPortURLs and GCGetPortURLInfo. In this case this function may only return a subset of the available URLs in the string list. It is up to the implementor which URL to return.

##### Parameters

[in] hPort

Module or remote device port handle to access Port from.