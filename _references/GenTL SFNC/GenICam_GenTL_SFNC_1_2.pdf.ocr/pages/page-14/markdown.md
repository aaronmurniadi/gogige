|  GEN<ICAM |   | emva  |
| --- | --- | --- |
|  Version 1.2.0 | GenTL Standard Features Naming Convention  |   |

### 1.3 Standard Definitions

This section defines the terms used in this document. See Transport layer specific standards, as well as GenICam and SFNC for detailed information.

|  Entity | An Entity is an end point located at either side (Host or Device) of a Communication.  |
| --- | --- |
|  Host System | The Host System is the Entity that takes control over a Device. A Host System can be the sink or the source for the data being streamed. Under GenICam the Host System must read and use the GenICam compliant XML file of the Device to control it.  |
|  Device | The Device is an Entity that is controlled by a Host System. A Device can be the source or the sink for streaming data. It can be remote (outside the Host System) or local (in the Host System). Under GenICam the Device must provide a GenICam compliant XML file and a register-based control access.  |
|  Link | A Link is the virtual binding between a Host System and a Device to establish a Communication. A Link is logical and may use one or more physical Connections.  |
|  Connection | A Connection is the physical binding between a Host System and a Device.  |
|  Interface | A: A virtual endpoint of the Link between a Device and a Host System. B: A GenICam programming interface class, e.g. Uint or Command.  |
|  Consumer | A library or application using an implementation of a GenTL Transport Layer Interface.  |
|  Producer | GenTL Transport Layer Interface implementation.  |
|  Adapter | A physical entity located in the Host System that has one or many Interfaces.  |
|  Communication | A Communication is an exchange of information between two Entities using a Link.  |
|  Channel | A logical point-to-point Communication over a Link. There may be multiple Channels on a single Link.  |
|  Transport Layer | The layer of Communication responsible to transport information between Entities.  |
|  Transmitter | An Entity that acts as the source for streaming data. This may apply to a Host System or a Device.  |
|  Receiver | An Entity that acts as the sink for streaming data. This may apply to a Host System or a Device.  |
|  Transceiver | An Entity that can receive and transmit streaming data. This may apply to a Host System or a Device.  |
|  Peripheral | An Entity that neither acts as a source nor as a sink for streaming data but can be controlled.  |
|  Stream | A flow of data that comes from a source and goes to a sink. A data Stream can be composed of images or chunk of data.  |
|  Stream Channel | A Communication Channel used to transmit a data Stream from a Transmitter (or Transceiver) to a Receiver (or Transceiver).  |
|  Event Channel | A Communication Channel used by the Device to notify the Host System asynchronously of Events. The Host System could also use an Event Channel to  |