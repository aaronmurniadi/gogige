|  GEN<i>CAM |   | ![img-14.jpeg](img-14.jpeg)emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

- If the new command has the same request_id as the command currently processed, another pending acknowledge packet is sent. In this case the pending acknowledge timeout from the original command is used.
- If the new command has a different request_id the device responds with a GENCP_BUSY status code.

The Processing Time for the inquiry of the Maximum Device Response Time register must not take longer than 50ms.

After the cycle finishes, the host timeout resets to the previously calculated timeout using Maximum Device Response Time and the heartbeat mechanism in the device works as configured before.