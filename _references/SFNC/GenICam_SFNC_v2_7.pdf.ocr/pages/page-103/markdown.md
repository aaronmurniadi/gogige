|  **Category** | DeviceControl  |
| --- | --- |
|  **Level** | Optional  |
|  **Interface** | IEnumeration  |
|  **Access** | Read/Write  |
|  **Unit** | -  |
|  **Visibility** | Expert  |
|  **Values** | On Off  |

Activate or deactivate the Link's heartbeat.

Possible values are:

- **Off:** Disables the Link heartbeat.
- **On:** Enables the Link heartbeat.

### 3.40 DeviceLinkHeartbeatTimeout

|  **Name** | DeviceLinkHeartbeatTimeout [DeviceLinkSelector]  |
| --- | --- |
|  **Category** | DeviceControl  |
|  **Level** | Optional  |
|  **Interface** | IFloat  |
|  **Access** | Read/Write  |
|  **Unit** | us  |
|  **Visibility** | Guru  |
|  **Values** | >0  |

Controls the current heartbeat timeout of the specific Link.

### 3.41 DeviceLinkCommandTimeout

|  **Name** | DeviceLinkCommandTimeout [DeviceLinkSelector]  |
| --- | --- |
|  **Category** | DeviceControl  |
|  **Level** | Optional  |
|  **Interface** | IFloat  |
|  **Access** | Read  |
|  **Unit** | us  |
|  **Visibility** | Guru  |