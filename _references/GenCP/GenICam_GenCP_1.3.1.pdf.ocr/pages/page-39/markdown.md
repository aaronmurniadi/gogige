|  GEN<i>CAM |   | ![img-37.jpeg](img-37.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

the Device Capability register. In case the length_written field (and the 2 reserved bytes) is sent, the length field is to be set to 4. In case the length_written field is not sent the length field is 0.

#### 4.4.5. Pending Acknowledge

The pending acknowledge informs the sender that the command, sent with the given request_id, needs more time to execute than stated in the MDRT register. This allows the temporary adjustment of the timeout mechanism on the command sender side. This “new” temporary timeout is only valid for the command referenced by request_id. Multiple pending acknowledges can be sent consecutively. The start time for the timeout specified is the time when the pending ack is sent, assuming that the time needed to transfer the command is roughly known. The timeout is not referring to the time the original command is sent.

|  Width (Bytes) | Offset (Bytes) | Description  |
| --- | --- | --- |
|  Prefix  |   |   |
|  CCD-ACK (command_id = PENDING_ACK)  |   |   |
|  2 | 0 | reservedReserved, set to 0.  |
|  2 | 2 | temporary timeoutTemporary timeout for the command sent with the givenrequest_id. The timeout is specified in ms. The reference time/start time for the temporary timeout is the time the PendingAck is sent.  |
|  Postfix  |   |   |

Table 12 – Pending Ack SCD-Fields

#### 4.4.6. ReadMemStacked Command

The ReadMemStacked Command allows sending multiple read requests in one packet. The resulting data must not exceed the maximum packet size. Start address and length of any read access is byte aligned unless the underlying technology is not. The count of read commands within the packet n has to be deduced by the receiver using the packet size sent by the transmitter.

|  Width (Bytes) | Offset (Bytes) | Description  |
| --- | --- | --- |