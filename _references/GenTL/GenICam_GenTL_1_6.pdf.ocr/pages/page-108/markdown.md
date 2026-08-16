|  ![img-148.jpeg](img-148.jpeg) CAM |   | ![img-149.jpeg](img-149.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

provides the results. The details of the structure members and their use for the info exchange are defined in

DS BUFFER PART INFO STACKED (6.5.1.3).

[in] iNumInfos

Number of stacked buffer part infos to retrieve.

# Returns

GC_ERR_SUCCESS:

Operation was successful; no error occurred. This does not indicate whether every buffer part info was retrieved successfully. Therefore each DS_BUFFER_PART_INFO_STACKED::iResult needs to be checked individually as described in DS_BUFFER_PART_INFO_STACKED.

GC_ERR_NOT_INITIALIZED:

No preceding call to GCInitLib.

GC_ERR_INVALID_HANDLE:

The handle hDataStream is invalid (NULL) or does not reference an open Data Stream module retrieved through a call to DevOpenDataStream or the handle hBuffer is invalid (NULL) or does not reference an announced Buffer module.

GC_ERR_NOT_IMPLEMENTED:

The GenTL implementation does not support querying information about buffer parts.

GC_ERR_INVALID_PARAMETER

Parameter pInfoStacked is an invalid pointer (NULL or ~0x0) or iNumInfos is 0.

GC_ERR_NO_DATA:

The buffer referenced by hBuffer does not contain data parts.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.5.19 DSGetNumFlows

GC ERROR

DSGetNumFlows

( DS HANDLE

hDataStream,

uint32_t *

piNumFlows )

Inquires the number of flows currently configured for this data stream.

Note that the number of flows and their structure can be dynamic and depend on the configuration of the data stream and of the device output. It is recommended to query it after any configuration possibly affecting the stream and its contents is finished.

Note also that in case of GenDC streaming the number of flows corresponds to the size of the GenDC flow mapping table.

If querying flow-related information is not relevant in the moment, for example because the device itself does not support flows (or does not expose them in current working mode), the