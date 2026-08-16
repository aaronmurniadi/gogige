### 3.54 DeviceFeaturePersistenceEnd

|  **Name** | DeviceFeaturePersistenceEnd  |
| --- | --- |
|  **Category** | DeviceControl  |
|  **Level** | Optional  |
|  **Interface** | ICommand  |
|  **Access** | (Read)/Write  |
|  **Unit** | -  |
|  **Visibility** | Guru  |
|  **Values** | -  |

Indicate to the device the end of feature persistence.

### 3.55 DeviceRegistersStreamingStart

|  **Name** | DeviceRegistersStreamingStart  |
| --- | --- |
|  **Category** | DeviceControl  |
|  **Level** | Recommended  |
|  **Interface** | ICommand  |
|  **Access** | (Read)/Write  |
|  **Unit** | -  |
|  **Visibility** | Guru  |
|  **Values** | -  |

Prepare the device for registers streaming without checking for consistency.

If the camera implements this feature, GenApi guarantees using it to announce register streaming.

If the feature is present, but currently not writable (locked), the application must not start register streaming and must avoid switching the access mode and range verification off until the feature becomes writable again.

### 3.56 DeviceRegistersStreamingEnd

|  **Name** | DeviceRegistersStreamingEnd  |
| --- | --- |
|  **Category** | DeviceControl  |
|  **Level** | Recommended  |
|  **Interface** | ICommand  |
|  **Access** | (Read)/Write  |
|  **Unit** | -  |