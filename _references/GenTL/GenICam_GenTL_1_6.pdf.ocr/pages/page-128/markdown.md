|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

GENTL INVALID HANDLE before calling

GCRegisterEvent to indicate an invalid handle.

Returns

|  GC_ERR_SUCCESS | Operation was successful; no error occurred.  |
| --- | --- |
|  GC_ERR_NOT_INITIALIZED | No preceding call to GCInitLib.  |
|  GC_ERR_INVALID_HANDLE | The handle hModule is invalid (NULL) or does not reference a previously instatiated module.  |
|  GC_ERR_RESOURCE_IN_USE | The given iEventID has been registered before on the given hModule.  |
|  GC_ERR_NOT_IMPLEMENTED | The specified event type is not implemented in the provided module of the GenTL Producer. Applies also if specified iEventID is not a valid event type for given module.  |
|  GC_ERR_NOT_AVAILABLE | The specified event type is not available in the provided module hModule (for example because the remote device does not implement it).  |
|  GC_ERR_INVALID_PARAMETER | Parameter phEvent is an invalid pointer (NULL or ~0x0)  |

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.7.7 GCUnregisterEvent

GC_ERROR GCUnregisterEvent ( EVENTSRC_HANDLE hModule, EVENT_TYPE iEventID )

A call to this function will unregister the given iEventID from the given hModule. This will terminate all pending wait operations of EventGetData with the error code GC_ERR_ABORT. Pending events are silently discarded.

For the EVENT_NEW_BUFFER all pending buffers in the output queue are set to a non-queued state to match the behavior of normal events. All buffers in the input pool or buffers currently being filled are not touched.

Parameters

|  [in] | hModule | Module handle to unregister event with.  |
| --- | --- | --- |
|  [in] | iEventID | Event type to unregister as defined in EVENT_TYPE.  |

Returns

|  GC_ERR_SUCCESS | Operation was successful; no error occurred.  |
| --- | --- |