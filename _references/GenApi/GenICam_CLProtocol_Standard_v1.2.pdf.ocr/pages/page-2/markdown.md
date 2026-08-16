|  GEN<i>CAM |   | ![img-1.jpeg](img-1.jpeg) emva  |
| --- | --- | --- |
|  V1.2 | CLProtocol Standard Module  |   |

## Table of Contents

1 OVERVIEW...3
2 INSTALLING AND REGISTERING CLPROTOCOL DRIVER LIBRARIES 4
3 SELECTING A CLPROTOCOL DRIVER LIBRARY AND IDENTIFYING A CAMERA...5
4 RETRIEVING AN XML FILE FOR A CAMERA...9
5 HANDLING THE BAUD RATE...10
6 STANDARDIZED PROGRAMMING INTERFACES...11

6.1 ISERIAL INTERFACE 11

6.2 CLPROTOCOL INTERFACE...11

HISTORY

|  Version | Date | Changed by | Change  |
| --- | --- | --- | --- |
|  1.0 | 08.12.2009 | Fritz Dierks, Basler | First Draft  |
|  1.0.1 | 30.08.2010 | Fritz Dierks, Basler | Added bootstrap registers  |
|  1.1 | 04.06.2011 | Fritz Dierks, Basler | Added v1.1 extensions  |
|  1.1.1 | 31.10.2011 | Fritz Dierks, Basler | Clarified baud rate handling  |
|  1.1.2 | 08.07.2015 | Fritz Dierks, Basler | Clarified that in CCLPort initialize the cookie value is a reserved value  |
|  1.2 | 09.07.201818.09.2018 | Silvio Voitzsch, Baumer Christoph Zierl, MVTec | - Added Linux support- Added new parameter to stop probing- Added new function to deal with device events- Fixed some spelling mistakes and added some clarifications- Added enumeration procedure for CLSerXXX modules  |