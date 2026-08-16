|  ![img-146.jpeg](img-146.jpeg) CAM |   | ![img-147.jpeg](img-147.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

GC_ERR_INVALID_HANDLE

The handle hDataStream is invalid (NULL) or does not reference an open Data Stream module retrieved through a call to DevOpenDataStream or the handle hBuffer is invalid (NULL) or does not reference an announced Buffer module.

GC_ERR_INVALID_PARAMETER

Parameter pInfoStacked is an invalid pointer (NULL or ~0x0) or iNumInfos is 0.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.5.18 DSGetBufferPartInfoStacked

GC ERROR DSGetBufferPartInfoStacked ( DS_HANDLE hDataStream,
    BUFFER_HANDLE hBuffer,
    DS_BUFFER_PART_INFO_STACKED *
    pInfoStacked,
    size_t iNumInfos )

Inquires various information about individual data parts of the buffer encapsulated in the Buffer module associated with hBuffer on the hDataStream instance as defined in BUFFER_PART_INFO_CMD.

With this function, multiple pieces of information can be queried through a single call to the GenTL Producer without the need to combine that information into a custom structure.

Each buffer part info is grouped in a structure as defined in DS_BUFFER_PART_INFO_STACKED. A pointer to an array of one or more of these structures is used as in and out parameter. Each structure DS_BUFFER_PART_INFO_STACKED of that array passes part index and a BUFFER_PART_INFO_CMD as input and retrieves the required info as output. The details of handling the data members of the DS_BUFFER_PART_INFO_STACKED structure are defined in 6.5.1.3.

Note that the results of the individual queries requested in pInfoStacked do not affect return value of the function, nor the last error information reported by GCGetLastError. Even if certain individual queries fail (for example if given info is not available), the function attempts to process all required infos and reports success, unless the call fails as a whole.

##### Parameters

[in] hDataStream
[in] hBuffer
[in,out] pInfoStacked

Data Stream module to work on.

Buffer handle to retrieve information about.

User allocated array of structures as defined in

DS_BUFFER_PART_INFO_STACKED to receive the requested information. Its length is defined by iNumInfos.

The array contains the various information to be retrieved as defined in BUFFER_PART_INFO_CMD, on output it