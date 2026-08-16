|  **GEN<i>CAM** |   |   |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

#### **4.2. Prefix**

In case the underlying technology does not provide an addressing schema for multiple communication channels or does not provide a checksum mechanism, the protocol needs to provide such services. A packet then contains not only command specific data but also has to mimic an addressing scheme between the device and host. Also we need to be able to support multiple communication channels on a given Link and a checksum.

In case such services are provided by the underlying technology, the Prefix can simply be omitted.

#### **4.3. Common Command Data**

The Common Command Data section is technology agnostic.