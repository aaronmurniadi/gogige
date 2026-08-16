|  GEN<I>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

Possible values are:

- Off: The device will only stream data in its native format.
- On: The device will stream all its data in the generic GenDC format.

### 27.2.8 GenDCDescriptor

|  Name | GenDCDescriptor  |
| --- | --- |
|  Category | TransportLayerControl  |
|  Level | Recommended  |
|  Interface | IRegister  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | GenDC specific  |

Returns a preliminary GenDC Descriptor that can be used as reference for the data Container to be streamed out by the device in its current configuration. This information can be used to set up the receiver in advance to be ready for the data Containers to come.

The format of the GenDC Descriptor returned by this feature is defined by the GenICam GenDC standard. The receiver must interpret its content based on the version of this Descriptor.

It is updated in synchronization with PayloadSize and must be static after TLParamsLocked is asserted to guarantee that it accurately represents the typical Descriptor (all the headers) of the Containers that will be streamed out when AcquisitionStart is executed.

In the case of variable Containers, this preliminary Descriptor must represent the maximum size, number of Components and number of Parts that can be sent during the Acquisition. All the Offset fields of the Descriptor must also stay fixed during an acquisition.

The GenDCDescriptor feature must be provided for payloads transmitted using GenDC unless the Container Descriptor is totally variable due to some special device setting (e.g. special Sequencer usage). In that case, this feature should be made unavailable. The Receiver then cannot make any assumption about the Container format to come and should rely solely on the Container Descriptor sent on the TL for each individual Container.

### 27.2.9 GenDCFlowMappingTable

|  Name | GenDCFlowMappingTable  |
| --- | --- |