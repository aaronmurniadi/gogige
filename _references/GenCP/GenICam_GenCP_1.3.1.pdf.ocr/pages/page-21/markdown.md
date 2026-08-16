|  GEN<i>CAM |   | ![img-18.jpeg](img-18.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

CRC of the packet to be wrong or the sender sent the wrong CRC) or if it is too short to carry a correct CCD plus Prefix. In this case the received data is discarded and no answer is sent back to the sender.

The receive buffer should be flushed until no data is received within a maximum packet transfer time or longer.

- The sender must wait after a communication error until all corrupt data is removed and then it sends its command again.
- The receiver discards all corrupt data after a communication error and waits for the sender to resend its command.
- If the underlying technology controls packet handling, it is not necessary to wait for a packet transfer time on failure.
- There is no acknowledge carrying a failure status code in order to prevent the link being flooded with garbage acknowledges.

In case the received Prefix and CCD is correct, the receiver must answer as requested with an appropriate status code and the originator can resend the command.

When there are errors on either side, the original command packet is resent from the sender as described in chapter 3.1.4.3.

In case of failure the sender should retry 3 times to transmit the packet.

##### 3.1.4.2. Timeout

A packet is considered “too short” if the data for a packet has not completely been received within the Packet Transfer Time (PTT) after the first byte of the packet has arrived. The PTT is depending on

- the link speed
- the maximum packet size allowed on the link
- the timeout for the transfer of two consecutive bytes on a link

If an error occurs on either side, the original command packet is resent from the sender as described in chapter 3.1.4.3.

In case of failure, the sender should retry 3 times to transmit the packet.