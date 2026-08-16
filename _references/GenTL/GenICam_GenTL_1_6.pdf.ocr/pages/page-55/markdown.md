|  ![img-76.jpeg](img-76.jpeg)CAN |   | ![img-77.jpeg](img-77.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

### 5.7 Structured Data Acquisition

The universal acquisition into contiguous buffers (announced using DSAllocAndAnnounceBuffer or DSAnnounceBuffer) described in previous sections is sufficient in many use cases and ensures backward compatibility with older implementations.

However, GenTL specification version 1.6 introduces additional options allowing more flexible control of the acquisition engine: composite buffers and data stream flows.

#### 5.7.1 Composite (Non-contiguous) Buffers

In some use cases, the GenTL Consumer might prefer that specific parts of the data acquired in each acquisition step (each “New Buffer” event or “block” as referred in some transport layer technologies) are stored in discrete memory locations.

To achieve that, the consumer announces each buffer as a set of memory segments, using DSAnnounceCompositeBuffer which returns a unique BUFFER_HANDLE similar to the other buffer announcement functions. From buffer flow and lifetime perspective the composite buffers are treated by the acquisition engine exactly same as the “traditional” contiguous buffers, they are queued, info-queried, signaled (through “New Buffer” events) or revoked same as other buffers (as described in 5.2), the only difference is that they have an internal structure, consisting of discrete memory areas. To use them effectively, the GenTL Consumer and the GenTL Producer’s acquisition engine need to act and interpret the composite buffers’ structure in concert.

It is important to realize that announcing a composite buffer by itself is not sufficient to get the data properly split in its buffer segments during acquisition. The GenTL Consumer is responsible to care that the GenTL Producer and/or device are capable and configured for such operation.

The configuration of how the acquisition data are mapped into buffer segments is in general beyond scope of the GenTL specification and might be controlled for example by means of stream and data format related features in SFNC or GenTL SFNC. The composite buffers are primarily intended as targets for structured data acquired by data stream flows mechanism (5.7.2) and hence for acquisition in the GenDC format (5.7.3). In these cases, the mapping is implied by means of the flows and GenDC configuration.

#### 5.7.2 Data Stream Flows

The acquisition data stream transported into the target buffers by means of a Data Stream module might be split into multiple “flows”. The flows are independent channels within the data stream, responsible to transfer individual components of the acquisition data.

The flow mechanism serves two main purposes:

- Allow transferring independent data parts into a single buffer in parallel as they become available, without need to buffer them within the source (remote device), which might be required for sequential transfer.