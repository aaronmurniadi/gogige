|  GEN<i>CAM |   | ![img-32.jpeg](img-32.jpeg)emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

|  Status Code (Hex) | Name | Description  |
| --- | --- | --- |
|  0x0000 | GENCP_SUCCESS | Success  |
|  0x8001 | GENCP_NOT_IMPLEMENTED | Command not implemented in the device. This covers for example- Unknown/Unsupported command_id  |
|  0x8002 | GENCP_INVALID_PARAMETER | At least one command parameter of CCD or SCD is invalid or out of range. This covers for example:- CCD-Length field which does not fit to the SCD-Part- Invalid content of the reserved field in the SCD- Write with request_id = 0  |
|  0x8003 | GENCP_INVALID_ADDRESS | Attempt to access a not existing register address.  |
|  0x8004 | GENCP_WRITE_PROTECT | Attempt to write to a read only register.  |
|  0x8005 | GENCP_BAD_ALIGNMENT | Attempt to access registers with an address which is not aligned according to the underlying technology.  |
|  0x8006 | GENCP_ACCESS_DENIED | Attempt to read a non-readable or write a non-writable register address.  |
|  0x8007 | GENCP_BUSY | The command receiver is currently busy.  |
|  0x800B | GENCP_MSG_TIMEOUT | Timeout waiting for an acknowledge.  |