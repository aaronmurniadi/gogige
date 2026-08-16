Returns the frequency of the selected Clock.

### 3.64 DeviceSerialPortSelector

|  **Name** | DeviceSerialPortSelector  |
| --- | --- |
|  **Category** | DeviceControl  |
|  **Level** | Recommended  |
|  **Interface** | IEnumeration  |
|  **Access** | Read/(Write)  |
|  **Unit** | -  |
|  **Visibility** | Expert  |
|  **Values** | CameraLink Device-specific  |

Selects which serial port of the device to control.

Possible values are:

- **CameraLink:** Serial port associated to the Camera link connection.

### 3.65 DeviceSerialPortBaudRate

|  **Name** | DeviceSerialPortBaudRate [DeviceSerialPortSelector]  |
| --- | --- |
|  **Category** | DeviceControl  |
|  **Level** | Recommended  |
|  **Interface** | IEnumeration  |
|  **Access** | Read/(Write)  |
|  **Unit** | -  |
|  **Visibility** | Expert  |
|  **Values** | Baud9600 Baud19200 Baud38400 Baud57600 Baud115200 Baud230400 Baud460800 Baud921600 ...  |

This feature controls the baud rate used by the selected serial port.