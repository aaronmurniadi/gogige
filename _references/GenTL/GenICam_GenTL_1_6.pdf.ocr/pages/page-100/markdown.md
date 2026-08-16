|  ![img-139.jpeg](img-139.jpeg) CAM |   | ![img-140.jpeg](img-140.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

GC_ERR_SUCCESS

GC_ERR_NOT_INITIALIZED

GC_ERR_INVALID_HANDLE

GC_ERR_NOT_IMPLEMENTED

GC_ERR_INVALID_PARAMETER

GC_ERR_INVALID_BUFFER

GC_ERR_RESOURCE_IN_USE

GC_ERR_BUFFER_TOO_SMALL

Operation was successful; no error occurred.

No preceding call to GCInitLib.

The handle hDataStream is invalid (NULL) or does not reference an open Data Stream module retrieved through a call to DevOpenDataStream.

One or more flags set in iStartFlags referencing functionality which is not implemented.

iNumToAcquire is 0.

The number of buffers announced through one of the buffer announcement functions is smaller than the number retrieved through a call to DSGetInfo using the STREAM_INFO_BUFF_ANNOUNCE_MIN command.

The Acquisition is already active.

One or more of the announced buffers are smaller than the expected payload size required. This is optional to the GenTL Producer implementation if it chooses to not start acquisition in this case or if the acquisition is started and the buffers are not or only partially filled (see chapter 5.2.1).

Error cases not covered in the list above may return error codes according to chapter 6.1.5 Error Handling on page 61.

##### 6.3.5.12 DSStopAcquisition

GC_ERROR DSStopAcquisition (DS_HANDLE hDataStream, ACQ_STOP_FLAGS iStopFlags)

Stops the acquisition engine on the host. There must be a call to DSStopAcquisition accompanying each call to DSStartAcquisition even though the stream already stopped because the number of frames to acquire was reached. This is also independent of the acquisition modes.

##### Parameters

[in] hDataStream

[in] iStopFlags

Data Stream module to work on.

Stops the acquisition as defined in ACQ_STOP_FLAGS.

##### Returns

GC_ERR_SUCCESS

GC_ERR_NOT_INITIALIZED

Operation was successful; no error occurred.

No preceding call to GCInitLib.