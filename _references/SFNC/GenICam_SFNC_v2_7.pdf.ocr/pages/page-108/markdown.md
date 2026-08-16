Resets the device to its power up state. After reset, the device must be rediscovered.

Note that some Transport Layers require the acknowledgement of the DeviceReset command before starting the actual reset of the device.

### 3.52 DeviceIndicatorMode

|  **Name** | DeviceIndicatorMode  |
| --- | --- |
|  **Category** | DeviceControl  |
|  **Level** | Optional  |
|  **Interface** | IEnumeration  |
|  **Access** | Read/Write  |
|  **Unit** | -  |
|  **Visibility** | Expert  |
|  **Values** | Inactive Active ErrorStatus  |

Controls the behavior of the indicators (such as LEDs) showing the status of the Device.

Possible values are:

- **Inactive:** Device's indicators are inactive (Off).
- **Active:** Device's indicators are active showing their respective status.
- **ErrorStatus:** Device's indicators are inactive unless an error occurs.

### 3.53 DeviceFeaturePersistenceStart

|  **Name** | DeviceFeaturePersistenceStart  |
| --- | --- |
|  **Category** | DeviceControl  |
|  **Level** | Optional  |
|  **Interface** | ICommand  |
|  **Access** | (Read)/Write  |
|  **Unit** | -  |
|  **Visibility** | Guru  |
|  **Values** | -  |

Indicate to the device and GenICam XML to get ready for persisting of all streamable features.

Note that device persistence is done by reading the device features and saving them outside of the device.