|  GEN<i>CAM |   | ![img-24.jpeg](img-24.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

##### 3.1.4.5. Pending Acknowledge Packet Failure

There are two possible failure cases using pending acknowledge.

- A complete pending acknowledge packet is lost. In this case the sender will generate a timeout as if the pending acknowledge would not have been sent and it will issue a resend of the command packet with the same request_id. Following chapter 3.1.2, the receiver will reissue a pending acknowledge packet.
- A pending acknowledge packet is received corrupt by the sender. This will trigger a resend of the command packet.

### 3.2. Heartbeat

In order to maintain control in case of an unexpected abrupt detach of the controlling application, a watchdog timer is implemented in the device. This mechanism is called Heartbeat. On start-up of the command sender application, the Access Privilege Register in the device's BRM must be set. With that the Heartbeat timer in the device starts. This Heartbeat timer has to be triggered periodically by a read/write register access from the host to the device. The timeout of the Heartbeat can be adjusted through a register in the bootstrap register map. The presence of a Heartbeat mechanism is indicated by a bit in the device capability register in the device's BRM. It may be disabled through a bit in the device configuration register in the BRM.

In case the Heartbeat counter is not triggered by a register access longer than specified in the Heartbeat Timeout register, the device stops streaming and resets the access privilege status and resets communication parameters. After a Heartbeat timeout, it should be possible to communicate with a device using default communication parameters, for example the baud rate of serial devices. It is technology dependent which parameters are affected.

The Access Privilege register can be set to

- Available – The device is available. The device does not stream data.
- Open (Exclusive) – Only the controlling application has read and write access to the device. It is depending on the technology how this is observed. Other applications/hosts will receive an error trying to access the device's register map.
The exception to this rule is the Access Privilege register itself. This register can be read any time.

When the host changes the state of the Access Privilege register from Open (Exclusive) to Available the device must switch back to default communication parameters after the acknowledge for the write command was sent. The behavior is the same as if the Heartbeat Timeout would run out. This is to allow another application to establish a communication with the device.