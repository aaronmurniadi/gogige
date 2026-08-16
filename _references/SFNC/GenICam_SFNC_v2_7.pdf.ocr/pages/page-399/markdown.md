|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

Note: Pixel unit can be used if the distance data is given as a disparity image (e.g. output of stereo matching), since disparity is a displacement in the image plane, which is measured in pixel.

Device specific values can be used to indicate other meaning to distance data. The angle unit used is degrees.

### 21.4.6 Scan3dCoordinateSystem

|  Name | Scan3dCoordinateSystem[Scan3dExtractionSelector]  |
| --- | --- |
|  Category | Scan3dControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Expert  |
|  Values | Cartesian Spherical Cylindrical Device-specific  |

Specifies the Coordinate system to use for the device.

Note that the third coordinate, C, is the "distance" (Z or Rho) coordinates independent of the coordinate system used.

Possible values are:

- Cartesian: Default value. 3-axis orthogonal, right-hand X-Y-Z.
- Spherical: A Theta-Phi-Rho coordinate system.
- Cylindrical: A Theta-Y-Rho coordinate system.

In addition to the previous standard values, a device might also provide device-specific values (e.g. non-orthogonal coordinate system).

### 21.4.7 Scan3dOutputMode

|  Name | Scan3dOutputMode[Scan3dExtractionSelector]  |
| --- | --- |
|  Category | Scan3dControl  |
|  Level | Optional  |
|  Interface | IEnumeration  |
|  Access | Read/(Write)  |