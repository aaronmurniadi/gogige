|  GEN<i>CAM |   | ![img-21.jpeg](img-21.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

There is a corner case if the device was opened and only one single command was sent or if the request_id got a wraparound to 0, the device was closed and a new application starts with request_id being 0. In this case the CommandResent bit would not be set but the receiver should not discard the command. Therefore, commands with request_id equal to 0 must always be read commands and must always be executed.

If a received command is invalid (combination of command and flags) or is not supported/unknown by the receiver, but at least the CCD is correct (guaranteed by the underlying technology or by CRC) so that the content of the packet is as sent by the originator and the RequestAck bit is set in the flags field, an acknowledge must be sent back with the following content:

- the status code is to be set to GENCP_INVALID_HEADER or GENCP_NOT_IMPLEMENTED (see 4.3.2.1)
- the command_id is copied from the received packet and the acknowledge flag (see 4.3.3) is set
- the length is set to 0, the SCD is discarded
- the request_id is copied from the received packet and left untouched
- CRCs (if existing) must be adjusted

and then it is sent back to the originator.

##### 3.1.4.4. Acknowledge packet failure

If an acknowledge packet is lost on the link, if the CRC of the acknowledge packet is corrupt or if the content is not as expected, the following actions are supposed to happen:

The resend of the command packet uses the same request_id as the original. This allows the receiver to identify a resend in case the request_id is already processed. In this case the command must not be processed again but the previous result should be resent.