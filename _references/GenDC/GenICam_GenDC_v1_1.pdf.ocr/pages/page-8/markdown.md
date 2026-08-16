|  GENICAM |   | emva  |
| --- | --- | --- |
|  Version 1.1.0 | GenDC  |   |

## 1.2 Terms and Definitions

|  Name | Description  |
| --- | --- |
|  Component | A single individual element of a Container. A Component has one Header that describes the Component and includes one Part Header for each of its Part(s). A Component can have one or more Parts that belong together.  |
|  Container | An object that contains the complete description and data of simple or complex data buffers. A Container has one Descriptor and one or more Components.  |
|  Descriptor | A structure describing the Container's organization and its data. The Container's Descriptor groups the Container Header, all the Component Headers and their associated Part Headers and fully describes the content of the Container including the data Offsets. To support various use cases in the Transport Layers, three kinds of Descriptors are defined: Prefetch, Preliminary and Final Descriptor but they all share the same layout.  |
|  Flow | A data transport entity that can carry a Descriptor and/or one or more Components and Parts. Allows mapping of a Container's contents to different memory locations and/or parallel transport of it.  |
|  Flow Offset | The position of a specific Part in a Flow. It is specified in bytes from the Flow start.  |
|  Group | An ensemble of GenDC Components related together.  |
|  Header | Structure describing a particular member of the Container's Descriptor.  |
|  Offset | The position of a specific element in the Container. The Offset is always in bytes from the Container start except the Part's FlowOffset.  |
|  Part | A Part is the basic constituent of a Component and contains the data.  |
|  Product | A functional entity using GenDC.  |
|  Transport Layer | An entity transporting the data between receivers and transmitters. It typically handles the transfer and representation of data on a particular physical layer like Ethernet, USB3 or coax cables using a well-defined protocol. Examples of Transport Layers protocol in the context of this specification are GigE Vision™, USB3 Vision™, CoaXPress™ and Camera Link HS™.  |

Table 1-1: Terms and Definitions

## 1.3 Normative References

|  GenICam: | EMVA GenICam Specification  |
| --- | --- |
|  GenICam SFNC: | EMVA GenICam Standard Features Naming Convention  |
|  GenICam PFNC: | EMVA GenICam Pixel Format Naming Convention  |
|  GenICam PFNC Pixel Format Values: | EMVA GenICam Pixel Format Naming Convention Pixel format values list.  |