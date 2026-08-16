### 3 Device Control

Device control features provides general information and control for the device (camera) and its sensor. This is mainly used to identify the device during the enumeration process and to obtain information about the sensor resolution. Other information and controls pertaining to the general state of the device are also included in this category.

#### 3.1 DeviceControl

|  **Name** | DeviceControl  |
| --- | --- |
|  **Category** | Root  |
|  **Level** | Recommended  |
|  **Interface** | ICategory  |
|  **Access** | Read  |
|  **Unit** | -  |
|  **Visibility** | Beginner  |
|  **Values** | -  |

Category for device information and control.

#### 3.2 DeviceType

|  **Name** | DeviceType  |
| --- | --- |
|  **Category** | DeviceControl  |
|  **Level** | Optional  |
|  **Interface** | IEnumeration  |
|  **Access** | Read  |
|  **Unit** | -  |
|  **Visibility** | Guru  |
|  **Values** | Transmitter Receiver Transceiver Peripheral  |

Returns the device type.

Possible values are:

- **Transmitter:** Data stream transmitter device.
- **Receiver:** Data stream receiver device.
- **Transceiver:** Data stream receiver and transmitter device.
- **Peripheral:** Controllable device (with no data stream handling).