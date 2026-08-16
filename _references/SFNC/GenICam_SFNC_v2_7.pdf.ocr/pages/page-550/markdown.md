|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

### 27.4.72 GevSCZoneDirectionAll

|  Name | GevSCZoneDirectionAll[GevStreamChannelSelector]  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IInteger  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | 00000000h to FFFFFFFFh  |

Reports the transmission direction of each zone transmitted on the selected stream channel.

This feature is represented as an unsigned integer. The most significant bit of the range of valid values reports the direction of the first zone (Zone ID 0) while the least significant bit represents the direction of the last zone (Zone ID 1).

### 27.4.73 GevSCZoneConfigurationLock

|  Name | GevSCZoneConfigurationLock[GevStreamChannelSelector]  |
| --- | --- |
|  Category | GigEVision  |
|  Level | Optional  |
|  Interface | IBoolean  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | True False  |

Controls whether the selected stream channel multi-zone configuration is locked. When locked, the GVSP transmitter is not allowed to change the number of zones and their direction during block acquisition and transmission.