|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|   | Receiver Transceiver Peripheral  |
| --- | --- |

This feature is deprecated (See DeviceType). It was representing the class of the device.

Note: The GevDeviceClass feature returns Transmitter for cameras.

### 27.4.6 GevDeviceModeCharacterSet (Deprecated)

|  Name | GevDeviceModeCharacterSet  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Invisible  |
|  Values | UTF8  |

This feature is deprecated (See DeviceCharacterSet). It was representing the character set used by all the strings of the device.

### 27.4.7 GevPhysicalLinkConfiguration

|  Name | GevPhysicalLinkConfiguration  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | SingleLink MultiLink StaticLAG DynamicLAG  |

Controls the principal physical link configuration to use on next restart/power-up of the device.