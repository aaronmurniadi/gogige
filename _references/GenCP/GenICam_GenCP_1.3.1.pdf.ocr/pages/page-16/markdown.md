|  GEN<i>CAM |   | ![img-11.jpeg](img-11.jpeg) emva  |
| --- | --- | --- |
|  Version 1.3.1 | GenCP Standard  |   |

This is to prevent that with the start of a communication an application uses request_id = 0 and sends just 1 command. Then a second application would also start a new communication and would again use request_id = 0. In this case it needs to be ensured that the second communication does not get an “old” ack.

The round trip time for a command and the according acknowledge is

Command Transfer Time + Processing Time + Acknowledge Transfer time

When calculating the timeout time for the command cycle, a host must therefore consider:

- the transfer time of the maximum packet size on a given link speed
- the Maximum Device Response Time, which is provided via a bootstrap register
- some margin for technology-dependent delays, which may occur on the link

Reading the Maximum Device Response Time (MDRT) register should not exceed 50 ms in order to guarantee a responsive device. The maximum device response time for any other read or write operation should not exceed 300 ms. This plus the maximum packet transfer time allows the host to calculate a timeout value.