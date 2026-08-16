|  GEN<ì>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

## 27.5 Network Statistics features

This section describes the network statistics specific features.

### 27.5.1 NetworkStatistics

|  Name | NetworkStatistics  |
| --- | --- |
|  Category | TransportLayerControl  |
|  Level | Optional  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | -  |

Category that contains statistics pertaining to various modules of the GigE Vision transport layer.

### 27.5.2 oMACControlFunctionEntity

|  Name | oMACControlFunctionEntity  |
| --- | --- |
|  Category | NetworkStatistics  |
|  Level | Optional  |
|  Interface | ICategory  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | -  |

Category that contains statistics pertaining to the MAC control PAUSE function of the device.

When multiple links are aggregated together, this category represents the aggregated statistics of all associated MACs. This is because it is not possible to select each physical network interface in this case.

The counters in this section are defined by the IEEE 802.3 standard. They are generally nonresetable. Rollover behavior and maximum value are device-specific.