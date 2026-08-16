### 1.4 Standard Definitions

This section defines the terms used in this document. An illustration of their inter-relation is provided in the

Device Communication Model section below (See Figure 1-1).

|  Entity | An Entity is an end point located at either side (Host or Device) of a Communication.  |
| --- | --- |
|  Host System | The Host System is the Entity which takes control over a Device. A Host System can be the sink or the source for the data being streamed. Under GenICam the Host System must read and use the GenICam compliant XML file of the Device to control it.  |
|  Device | The Device is an Entity which is controlled by a Host System. A Device can be the source or the sink for streaming data. It can be remote (outside the Host System) or local (in the Host System). Under GenICam the Device must provide a GenICam compliant XML file and a register-based control access.  |
|  Link | A Link is the virtual binding between a Host System and a Device to establish a Communication. A Link is logical and may use one or more physical Connections.  |
|  Connection | A Connection is the physical binding between a Host System and a Device.  |
|  Interface | A virtual end point of the Link between a Device and a Host System.  |
|  Adapter | A physical entity located in the Host System that has one or many Interfaces.  |
|  Communication | A Communication is an exchange of information between two Entities using a Link.  |
|  Channel | A logical point-to-point Communication over a Link. There may be multiple Channels on a single Link.  |
|  Transport Layer | The layer of Communication responsible to transport information between Entities.  |
|  Transmitter | An Entity which acts as the source for streaming data. This may apply to a Host System or a Device.  |
|  Receiver | An Entity which acts as the sink for streaming data. This may apply to a Host System or a Device.  |
|  Transceiver | An Entity which can receive and transmit streaming data. This may apply to a Host System or a Device.  |
|  Peripheral | An Entity which neither acts as a source nor as a sink for streaming data but can be controlled.  |
|  Stream | A flow of data that comes from a source and goes to a sink. A data Stream can be composed of images or chunk of data.  |
|  Stream Channel | A Communication Channel used to transmit a data Stream from a Transmitter (or Transceiver) to a Receiver (or Transceiver).  |
|  Event Channel | A Communication Channel used by the Device to notify the Host System asynchronously of Events. The Host System could also use a Event Channel to communicate events to the Device.  |
|  Control Channel | A Communication Channel used to configure and control a Device. For a Control Channel the Device acts as a server that provides the initial point of Communication for the Host System that acts as a Client. The Communication on a Control Channel  |