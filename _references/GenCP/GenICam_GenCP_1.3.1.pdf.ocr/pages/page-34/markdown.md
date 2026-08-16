|  0x800E | GENCP_INVALID_HEADER | The header of the received command is invalid. This includes CCD and SCD fields but not the command payload. This covers for example: - Invalid combinations of flags in the CCD-Flags field - The transmitted packet length does not fit to expected size with the given command and CCD-Length incl. Prefix and Postfix.  |
| --- | --- | --- |
|  0x800F | GENCP_WRONG_CONFIG | The current receiver configuration does not allow the execution of the sent command.  |
|  ... |  |   |
|  0x8FFF | GENCP_ERROR | Generic error.  |

Table 6 – Status Codes