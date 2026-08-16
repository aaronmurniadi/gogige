|  GEN<I>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

#### 6.3.7.5 EventKill

GC ERROR EventKill (EVENT_HANDLE hEvent)

Terminates a waiting operation on a previously registered event object. In case of multiple pending wait operations EventKill causes one wait operation to return with a GC_ERR_ABORT error code. Therefore in order to cancel all pending wait operations EventKill must be called as many times as wait operations are pending. In case this function is called while no wait operation was pending the next call to EventGetData will return a GC_ERR_ABORT. This behavior can be cleared by unregistering and reregistering the event.

In case there are pending events in the queue the EventKill has higher priority and on the pending/next call to EventGetData a GC_ERR_ABORT is returned.

EventKill does not free any resources.

# Parameters

[in] hEvent Handle to event object.

# Returns

GC_ERR_SUCCESS Operation was successful; no error occurred.

GC_ERR_NOT_INITIALIZED No preceding call to GCInitLib.

GC_ERR_INVALID_HANDLE The handle hEvent is invalid (NULL) or does not reference a previously registered event.

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

#### 6.3.7.6 GCRegisterEvent

GC_ERROR GCRegisterEvent ( EVENTSRC_HANDLE hModule,
EVENT_TYPE iEventID,
EVENT_HANDLE * phEvent )

Registers an event object to a certain iEventID. The implementation might change depending on the platform.

Every event registered must be unregistered with the GCUnregisterEvent function.

# Parameters

[in] hModule Module handle to access to register event to.
[in] iEventID Event type to register as defined in EVENT_TYPE.
[out] phEvent New handle to an event object to work with. It is recommended to initialize *phEvent to