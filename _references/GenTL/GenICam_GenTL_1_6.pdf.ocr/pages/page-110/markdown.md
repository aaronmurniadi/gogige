|  GEN<ICAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

[in] iFlowIndex

[in] iInfoCmd

[out] piType

[in,out] pBuffer

[in,out] piSize

Zero based index of the flow to query.

Information to be retrieved as defined in FLOW_INFO_CMD.

Data type of the pBuffer content as defined in the FLOW_INFO_CMD and INFO_DATATYPE.

Pointer to a user allocated buffer to receive the requested information. If this parameter is NULL, piSize will contain the minimal size of pBuffer in bytes. If the piType is a string the size includes the terminating 0.

pBuffer equal NULL:

out: minimal size of pBuffer in bytes to hold all information.

pBuffer unequal NULL:

in: size of the provided pBuffer in bytes.

out: number of bytes filled by the function.

### Returns

GC_ERR_SUCCESS:

GC_ERR_NOT_INITIALIZED:

GC_ERR_INVALID_HANDLE:

GC_ERR_NOT_IMPLEMENTED:

GC_ERR_INVALID_PARAMETER:

GC_ERR_INVALID_INDEX:

GC_ERR_NO_DATA:

GC_ERR_BUFFER_TOO_SMALL_:

GC_ERR_NOT_AVAILABLE:

Operation was successful; no error occurred.

No preceding call to GCInitLib.

The handle hDataStream is invalid (NULL) or does not reference an open Data Stream module retrieved through a call to DevOpenDataStream.

Specified iInfoCmd is not implemented or the GenTL implementation does not support querying information about flows.

Parameters piSize and/or piType are invalid pointers (NULL or ~0x0).

iFlowIndex is greater than the number of available flows - 1 retrieved through a call to DSGetNumFlows.

The data stream referenced by hDataStream has currently no configured flows. This applies also if the actual device does not support flows or is in mode not involving flows

pBuffer is not NULL and the value of *piSize is too small to receive the expected amount of data.

The request is implemented but the requested information is currently not available for any reason.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.