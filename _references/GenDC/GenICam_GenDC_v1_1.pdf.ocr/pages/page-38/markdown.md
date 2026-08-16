|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.1.0 | GenDC  |   |

[ R-009] The Descriptor is always transferred in Flow 0.

[ R-010] A Part must only be mapped to a single Flow.

[ R-011] Flow must be numbered sequentially starting from 0.

GenDC variable Container Transmission: In the context of data exchange between a Transmitter and a Receiver, if a GenDC Container has some elements that are variable (e.g. Component's Size Y, Region X,Y Offset, Data Size Timestamps, Sequence length...), the corresponding flag(s) of the "VariableFields" field of the Container Header must be set. If the "VariableFields" field is not null, a preliminary Descriptor should be sent to describe the characteristics of the maximum Container that will be transmitted. This preliminary Descriptor must represent the maximum size, number of Components and number of Parts that can be sent. Also, for variable Container, a final GenDC Descriptor including updated Container information must be sent immediately after transmission of the Container's data section.

[ CR-012] If a Container has variable content during the transmission the VariableFields flags of the Container must be set accordingly.

[ CR-013] If any of the Container's VariableFields flags is set, a final Descriptor must be sent as soon as possible but at least just after the transmission of the data section.

GenDC Container Storage: A GenDC Container can be stored on external storage (e.g. as a file). If stored in GenDC binary format, it must use the extension ".gendc" for the filename.

[ R-014] When storing a GenDC Container using the ".gendc" file extension, the standard GenDC binary format must be used.

[ R-015] When storing a GenDC Container, a linear Container must be used.

GenDC Container Metadata: Metadata related to a Container must be added to the Container as an additional and separate Metadata Component (e.g. A Container including the Components of a 3D scene (Range, Confidence, ... ) with an additional GenICam Chunk Metadata Component to store the 3D scene additional information).

[ R-016] When adding metadata to a Container, this must be done using a separate Metadata Component.