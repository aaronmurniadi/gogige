|  GEN<i>CAM |   | ![img-25.jpeg](img-25.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

### 3.3. GenICam File

A GenCP device must be register based. A manufacturer must provide access to a GenICam file describing the register map of the device.

The GenICam file must be stored within the device so that it can be retrieved by the host. The file may be stored and delivered either in uncompressed or compressed format. In case it is compressed it is up to the controlling host to deflate the file.

#### 3.3.1. Manifest Table

A GenCP device may provide multiple GenICam files complying with different GenICam Schema versions. A so called “Manifest Table” register block contains a list of entries, providing information like file versions, complying schema versions, and register addresses. A description of the Manifest register block can be found in the Bootstrap Register Map section of this document.

#### 3.3.2. Retrieval

It is the responsibility of the host software to retrieve the file from the device reading the device's register space using the GenCP Protocol.

#### 3.3.3. Compression

The compression methods used in case the GenICam file is stored in the device in a compressed format are DEFLATE and STORE of the .zip file format. File extension for compressed files is zip.