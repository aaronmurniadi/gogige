|  GEN<ICAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

a call to DevOpenDataStream or the handle hBuffer is invalid (NULL) or does not reference an announced Buffer.

GC_ERR_NOT_IMPLEMENTED: The GenTL implementation does not support querying information about buffer parts.

GC_ERR_INVALID_PARAMETER: Parameter piNumParts is an invalid pointer (NULL or ~0x0).

GC_ERR_NO_DATA: The GenTL implementation supports querying information about buffer parts, but the information about number of data parts in the buffer is currently not available for any reason, for example because the buffer does not contain multi-part payload.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.5.16 DSGetBufferPartInfo

GC_ERROR DSGetBufferPartInfo ( DS_HANDLE hDataStream,
    BUFFER_HANDLE hBuffer,
    uint32_t iPartIndex,
    BUFFER_PART_INFO_CMD iInfoCmd,
    INFO_DATATYPE * piType,
    void * pBuffer,
    size_t * piSize )

Inquires information about individual data parts of the buffer encapsulated in the Buffer module associated with hBuffer on the hDataStream instance as defined in BUFFER_PART_INFO_CMD.

To retrieve multiple infos about one or more buffer parts at once and reduce the number of calls from the GenTL Consumer to the GenTL Producer, DSGetBufferPartInfoStacked function can be used instead.

##### Parameters

|  [in] | hDataStream | Data Stream module to work on.  |
| --- | --- | --- |
|  [in] | hBuffer | Buffer handle to retrieve information about.  |
|  [in] | iPartIndex | Zero based index of the buffer part to query.  |
|  [in] | iInfoCmd | Information to be retrieved as defined inBUFFER_PART_INFO_CMD.  |
|  [out] | piType | Data type of the pBuffer content as defined in theBUFFER_PART_INFO_CMD and INFO_DATATYPE.  |
|  [in,out] | pBuffer | Pointer to a user allocated buffer to receive the requested information. If this parameter is NULL, piSize will contain the  |