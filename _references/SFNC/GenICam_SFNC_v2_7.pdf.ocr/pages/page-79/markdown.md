|  **GEN<i>CAM** |   |   |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  ValueArrayCandidates | O | IString | R | - | I | Advertises selector/value feature pairs suitable for use with GenApi ValueArrayAdapter utility which allows efficient extraction of features under selectors (for example Chunk features).  |
| --- | --- | --- | --- | --- | --- | --- |

## 2.25 Transport Layer Control

Contains the features related to the Transport Layer Control (See the Transport Layer Control chapter for details).

Table 2-25: Transport Layer Control Summary

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  TransportLayerControl | R | ICategory | R | - | B | Category that contains the transport Layer control features.  |
|  TLParamsLocked | M | IInteger | R/W | - | I | Used by the Transport Layer to prevent critical features from changing during acquisition.  |
|  TLParamsLockedSelector | O | IEnumeration | R/(W) | - | G | Selects the type of feature for which the locking behavior will be configured.  |
|  TLParamsLockedState[TLParamsLockedSelector] | O | IBoolean | R/(W) | - | G | Controls if the selected parameters are locked during acquisition.  |
|  PayloadSize | R | IInteger | R | B | E | Provides the number of bytes transferred for each data buffer or chunk on the stream channel.  |
|  GenDCStreamingMode | R | IEnumeration | R/W | - | G | Controls the device's streaming format.  |
|  GenDCStreamingStatus | R | IEnumeration | R | - | G | Returns whether the current device data streaming format is GenDC.  |
|  GenDCDescriptor | R | IRegister | R | - | G | Returns a preliminary GenDC Descriptor that can be used as reference for the data Container to be streamed out by the device in its current configuration.  |
|  GenDCFlowMappingTable | R | IRegister | R | - | G | Returns the GenDC Container data Flow mapping table that will be used to transport the GenDC Container.  |
|  DeviceTapGeometry | R | IEnumeration | R/(W) | - | E | This device tap geometry feature describes the geometrical properties characterizing the taps of a camera as presented at the output of the device.  |