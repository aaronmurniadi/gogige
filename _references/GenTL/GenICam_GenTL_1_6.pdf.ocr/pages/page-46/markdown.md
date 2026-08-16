|  GENICAM |   | ![img-61.jpeg](img-61.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

buffers are announced prior to the call to DSStartAcquisition it is also possible to announce or revoke buffers in between calls to DSStartAcquisition and DSStopAcquisition while the acquisition is ongoing in case the underlying GenTL Producer supports this. In order to revoke a buffer it is additionally necessary that the particular buffer is only referenced in the announced buffer pool which means that it is neither in any of the acquisition queues and that it is currently not acquired to. In case the underlying GenTL Producer does not support the announcing or revoking buffers while the acquisition is active (in between calls to DSStartAcquisition and DSStopAcquisition) it is also valid for the GenTL Producer to return a GC_ERR_BUSY from a call to DSAnnounceBuffer, DSAllocAndAnnounceBuffer or DSRevokeBuffer.

Along with the buffer memory a pointer to user data is passed which may point to a buffer specific implementation. That pointer is delivered along with the Buffer module handle in the New Buffer event.

The DSAnnounceBuffer and DSAllocAndAnnounceBuffer functions return a unique BUFFER_HANDLE to identify the buffer in the process. The minimum number of buffers that must be announced depends on the technology used. This information can be queried from the Data Stream module features. If there is a known maximum this can also be queried. Otherwise the number of buffers is only limited by available memory.

The acquisition engine normally stores additional data with the announced buffers to be able to, e.g., use DMA transfer to fill the buffers.

#### 5.2.3 Queue Buffers

To acquire data at least one buffer has to be queued with the DSQueueBuffer function. When a buffer is queued it is put into the Input Buffer Pool. The user has to explicitly call DSQueueBuffer to place the buffers into the Input Buffer Pool. The order in which the buffers are queued does not need to match the order in which they were announced. The queue order also does not necessarily have an influence in which order the buffers are filled. This depends only on the buffer handling mode.

#### 5.2.4 Register New Buffer Event

An event object to the data stream must be registered using the NewBufferEvent ID in order to be notified on newly filled buffers. The GCRegisterEvent function returns a unique EVENT_HANDLE which can be used to obtain event specific data when the event was signaled. For the "New Buffer" event this data is the BUFFER_HANDLE and the user data pointer.

#### 5.2.5 Start Acquisition

First the acquisition engine on the host is started with the DSStartAcquisition function. After that the acquisition on the remote device is to be started by setting the "AcquisitionStart" standard feature via the GenICam GenApi.

If a device implements the SFNC Transfer Control features, the GenTL Consumer may need to start the transfer on the remote device as well, depending on the operating mode.