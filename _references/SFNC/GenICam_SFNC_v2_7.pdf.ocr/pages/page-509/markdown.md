|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Category | TransportLayerControl  |
| --- | --- |
|  Level | Recommended  |
|  Interface | IRegister  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | GenDC specific  |

Returns the GenDC Container data Flow mapping table that will be used to transport the GenDC Container.

The format of the GenDC Descriptor returned by this feature is defined by the GenICam GenDC standard. The receiver must interpret its content based on the version of this Descriptor.

It is updated in synchronization with PayloadSize and must be static after TLParamsLocked is asserted to guarantee that it accurately represents the Flow mapping of the Containers streamed out when AcquisitionStart is executed.

In the case of variable Containers, this flow mapping table must represent the maximum number of Flows that can be sent during the Acquisition. The size of the Flows and offset of the Flow mapping table must also stay fixed during an acquisition.

The FlowMappingTable feature must be provided for payloads transmitted using GenDC.

#### 27.2.10 DeviceTapGeometry

|  Name | DeviceTapGeometry  |
| --- | --- |
|  Category | TransportLayerControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Geometry_1X_1YGeometry_1X2_1YGeometry_1X2_1Y2Geometry_2X_1YGeometry_2X_1Y2Geometry_2XE_1YGeometry_2XE_1Y2  |