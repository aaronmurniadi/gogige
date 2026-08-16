|  GENICAM |   | ![img-65.jpeg](img-65.jpeg) emva  |
| --- | --- | --- |
|  Version 2.1.1 | Standard  |   |

Applications (libraries) supporting schema versions 1.1 or newer should check, whether the XML file is based on schema version 1.1 or newer. If so, it must implement the endianess as follows:

- When reading/writing the data of any size, it has free choice of the access method, READMEM/WRITEMEM or READREG/WRITEREG, provided it follows the guidelines listed below. The READMEM/WRITEMEM more naturally matches the GenApi data access model and is therefore recommended option, wherever suitable.
- When reading/writing the data using READMEM/WRITEMEM (recommended option, particularly if WRITEMEM is supported by the camera), the data must be passed "as is" to the GenApi.
- When READREG/WRITEREG has to be used for any reason (eg. because WRITEMEM is not implemented by the camera), the transport layer (or other IPort implementation) has to revert the data back to the "camera order" before passing it to GenApi. This means no extra operation for big endian cameras. For little endian cameras it means to flip each 4-byte word read/written before passing it to GenApi, so that it works with the same data layout (camera's native byte order) as if READMEM/WRITEMEM was used. Note that the application knows the camera order, which can be read from the DeviceMode bootstrap register.

#### 3.1.2 Behavior of products based on schema version 1.0

Cameras providing XML file based on schema version 1.0 should implement endianess as follows to reach best possible compatibility (although for historical reasons slightly different implementations exist in the field):

- The <endianess> tags of all registers have to be set to "BigEndian", regardless its actual endianess.
- String registers should always be bigger than 4 bytes (the Length attribute of the registers should be bigger than 4). Device specific (non-bootstrap) strings should be read-only, particularly on little endian cameras.
- All integer and float registers should be exactly 32-bit, particularly on little endian cameras.</endianess>

Applications supporting schema version 1.0 (or applications supporting multiple schema versions, when working with cameras based on schema version 1.0) should behave as follows to reach best possible compatibility (although for historical reasons slightly different implementations exist in the field):

- When reading/writing a 4-byte data, it should assume it is a integer/float register and use READREG/WRITEREG. The READREG/WRITEREG data are always in network order and it should be passed as such to GenApi.
- When reading data longer than 4 bytes, it should assume it is a string register and use READMEM. Pass the data "as is" to GenApi.
- Writing data longer than 4 bytes should be considered as non-reliable for little endian cameras.

#### 3.1.3 Passing the schema version to the IPort implementation

The register access happens in the IPort implementation (in the IPort's Read/Write functions), ie. in the "transport layer" software component. This component typically does not retrieve and parse the XML (this is task of a top-level client component). The IPort implementation