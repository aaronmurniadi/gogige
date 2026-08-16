|  ![img-171.jpeg](img-171.jpeg) CAM |   | ![img-172.jpeg](img-172.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  [in,out] piSize | the queue without being delivered. In case of a New Buffer Event being retrieved with pBuffer equals NULL the buffer is removed from the output queue and not requeued. Size of the provided pBuffer in bytes; after the write operation this parameter holds the information about the bytes actually written.  |
| --- | --- |
|  [in] iTimeout | Timeout for the wait in ms. If set to GENTL INFINITE the timeout is infinite and the function will only return if the operation is completed or if EventKill is called on this event object.A value of 0 checks the state of the event object and returns immediately either with a timeout or with event data.  |

# Returns

|  GC_ERR_SUCCESS | Operation was successful; no error occurred.  |
| --- | --- |
|  GC_ERR_NOT_INITIALIZED | No preceding call to GCInitLib.  |
|  GC_ERR_INVALID_HANDLE | The handle hEvent is invalid (NULL) or does not reference a previously registered event.  |
|  GC_ERR_INVALID_PARAMETER | Parameter piSize is an invalid pointer (NULL or ~0x0).  |
|  GC_ERR_BUFFER_TOO_SMALL | pBuffer is not NULL and the value of *piSize is too small to receive the expected amount of data.  |
|  GC_ERR_ABORT | The current wait operation has been terminated through a call to EventKill.  |
|  GC_ERR_TIMEOUT | The specified iTimeout expired before the event hEvent occurred.  |

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.7.3 EventGetDataInfo

|  GC_ERROR | EventGetDataInfo | ( EVENT_HANDLE hEvent, const void * pInBuffer, size_t iInSize, EVENT_DATA_INFO_CMD iInfoCmd, INFO_DATATYPE * piType, void * pOutBuffer, size_t * piOutSize )  |
| --- | --- | --- |

Parses the transport layer technology dependent event info.

# Parameters