|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

function must return the error GC_ERR_NO_DATA. The GenTL Consumer would in this case avoid querying information about flows using DSGetFlowInfo.

# Parameters

|  [in] | hDataStream | Data Stream module to work on.  |
| --- | --- | --- |
|  [out] | piNumFlows | Number of flows currently configured for this data stream.  |

# Returns

|  GC_ERR_SUCCESS: | Operation was successful; no error occurred.  |
| --- | --- |
|  GC_ERR_NOT_INITIALIZED: | No preceding call to GCInitLib.  |
|  GC_ERR_INVALID_HANDLE: | The handle hDataStream is invalid (NULL) or does not reference an open Data Stream module retrieved through a call to DevOpenDataStream.  |
|  GC_ERR_NOT_IMPLEMENTED: | The GenTL implementation does not support the flow functionality.  |
|  GC_ERR_INVALID_PARAMETER: | Parameter piNumFlows is an invalid pointer (NULL or ~0x0).  |
|  GC_ERR_NO_DATA: | The GenTL implementation supports flows, but the flow functionality is currently not available for any reason, for example because the given device/stream does not support flows or is in mode not involving flows.  |

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.5.20 DSGetFlowInfo

|  GC_ERROR | DSGetFlowInfo | ( DS_HANDLE | hDataStream,  |
| --- | --- | --- | --- |
|   |  | uint32_t | iFlowIndex,  |
|   |  | FLOW_INFO_CMD | iInfoCmd,  |
|   |  | INFO_DATATYPE * | piType,  |
|   |  | void * | pBuffer,  |
|   |  | size_t * | piSize )  |

Inquires information about individual flows currently configured for this data stream as defined in FLOW_INFO_CMD.

Note that the number of flows and their structure can be dynamic and depend on the configuration of the data stream and of the device output. It is recommended to query it after any configuration possibly affecting the stream and its contents is finished.

# Parameters

[in] hDataStream Data Stream module to work on.