|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.1.0 | GenDC  |   |

## 4 GenDC Container formatting, requirements and recommendations

Since the GenDC format is very flexible, it might be suitable to restrict the number of scenarios currently supported in order to simplify and make its usage uniform without restricting its expandability and future usage. This chapter exposes some rules and recommendations that standardize and simplify GenDC encoding, transmission and decoding. Note that none of the possible rules and recommendations below is dictated by the GenDC format itself which aims to be general and flexible. The generic GenDC format targets to be usable in many different scenarios and even outside of the Transmitter to Receiver transmission context.

Each individual Transport Layer Protocol (TLP) supporting GenDC is also free to add its own particular rules and restrictions regarding the GenDC Container transport to simplify the GenDC usage in its particular context. Recommendations can be ignored if necessary but the TLP must not modify or override basic formal requirements stated in the GenDC standard.

### 4.1.1 Requirements and recommendations

This section describes requirements and recommendations to be compliant with the GenDC specification: The GenDC Containers are always stored in little Endian. This permits to have a unified way to encode and decode a GenDC Container Descriptor and data independently of the CPU, Transport Layer Protocol or context where it is used.

[ R-004] The GenDC Container Descriptor must be always stored in little-endian ordering.

[ R-005] Container data Part in PFNC format must use little-endian ordering.

GenDC Container format and ordering: In the context of data exchange between a Transmitter and a Receiver, a virtual GenDC Container is always represented as a continuous block starting with a GenDC Descriptor immediately followed by a continuous Container's data section. The DataOffset field represents the offset of the data from the Descriptor start.

[ R-006] For Transmission, a GenDC Container is always represented as a continuous block of linear memory starting with the Descriptor.

[ R-007] The Part's DataOffset is always the offset of the data in bytes from the start of the Descriptor.

GenDC Container transmission Flow(s): In the context of data exchange between a Transmitter and a Receiver, the data section(s) of a virtual GenDC Container are transmitted using one or more data Flow(s). Flows allow splitting a Container in sections to facilitate parallel transmission. Flows have a FlowId that is numbered sequentially starting with 0. A Flow can contain a Container Descriptor and/or one or more data Parts. The GenDC Descriptor is always sent as early as possible in Flow 0. The Container Descriptor should be sent as soon as possible to provide its information early to the receiver and help decoding and preprocessing of the data. A Flow can carry one or many data Parts but a Part must only be mapped to a single Flow. All the Component's Parts sharing a Flow are mapped in the same order as they are listed in the GenDC Descriptor. Due the support of variable data content and hardware based memory alignment constraints a Container can contain unused memory areas. It is up to the receiver to store the data sections of a Container transported on different Flows separately in a single or multiple target buffers after the transport.

[ R-008] The Part's FlowOffset is always the offset of the data in bytes from the start of the Flow specified by the Part's FlowId.