|  ![img-55.jpeg](img-55.jpeg)CAN |   | ![img-56.jpeg](img-56.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

buffer is removed from the output buffer queue (delivered), the acquisition engine must not write data into it. Thus this is effectively a buffer locking mechanism.

In order to reuse this buffer a GenTL Consumer has to put the buffer back into the Input Buffer Pool (requeue).

The order of the buffers is defined by the buffer handling mode. Buffers are retrieved by the New Buffer event in a logical first-in-first-out manner. If the acquisition engine does not remove or reorder buffers in the Output Buffer Queue it is always the oldest buffer from the queue that is returned to the GenTL Consumer. Only buffers present in the Announced Buffer Pool which were filled can be in this queue.

### 5.2 Acquisition Chain

Note: this section describes the acquisition chain based on the “traditional” approach using linear contiguous buffers, available since first versions of GenTL specification. This approach is backwards compatible and works well for all payload types. To learn about the more advanced option using structured “composite buffers”, refer to chapter 5.7.

The following description shows the steps to acquire an image from the GenTL Consumer's point of view (default buffer handling mode). Image or data acquisition is performed on the Data Stream module with the functions using the DS_HANDLE. Thus before an acquisition can be carried out, an enumeration of a Data Stream module has to be performed (see chapter 3 Module Enumeration (page 20ff). For a detailed description of the C functions and data types see chapter 6 Software Interface (page 59ff).

Prior to the following steps the remote device and, if necessary (in case a grabber is used), the GenTL Device module should be configured to produce the desired image format. The remote device's PORT_HANDLE can be retrieved from the GenTL Device module's DevGetPort function.