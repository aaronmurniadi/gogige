|  **Level** | Optional  |
| --- | --- |
|  **Interface** | Integer  |
|  **Access** | Read  |
|  **Unit** | Bps  |
|  **Visibility** | Expert  |
|  **Values** | >0  |

Indicates the speed of transmission of the specified Connection

### 3.33 DeviceConnectionStatus

|  **Name** | DeviceConnectionStatus [DeviceConnectionSelector]  |
| --- | --- |
|  **Category** | DeviceControl  |
|  **Level** | Optional  |
|  **Interface** | Enumeration  |
|  **Access** | Read  |
|  **Unit** | -  |
|  **Visibility** | Expert  |
|  **Values** | Active Inactive  |

Indicates the status of the specified Connection.

Possible values are:

- **Active**: Connection is in use.
- **Inactive**: Connection is not in use.

### 3.34 DeviceLinkSelector

|  **Name** | DeviceLinkSelector  |
| --- | --- |
|  **Category** | DeviceControl  |
|  **Level** | Recommended  |
|  **Interface** | Integer  |
|  **Access** | Read/(Write)  |
|  **Unit** | -  |
|  **Visibility** | Beginner  |
|  **Values** | ≥0  |