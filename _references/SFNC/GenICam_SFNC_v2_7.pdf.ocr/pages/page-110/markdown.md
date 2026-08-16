|  **Visibility** | Guru  |
| --- | --- |
|  **Values** | -  |

Announce the end of registers streaming. This will do a register set validation for consistency and activate it. This will also update the **DeviceRegistersValid** flag.

### 3.57 DeviceRegistersCheck

|  **Name** | DeviceRegistersCheck  |
| --- | --- |
|  **Category** | DeviceControl  |
|  **Level** | Recommended  |
|  **Interface** | ICommand  |
|  **Access** | (Read)/Write  |
|  **Unit** | -  |
|  **Visibility** | Expert  |
|  **Values** | -  |

Perform the validation of the current register set for consistency. This will update the **DeviceRegistersValid** flag.

### 3.58 DeviceRegistersValid

|  **Name** | DeviceRegistersValid  |
| --- | --- |
|  **Category** | DeviceControl  |
|  **Level** | Recommended  |
|  **Interface** | IBoolean  |
|  **Access** | Read  |
|  **Unit** | -  |
|  **Visibility** | Expert  |
|  **Values** | True False  |

Returns if the current register set is valid and consistent.

### 3.59 DeviceRegistersEndianness

|  **Name** | DeviceRegistersEndianness  |
| --- | --- |
|  **Category** | DeviceControl  |
|  **Level** | Optional  |
|  **Interface** | IEnumeration  |