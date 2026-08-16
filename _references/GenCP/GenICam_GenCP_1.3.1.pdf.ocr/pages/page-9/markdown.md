|  GEN<i>CAM |   | ![img-2.jpeg](img-2.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

For example, an ASCII based serial link protocol could be used in the generic CLProtocol module to communicate with a manufacturer's device over the Camera Link's serial link. At boot up, the generic CLProtocol module would allow the configuration of the serial link. A “generic” software could download the GenICam file by accessing the camera's registers. The software can then provide native GenICam (like GigE Vision) access to the device without the need for the camera vendor to provide a platform/operating system-specific software running on the host, implementing the translation between GenICam register access and manufacturer proprietary protocols.

### 1.3. Abstract

The protocol is packet based. It follows a simple command/acknowledge scheme to provide resend and timeout capabilities, adding minimum overhead.

The Bootstrap Register Map (BRM) resides in a 64-bit register space. The 64 Kbytes starting on address zero contain technology agnostic information like manufacturer name, model name, etc., and provide a directory for technology specific settings.

In order to locate the GenICam file for a device, software would need to retrieve a list of available GenICam files, called the manifest, from the device's register map. The software would then pick the best fitting GenICam file from the list and access via the device's register map.