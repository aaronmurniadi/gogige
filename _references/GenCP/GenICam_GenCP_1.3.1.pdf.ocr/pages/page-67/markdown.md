|  GEN<i>CAM |   | ![img-71.jpeg](img-71.jpeg)emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

##### 5.5.1.2. Manifest Entry

Each Manifest Entry describes the properties of a single file.

|  Width (Bytes) | Offset (Bytes) | Description  |   |   |
| --- | --- | --- | --- | --- |
|  4 | 0 | GenICam File Version  |   |   |
|   |   |  Bit offset (lsb << x) | Width (bits) | Description  |
|   |   |  0 | 16 | File-Subminor VersionSubminor version of the GenICam file referenced in this entry.  |
|   |   |  16 | 8 | File-Minor VersionMinor version of the GenICam file referenced in this entry.  |
|   |   |  24 | 8 | File-Major VersionMajor version of the GenICam file referenced in this entry.  |
|  4 | 4 | Schema / Filetype / Fileformat  |   |   |
|   |   |  Bit offset (lsb << x) | Width (bits) | Description  |
|   |   |  0 | 3 | File TypeFile type of the file this manifest entry points to.0 = Device XMLThis is the “normal” GenICam device xml containing all device features. This is the one file provided in GenCP until version 1.1.1 = Buffer XMLThis optional XML-file contains only the chunkdata related nodes. This allows the consumer to instantiate one nodemap per buffer in case the buffers containing chunk data and so work on multiple buffers in parallel.2-7 = reserved  |