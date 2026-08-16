|  ![img-13.jpeg](img-13.jpeg)CAN |   | ![img-14.jpeg](img-14.jpeg)emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

Interface module also presents Signaling and module configuration capabilities to the GenTL Consumer.

One system may contain zero, one or multiple interfaces. An interface represents only one transport layer technology. It is not allowed to have, e.g., a GigE Vision camera and a Camera Link camera on one interface. There is no logical limitation on the number of interfaces addressed by the system. This is limited solely by the hardware used.

#### 2.2.3 Device Module

The Device module represents the GenTL Producers' proxy for one physical remote device. The responsibility of the Device module is to enable the communication with the remote device and to enumerate and instantiate Data Stream modules. The Device module also presents Signaling and module configuration capabilities to the GenTL Consumer.

One Interface module can contain zero, one or multiple Device module instances. A device is always of one transport layer technology. There is no logical limitation on the number of devices attached to an interface. This is limited solely by the hardware used.

#### 2.2.4 Data Stream Module

A single (image) data stream from a remote device is represented by the Data Stream module. The purpose of this module is to provide the acquisition engine and to maintain the internal buffer pool. Beside that the Data Stream module also presents Signaling and module configuration capabilities to the GenTL Consumer.

One device can contain zero, one or multiple data streams. There is no logical limitation on the number of streams a device can have. This is limited solely by the hardware used and the implementation.

#### 2.2.5 Buffer Module

The Buffer module encapsulates a single memory buffer. Its purpose is to act as the target for acquisition. The memory of a buffer can be GenTL Consumer allocated or GenTL Producer allocated. The latter could be pre-allocated system memory. The Buffer module also presents Signaling and module configuration capabilities to the GenTL Consumer.

To enable streaming of data at least one buffer has to be announced to the Data Stream module instance and placed into the input buffer pool.

The GenTL Producer may implement preprocessing of the image data which changes image format and/or buffer size. Please refer to chapter 5.5 for a detailed list of the parameters describing the buffer.

The buffer memory might be contiguous or segmented, as described in 5.7.1.

### 2.3 GenTL Module Common Parts

Access and compatibility between GenTL Consumers and GenTL Producers is ensured by the C interface and the description of the behavior of the modules, the Signaling, the Configuration and the acquisition engine.