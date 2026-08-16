|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|   | StreamChannel0AllInTransmission StreamChannel0UnconditionalStreaming StreamChannel0ExtendedChunkData StreamChannel1BigAndLittleEndian StreamChannel1IPReassembly StreamChannel1MultiZone StreamChannel1PacketResendDestination StreamChannel1AllInTransmission StreamChannel1UnconditionalStreaming StreamChannel1ExtendedChunkData StreamChannel2BigAndLittleEndian StreamChannel2IPReassembly StreamChannel2MultiZone StreamChannel2PacketResendDestination StreamChannel2AllInTransmission StreamChannel2UnconditionalStreaming StreamChannel2ExtendedChunkData ...  |
| --- | --- |

Selects the GEV option to interrogate for existing support.

Note: The IP reassembly options (StreamChannel0IPReassembly, StreamChannel1IPReassembly, ...) are only applicable to GVSP receiver stream channels.

### 27.4.11 GevSupportedOption

|  Name | GevSupportedOption[GevSupportedOptionSelector]  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | True False  |

Returns if the selected GEV option is supported.