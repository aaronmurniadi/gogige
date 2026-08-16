|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.1.0 | GenDC  |   |

### 2.2.2 GenDC Container Header Description

|  Width (Bytes) | Offset (Bytes) | Description  |
| --- | --- | --- |
|  4 | 0 | **Signature** = “GNDC” Unique signature identifying a GenDC Container: a FourCC code encoded as 4 ASCII characters not null terminated (GDC_SIGNATURE = “GNDC” = 0x43444E47).  |
|  3 | 4 | **Version** = Container Descriptor version coded as Major.Minor.SubMinor. (e.g. GDC_VERSION => 01.00.00 = 0x01,0x00, 0x00 as three consecutive 8 bit fields). The version corresponds to the GenDC specification that the Container complies to. The following versioning rules apply: - The layout of the first 8 bytes of the Container Header will never change in a way that the binary compatibility is broken. - The major version is incremented if the layout of Headers changes or major rules change, that break the compatibility. - The minor version is incremented if new Components, Parts or flags are added that need to be interpreted. - The sub minor version is incremented for clarifications to the standard document. - If an implementation supporting a major version of the specification (e.g. 1.0.0), receives a Container with the same major version (e.g. 1.x.x) but a higher minor and/or sub minor version (e.g. 1.1.2), it is guaranteed to be able to interpret the known (1.0.0) Components and Parts.  |
|  1 | 7 | **Reserved** Reserved for future use, set to 0.  |
|  2 | 8 | **HeaderType** = Unique Header format identifier (Container Header) (GDC_CONTAINER_HEADER = 0x1000). A GenDC Container must always start with a Container Header.  |