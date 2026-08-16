|  **GEN<i>CAM** |   |   |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

### **1.9. Heartbeat**

In case a serial device supports multiple baud rates, the Heartbeat mechanism must be supported in order to ensure a fall back after a faulty baud rate configuration. In case the device loses the Heartbeat, the link falls back to the default 9600 baud so that the host can re-establish communication after a switch to a baud rate that is too high. In case the device only supports the default baud rate, the Heartbeat mechanism is optional.