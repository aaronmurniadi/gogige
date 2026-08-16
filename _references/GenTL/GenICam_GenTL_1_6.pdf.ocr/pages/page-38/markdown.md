|  ![img-47.jpeg](img-47.jpeg)CAN |   | ![img-48.jpeg](img-48.jpeg)  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Event Type | Modules | Description  |
| --- | --- | --- |
|   |  | occurred, initiated by the GenTL Producer module the event was registered on. The event ID and the optional data delivered with this event can be put into a GenApi Adapter which then invalidates all related nodes.  |

#### 4.2.2 Event Data Queue

The event data queue is the core of the Signaling. This is a thread safe queue holding event type specific data. Operations on this queue must be locked for example via a mutex in a way that its content may not change when either one of the event functions is accessing it or the module specific thread is accessing it. The GenTL Producer implementation therefore must make sure that access to the queue is as short as possible. Alternatively a lock free queue can be used which supports dequeue operations from multiple threads.

An event object's state is signaled as long as the event data queue is not empty.

Each event data queue must have its own lock if any to secure the state of each instance and to achieve necessary parallelism. Both read and write operations must be locked. The two operations of event data retrieval and the event object signal state handling in the EventGetData function must be atomic. Meaning that, if a lock is used, the lock on the event data queue must be maintained over both operations. Also the operation of putting data in the queue and the event object's state handling must be atomic.

#### 4.2.3 Event Handling

The handling of the event objects is always the same independently of the event type. The signal reason and the signal data of course depend on the event type. The complete state handling is done by the GenTL Producer driver. The GenTL Consumer may call the EventKill function to terminate a single instance of a waiting EventGetData operation. This means that if more than one thread waits for an event, the EventKill function terminates only one wait operation and other threads will continue execution.

##### 4.2.3.1 Registration

Before the GenTL Consumer can be informed about an event, the event object must be registered. After a module instance has been created in the enumeration process an event object can be created with the GCRegisterEvent function. This function returns a unique EVENT_HANDLE which identifies the registered event object. To get information about a registered event the EventGetInfo function can be used. There must be only one event registered per module and event type. If an event object is registered twice on the same module the GCRegisterEvent function must return the error GC_ERR_RESOURCE_IN_USE.

To unregister an event object the GCUnregisterEvent function must be called. If a module is closed all event registrations are automatically unregistered. Events that are still in the queue while an event object is unregistered are silently discarded. Pending wait operations