|  GEN<I>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

##### 6.3.5.17 DSGetBufferInfoStacked

|  GC_ERROR | DSGetBufferInfoStacked ( DS_HANDLE hDataStream, BUFFER_HANDLE hBuffer, DS_BUFFER_INFO_STACKED * pInfoStacked, size_t iNumInfos )  |
| --- | --- |

Inquire various information about the Buffer module associated with hBuffer on the hDataStream instance as defined in BUFFER_INFO_CMD.

With this function, multiple pieces of information can be queried through a single call to the GenTL Producer without the need to combine that information into a custom structure.

Each buffer info is grouped in a structure as defined in DS_BUFFER_INFO_STACKED. A pointer to an array of one or more of these structures is used as in and out parameter. Each structure DS_BUFFER_INFO_STACKED of that array passes a BUFFER_INFO_CMD as input and retrieves the required info as output. The details of handling the data members of the DS_BUFFER_INFO_STACKED structure are defined in 6.5.1.2.

Note that the results of the individual queries requested in pInfoStacked do not affect return value of the function, nor the last error information reported by GCGetLastError. Even if certain individual queries fail (for example if given info is not available), the function attempts to process all required infos and reports success, unless the call fails as a whole.

##### Parameters

|  [in] | hDataStream | Data Stream module to work on.  |
| --- | --- | --- |
|  [in] | hBuffer | Buffer handle to retrieve information about.  |
|  [in,out] | pInfoStacked | User allocated array of structures as defined in DS_BUFFER_INFO_STACKED to receive the requested information. Its length is defined by iNumInfos.The array contains the various information to be retrieved as defined in BUFFER_INFO_CMD, on output it provides the results. The details of the structure members and their use for the info exchange are defined in DS_BUFFER_INFO_STACKED (6.5.1.2).  |
|  [in] | iNumInfos | Number of stacked buffer infos to retrieve.  |

##### Returns

GC_ERR_SUCCESS

Operation was successful; no error occurred. This does not indicate whether every buffer info was retrieved successfully. Therefore each DS_BUFFER_INFO_STACKED::iResult needs to be checked individually as described in DS_BUFFER_INFO_STACKED.

GC_ERR_NOT_INITIALIZED

No preceding call to GCInitLib.