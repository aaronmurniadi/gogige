|  ![img-62.jpeg](img-62.jpeg) CAM |   | ![img-63.jpeg](img-63.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

The order in which buffers can be revoked depends on the method in which they were announced. Buffers can be revoked in any order if they were announced by the DSAnnounceBuffer function. More care has to be taken if the DSAllocAndAnnounceBuffer function is used. Normally underlying acquisition engines must not change the base pointer to the memory containing the data within a buffer object. If the DSAllocAndAnnounceBuffer function is used the base pointer of a buffer object may change after another buffer object has been revoked using the DSRevokeBuffer function. Nevertheless, it is recommended to keep the base pointer of a buffer for the lifetime of the buffer handle.

#### 5.2.10 Free Memory

If the GenTL Consumer provides the memory for the buffers using the DSAnnounceBuffer function it also has to free it. Memory allocated by the GenTL Producer implementation using the DSAllocAndAnnounceBuffer function is freed by the library if necessary. The GenTL Consumer must not free this memory.

### 5.3 Buffer Handling Modes

Buffer handling modes describe the internal buffer handling during acquisition. There is only one mandatory mode defined in this document which GenTL Producer implementations should default to. More modes are defined in the GenICam GenTL Standard Features Naming Convention document.

A certain mode may differ from another in these aspects:

- Which buffer is taken from the Input Buffer Pool to be filled
- At which time a filled buffer is moved to the Output Buffer Queue and at which position it is inserted
- Which buffer in the Output Buffer Queue is overwritten (if any at all) on an empty Input Buffer Pool

The graphical description in Figure 5-6 assumes that we are looking at an acquisition start scenario with five announced and queued buffers B0 to B4 ready for acquisition. When a buffer is delivered the image number is stated in the lower bar labeled 'User'. A solid bar on a buffer's time line illustrates its presence in a Buffer pool. A ramp indicates image transfer and therefore transition. During a thin line the Buffer is controlled by the GenTL Consumer and locked for data reception.

#### 5.3.1 Default Mode

The default mode is designed to be simple and flexible with only a few restrictions. This is done to be able to map it to most acquisition techniques used today. If a specific technique cannot be mapped to this mode the goal has to be achieved by copying the data and emulating the behavior in software.

In this scenario every acquired image is delivered to the GenTL Consumer if the mean processing time is below the acquisition time. Peaks in processing time can be mitigated by a larger number of buffers.