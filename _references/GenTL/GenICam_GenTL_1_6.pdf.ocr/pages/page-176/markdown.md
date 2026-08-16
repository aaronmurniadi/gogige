|  ![img-262.jpeg](img-262.jpeg)CAN |   | ![img-263.jpeg](img-263.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

Each iResult member of the many DS_BUFFER_INFO_STACKED structures represents the result of exactly one buffer info query. The applicable error codes of the DSGetBufferInfo function are valid for iResult. These are listed below:

GC_ERR_SUCCESS

Operation was successful; no error occurred.

GC_ERR_NOT_IMPLEMENTED

Specified iInfoCmd is not implemented.

GC_ERR_BUFFER_TOO_SMALL

pBuffer is not NULL and the value of iSize is too small to receive the expected amount of data.

GC_ERR_NOT_AVAILABLE

The request is implemented but the requested information is currently not available for any reason.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.5.1.3 DS_BUFFER_PART_INFO_STACKED

struct DS_BUFFER_PART_INFO_STACKED

Layout of the array elements being used in the function DSGetBufferPartInfoStacked to carry information about multiple buffer part infos as defined in BUFFER_PART_INFO_CMD.

|  Member | Type | Description  |
| --- | --- | --- |
|  iPartIndex [in] | uint32_t | Zero based index of the buffer part to query.  |
|  iInfoCmd [in] | BUFFER_PART_INFO_CMD | The buffer part info to be retrieved.  |
|  iType [out] | INFO_DATATYPE | Data type of the pBuffer content as defined in the BUFFER_PART_INFO_CMD and INFO_DATATYPE.  |
|  iResult [out] | GC_ERROR | The result of the buffer part info query.  |
|  pBuffer [in,out] | void* | Pointer to a user allocated buffer to receive the requested information. If this parameter is NULL, iSize will contain the minimal size of pBuffer in bytes. If the iType is a string the size includes the terminating 0.  |
|  iSize [in,out] | size_t | pBuffer equal NULL: out: minimal size of pBuffer in bytes to hold all information.pBuffer unequal NULL: in: size of the provided pBuffer in bytes. out: number of bytes filled by the function.  |

DSGetBufferPartInfoStacked queries multiple buffer part infos as defined in BUFFER_PART_INFO_CMD at once.

The purpose and “direction” (in/out) of the structure data members is same as corresponding parameters of the DSGetBufferPartInfo function. Similar as other get-info functions,