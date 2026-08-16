|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

#### 6.3.7 Signaling Functions

##### 6.3.7.1 EventFlush

GC_ERROR EventFlush (EVENT_HANDLE hEvent)

Flushes all events in the given hEvent object. This call empties the event data queue.

##### Parameters

[in] hEvent Event handle to flush queue on.

##### Returns

GC_ERR_SUCCESS Operation was successful; no error occurred.

GC_ERR_NOT_INITIALIZED No preceding call to GCInitLib.

GC_ERR_INVALID_HANDLE The handle hEvent is invalid (NULL) or does not reference a previously registered event.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.7.2 EventGetData

GC_ERROR EventGetData (EVENT_HANDLE hEvent,
void * pBuffer,
size_t * piSize,
uint64_t iTimeout)

Retrieves the next event data entry from the event data queue associated with the hEvent.

The data content can be queried by the EventGetDataInfo function.

The default buffer size which can hold all the event data can be queried with the EventGetInfo function. This needs to be queried only once. The default size must not change during runtime.

In case of a New Buffer event the EventGetData function return the EVENT_NEW_BUFFER_DATA structure in the provided buffer.

In case EventGetData returns an error (for example GC_ERR_ABORT) no event is removed from the internal queue and the event stays signaled. Event counters are not affected.

##### Parameters

[in] hEvent Event handle to wait for.
[out] pBuffer Pointer to a user allocated buffer to receive the event data. The data type of the buffer is dependent on the event ID of the hEvent. If this value is NULL the data is removed from